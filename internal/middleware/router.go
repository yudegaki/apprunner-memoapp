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
	// r.GET("/users/:id", controllers.GetUser)

	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:id", controllers.GetPostDetail)
	r.GET("/posts/:id/edit", controllers.UpdatePost)
	r.POST("/posts/create", controllers.CreatePost)
	r.GET("/posts/:id/delete", controllers.DeletePost)
}
