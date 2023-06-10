package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(usernameDB string, passwordDB string, host string, port string, nameDB string) (*gorm.DB, error) {
	var connectionAddress = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		usernameDB, passwordDB, host, port, nameDB,
	)
	return gorm.Open(mysql.Open(connectionAddress), &gorm.Config{})
}
