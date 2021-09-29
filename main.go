package main

import "github.com/gin-gonic/gin"

func main(){
	router := gin.Default()
	router.GET("/", func(c *gin.Context){
		c.String(http.StatusOK, "Hello")
	})
	router.Run(":29999")
}