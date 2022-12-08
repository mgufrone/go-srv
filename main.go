package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

const (
	defaultPort = 8000
)

func main() {
	srv := gin.Default()
	hello := &Hello{}
	srv.GET("/:name", func(c *gin.Context) {
		out := hello.Greet(c.Param("name"))
		c.String(http.StatusOK, out)
	})
	srv.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "OK")
	})
	srv.GET("/healthz", func(context *gin.Context) {
		context.String(http.StatusOK, "OK")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(defaultPort)
	}
	if err := srv.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}
