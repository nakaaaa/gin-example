package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Tests struct{}

func (t *Tests) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	}
}
