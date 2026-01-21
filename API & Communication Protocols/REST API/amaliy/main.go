package main

import (
	"amaliy/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)




func main(){
 r := gin.Default()
 routes.RegisterUserRoutes(r)
 fmt.Println("server running 8081")
 r.Run(":8081")
}