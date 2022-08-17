package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yudegaki/apprunner-memoapp/internal/external"
)

func Transaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := external.DB.Begin()
		defer func() {
			if 400 <= c.Writer.Status() {
				db.Rollback()
				return
			}
			db.Commit()
		}()
		c.Set("db", db)
		c.Next()
	}
}
