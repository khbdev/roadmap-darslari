package routes

import (
	"amaliy/model"
	"amaliy/storage"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func RegisterUserRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/users", GetUser)
		v1.POST("/users", PostUser)
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/users", GetUserV2)
	}
}

func GetUser(c *gin.Context){
	users := storage.GetUser() 
	c.JSON(http.StatusOK, users)
}

func PostUser(c *gin.Context){
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	created := storage.CreateUser(user)
	c.JSON(http.StatusOK, created)
}


func GetUserV2(c *gin.Context) {
	name := c.Query("name")
	limitStr := c.DefaultQuery("limit", "10")
	afterStr := c.DefaultPostForm("after_id", "0")


	limit, _ := strconv.Atoi(limitStr)
	afterID, _ := strconv.Atoi(afterStr)

	users, nextAfter := storage.GetUserFiltered(name, afterID, limit)
	c.JSON(http.StatusOK, gin.H{"users": users, "next_after_id": nextAfter})
}
