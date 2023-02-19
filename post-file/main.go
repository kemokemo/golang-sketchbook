package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	ver bool
)

func init() {
	flag.BoolVar(&ver, "v", false, "display version")
	flag.Parse()
}

func main() {
	os.Exit(run())
}

func run() int {
	if ver {
		fmt.Printf("post-file version: %s.%s\n", Version, Revision)
		return 0
	}

	router := gin.Default()
	// Limit file size to 512 MB or less
	router.MaxMultipartMemory = 512 << 20

	router.POST("/upload", uploadPost)

	err := router.Run(":8080")
	if err != nil {
		log.Println("failed to run server, ", err)
		return 1
	}

	return 0
}
