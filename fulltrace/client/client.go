package main

import (
	"context"
	trace "fulltrace/protobuf"
	"time"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgin"
	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.DialContext(context.TODO(), "fulltrace-server:50051", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithUnaryInterceptor(apmgrpc.NewUnaryClientInterceptor()))

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	grpcFulltrace := trace.NewFullTraceServiceClient(conn)

	g := gin.Default()

	g.Use(apmgin.Middleware(g))

	g.GET("/", func(c *gin.Context) {

		span, ctx := apm.StartSpan(c.Request.Context(), "SELECT * from USER", "db")

		defer span.End()

		time.Sleep(time.Second)

		grpcFulltrace.HelloWorld(ctx, &trace.HelloWorldRequest{
			Name: "HELLO WORLD",
		})

		c.String(200, "ok")

		return

	})

	g.Run(":3000")

}
