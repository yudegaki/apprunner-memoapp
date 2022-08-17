package external

import (
	"fmt"
	"log"
	"os"

	"github.com/yudegaki/apprunner-memoapp/internal/config"
	"github.com/yudegaki/apprunner-memoapp/internal/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// host := "127.0.0.1"
	host := os.Getenv("MYSQL_HOST")
	port := config.DBPort
	user := config.DBUser
	password := config.DBPassword
	dbName := config.DBName

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	log.Default().Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	migrate()
	Initialize()
}

func migrate() {
	DB.AutoMigrate(&repositories.User{})
}

func Initialize() {
	// User initialization
	for i := 1; i < 10; i++ {
		tmp := repositories.User{Name: fmt.Sprintf("user%d", i), Password: fmt.Sprintf("password%d", i)}
		DB.Create(&tmp)
	}
}
