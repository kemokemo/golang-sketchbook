package main

import (
	"gozip/internal"
	"log"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	archiver := internal.ZIP
	dirName := "test-dir"

	err := archiver.Archive(dirName, archiver.DestFmt()(dirName))
	if err != nil {
		log.Println("failed to archive to zip: ", err)
		return 1
	}

	return 0
}
