package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
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

		findMockDB(c.Request.Context())

		c.String(200, "users")
	})

	g.Run(":3000")
}

func findMockDB(ctx context.Context) {
	span, _ := apm.StartSpan(ctx, "SELECT * FROM USER", "db")

	defer span.End()

	randomDelay := rand.Intn(4)

	time.Sleep(time.Duration(randomDelay*100) * time.Millisecond)

}
