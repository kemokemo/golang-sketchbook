package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	err := json.NewEncoder(os.Stdout).Encode(fmt.Errorf("failed to execute").Error())
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to encode an error to json: ", err)
		return 1
	}

	err = json.NewEncoder(os.Stdout).Encode(
		struct {
			Result interface{} `json:"Result"`
			Error  error       `json:"Error"`
		}{"hi!", fmt.Errorf("failed to execute")})
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to encode an error to json: ", err)
		return 1
	}

	return 0
}
