package main

import (
	"fmt"
	"net/http"
	"net/url"
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

func TestParseURL(t *testing.T) {
	origin := "rtsp://192.168.3.101:9554/fss/vod/14/581806a8-3ca0-423f-a280-25d4dcfb38c5?mode=1&threshold=0&interval=5&mark=asdfasdfasdf&source=rtsp%3A%2F%2F192.168.1.160%3A10556%2Fplayback%3Fchannel%3D14%26starttime%3D1630295010%26endtime%3D1630295310%26isreduce%3D0"
	parse, err := url.Parse(origin)
	if err != nil {
		return
	}
	t.Log(parse.Query()["source"][0])

}
