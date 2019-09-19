package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUrl(t *testing.T) {
	func() {
		r := gin.Default()
		r.GET("/test/value", func(c *gin.Context) {
			req := struct {
				Key string `form:"Key" binding:"required"`
			}{}
			c.BindQuery(&req)
			fmt.Println(req.Key)
			c.JSON(http.StatusOK, gin.H{"Key": req.Key})
		})
		r.Run(":8000")
	}()

}
