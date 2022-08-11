package main

import (
	"net/http"

	"github.com/Nakaaaa/gin-example/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	rt := route.Tests{}
	v1 := r.Group("/")
	{
		{
			tests := v1.Group("/tests")
			{
				tests.GET("", rt.Index())
			}
		}
	}

	r.Run() // default port 8080
}
