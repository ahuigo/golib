package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("log before")
		// 默认 Next
		c.Next()
		fmt.Println("log end")
	}
}
