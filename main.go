package main

import (
	"net/http"

	newdb "github.com/Nakaaaa/gin-example/db"
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

	db := newdb.NewDB()
	rt := route.Tests{}
	v1 := r.Group("/")
	{
		{
			tests := v1.Group("/tests")
			{
				tests.GET("", rt.Index())
				tests.GET("/user/:id", rt.GetUser(db))
				tests.POST("/user", rt.AddOrUpdate(db))
				tests.DELETE("/user/:id", rt.DeleteUser(db))
			}
		}
	}

	r.Run() // default port 8080
}
