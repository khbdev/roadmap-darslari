package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

)



type  UserRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Age string `json:"age"`
}




func validateUser(req UserRequest) map[string]string {
	errors := make(map[string]string)

	if req.Name == "" {
		errors["name"] = "Name Is required"
	}
	if req.Email == "" {
		errors["email"] = "Email Is required"
	}
	if req.Age == "" {
		errors["age"] = "Age is required"
	}
	return  errors
}

func ValidationMiddleware() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		var reqs UserRequest
		if err := ctx.ShouldBindJSON(&reqs); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid Json"})
			ctx.Abort()
			return 
		}

		if err := validateUser(reqs); len(err) > 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": err})
			ctx.Abort()
			return 
		}
		ctx.Set("userRequest", reqs)
		ctx.Next()
	}
}


func LogginMIddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		path := ctx.Request.URL.Path
		fmt.Printf("Incoming request: %s %s\n", method, path)

		ctx.Next()

		status := ctx.Writer.Status()
		fmt.Printf("Completed request: %s %s with status %d\n", method, path, status)
	}
}

func main() {
	r := gin.Default()

	r.Use(LogginMIddleware())

	r.POST("/users", ValidationMiddleware(), func(ctx *gin.Context) {
		user := ctx.MustGet("userRequest").(UserRequest)
		ctx.JSON(http.StatusOK, gin.H{
		"Message" : "User created successfully",
		"user": user,
		})
	})

	r.Run(":8081")
}