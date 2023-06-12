package main

import (
	database "Jakpat_Test_2/database/mysql"
	"Jakpat_Test_2/handler"
	"Jakpat_Test_2/middleware"
	"Jakpat_Test_2/models"
	"Jakpat_Test_2/repository"
	"Jakpat_Test_2/usecase"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "Jakpat_Test_2/docs"

	cron "github.com/robfig/cron/v3"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal().Err(envErr).Msg("cannot load environment")
	}

	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "8080"
	}

	var dbUsername = os.Getenv("DB_USERNAME")
	var dbPassword = os.Getenv("DB_PASSWORD")
	var dbName = os.Getenv("DB_DATABASE")
	var dbHost = os.Getenv("DB_HOST")
	var dbPort = os.Getenv("DB_PORT")
	db, dbErr := database.ConnectDB(dbUsername, dbPassword, dbHost, dbPort, dbName)

	if dbErr != nil {
		log.Fatal().Err(dbErr).Msg("cannot connect to database")
	}

	//auto migrate database
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Inventory{})
	db.AutoMigrate(&models.SalesOrder{})

	//repository layer
	userRepository := repository.NewUserRepositoryImpl(db)
	inventoryRepository := repository.NewInventoryRepositoryImpl(db)
	orderRepository := repository.NewOrderRepositoryImpl(db)

	//usecase layer
	userUsecase := usecase.NewUserUsecaseImpl(userRepository)
	inventoryUsecase := usecase.NewInventoryUsecaseImpl(inventoryRepository)
	orderUsecase := usecase.NewOrderUsecaseImpl(orderRepository, inventoryRepository)

	//handler layer
	userHandler := handler.NewUserHandlerImpl(userUsecase)
	inventoryHandler := handler.NewInventoryHandlerImpl(inventoryUsecase)
	orderHandler := handler.NewOrderHandlerImpl(orderUsecase)

	router := gin.Default()
	router.Use(cors.Default())

	// route swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authMiddleware := middleware.JWTAuthMiddleware(userUsecase)
	api := router.Group("/api/v1")

	//route user group
	categoryRouter := api.Group("/user")
	categoryRouter.POST("/register", userHandler.Register)
	categoryRouter.POST("/login", userHandler.Login)

	//route inventory group
	inventoryRouter := api.Group("/inventory", authMiddleware)
	inventoryRouter.POST("", inventoryHandler.Create)
	inventoryRouter.PUT("/:id", inventoryHandler.Update)
	inventoryRouter.GET("/:id", inventoryHandler.GetById)
	inventoryRouter.GET("/sku/:sku", inventoryHandler.GetBySku)
	inventoryRouter.GET("", inventoryHandler.GetBySeller)
	inventoryRouter.DELETE("/delete/:id", inventoryHandler.DeleteById)

	//route order group
	orderRouter := api.Group("/order", authMiddleware)
	orderRouter.POST("", orderHandler.Create)
	orderRouter.PUT("/:id", orderHandler.Update)
	orderRouter.GET("/:id", orderHandler.GetById)
	orderRouter.GET("", orderHandler.GetBySeller)
	orderRouter.DELETE("/delete/:id", orderHandler.DeleteById)

	router.Run(":" + appPort)

	//cron scheduler
	//cron scheduler
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))
	defer scheduler.Stop()

	scheduler.AddFunc("*/5 * * * *", func() { orderUsecase.SetExpiredOrder() })
	go scheduler.Start()
}
