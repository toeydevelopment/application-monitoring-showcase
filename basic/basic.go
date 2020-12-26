package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func main() {

	g := gin.Default()

	g.Use(apmgin.Middleware(g))

	g.GET("/", func(c *gin.Context) {

		needError := c.Query("error")

		if needError != "" {
			c.String(500, "internal server error")
			return
		}

		c.String(200, "success")

	})

	g.GET("/users", func(c *gin.Context) {
		randomDelay := rand.Intn(4)

		time.Sleep(time.Duration(randomDelay*100) * time.Millisecond)

		c.String(200, "users")
	})

	g.Run(":3000")
}
