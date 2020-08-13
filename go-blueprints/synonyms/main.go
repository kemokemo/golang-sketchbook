package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/kemokemo/thesaurus"
)

type exitCode int

// exit code of this tool.
const (
	exitOK exitCode = iota
	exitFailedToGet
	exitNotFound
	exitEnvNotSet
)

func main() {
	os.Exit(int(run()))
}

func run() exitCode {
	apiKey := os.Getenv("BHT_APIKEY")
	if apiKey == "" {
		log.Println("can not load the BHT_APIKEY env")
		return exitEnvNotSet
	}

	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Printf("failed to get the synonyms of the %q: %v\n", word, err)
			return exitFailedToGet
		}
		if len(syns) == 0 {
			log.Printf("the synonyms of the %q was not found\n", word)
			return exitNotFound
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}

	return exitOK
}
