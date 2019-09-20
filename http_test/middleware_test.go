package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func ResponseTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		// 打印处理流程耗时
		ms := time.Since(start).Milliseconds()
		// 该 header 并不能设置成功
		c.Header("X-Response-Time", fmt.Sprintf("%d", ms))
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		ms := time.Since(start).Milliseconds()
		fmt.Printf("%s %s - %dms\n", c.Request.Method, c.Request.URL, ms)
	}
}

func Response() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 该 header 能设置成功
		c.Header("BeforeString", time.Now().String())
		c.String(http.StatusOK, "hello world")
		// 该 header 不能设置成功
		c.Header("AfterString", time.Now().String())
	}
}

func TestMiddleware(t *testing.T) {
	func() {
		r := gin.Default()
		r.Use(ResponseTime())
		r.Use(Logger())
		r.Use(Response())
		r.Run(":8080")
	}()
}
