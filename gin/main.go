package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!!")
	})
	// r.Any("/login", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "any hello world!!")
	// })

	// grp1 := r.Group("/v1")
	// {
	// 	grp1.GET("/user", func(c *gin.Context) {
	// 		c.String(http.StatusOK, "v1 get user")
	// 	})
	// }

	// grp2 := r.Group("/v2")
	// {
	// 	grp2.GET("/user", func(c *gin.Context) {
	// 		c.String(http.StatusOK, "v2 get user")
	// 	})
	// }

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
