package external

import (
	"fmt"

	"github.com/yudegaki/apprunner-memoapp/internal/config"
	"github.com/yudegaki/apprunner-memoapp/internal/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	port := config.DB_PORT
	host := config.DB_HOST
	user := config.DB_USER
	password := config.DB_PASS
	dbName := config.DB_NAME

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	Initialize()
}

func Initialize() {
	if exist := DB.Migrator().HasTable(&repositories.User{}); !exist {
		userInitialize()
	}
	if exist := DB.Migrator().HasTable(&repositories.Post{}); !exist {
		postInitialize()
	}
}

func userInitialize() {
	// Delete table
	DB.Migrator().DropTable(&repositories.User{})
	// Create table
	DB.AutoMigrate(&repositories.User{})
	// Initialize table
	for i := 1; i < 10; i++ {
		tmp := repositories.User{
			Name:     fmt.Sprintf("user%d", i),
			Password: fmt.Sprintf("password%d", i),
		}
		DB.Create(&tmp)
	}
}

func postInitialize() {
	// Delete table
	DB.Migrator().DropTable(&repositories.Post{})
	// Create table
	DB.AutoMigrate(&repositories.Post{})
	// Initialize table
	for i := 1; i < 20; i++ {
		tmp := repositories.Post{
			Title: fmt.Sprintf("POST-%d", i),
			Body:  fmt.Sprintf("HELLO WORLD HOGE FUGA PIYO %d", i),
		}
		DB.Create(&tmp)
	}
}
