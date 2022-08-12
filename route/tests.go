package route

import (
	"net/http"
	"strconv"

	"github.com/Nakaaaa/gin-example/gateway"
	"github.com/Nakaaaa/gin-example/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Tests struct{}

func (t *Tests) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	}
}

func (t *Tests) GetUser(db *gorm.DB) gin.HandlerFunc {
	rp := gateway.NewUserRepository()

	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}
		user, err := rp.GetByUserID(db, id)
		if err != nil {
			panic(err)
		}
		if user == nil {
			c.JSON(http.StatusNotFound, "User Not Found")
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func (t *Tests) AddOrUpdate(db *gorm.DB) gin.HandlerFunc {
	rp := gateway.NewUserRepository()
	model := &model.User{}

	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(&model); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		user, err := rp.AddOrUpdate(db, model)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, user)
	}
}

func (t *Tests) DeleteUser(db *gorm.DB) gin.HandlerFunc {
	rp := gateway.NewUserRepository()

	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}
		err = rp.DeleteUser(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, "User Not Found")
			return
		}

		c.JSON(http.StatusOK, "deleted user")
	}
}
