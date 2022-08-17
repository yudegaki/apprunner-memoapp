package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yudegaki/apprunner-memoapp/internal/config"
	"github.com/yudegaki/apprunner-memoapp/internal/external"
	"github.com/yudegaki/apprunner-memoapp/internal/middleware"
)

func main() {
	// Initialize the database
	external.InitDB()

	// Initialize the router
	r := gin.Default()
	r.Use(middleware.Transaction())
	middleware.Router(r)

	r.Run(fmt.Sprintf(":%d", config.Port))
}
