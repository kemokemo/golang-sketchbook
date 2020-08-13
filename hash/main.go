package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	exitCodeOK int = iota
	exitCodeFailed
)

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	f, err := os.Open(args[1])
	if err != nil {
		log.Println("failed to open file:", err)
		return exitCodeFailed
	}
	defer f.Close()

	hash, err := generateHash(f)
	if err != nil {
		log.Println("failed to generate hash:", err)
		return exitCodeFailed
	}
	fmt.Println(strings.ToUpper(hash))

	return exitCodeOK
}

func generateHash(r io.ReadCloser) (string, error) {
	// todo: If you want to use any other checksum kind, change this function.
	h := sha256.New()
	_, err := io.Copy(h, r)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
