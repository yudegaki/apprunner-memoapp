package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yudegaki/apprunner-memoapp/internal/controllers"
)

func Router(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	r.GET("/users", controllers.GetAllUsers)
}
