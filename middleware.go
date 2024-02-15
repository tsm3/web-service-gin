package main

import (
	"fmt"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func DummyMiddlewareSimple(c *gin.Context) {
	fmt.Println("Yuh")
	c.Next()
}

func DummyMiddlewareInitializer() gin.HandlerFunc {
	fmt.Println("This should only happen once")

	return func(c *gin.Context) {
		fmt.Println("This should happen each time.")
		c.Next()
	}
}
