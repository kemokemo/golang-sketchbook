package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	exitCodeOK int = iota
	exitCodeFailed
)

var (
	port = flag.String("port", "9000", "the address of the API service")
)

func init() {
	flag.Parse()
}

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	err := r.Run(fmt.Sprintf(":%s", *port))
	if err != nil {
		log.Println("Failed to start a robot", err)
		return exitCodeFailed
	}
	return exitCodeOK
}
