package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	source := `abcd0123`
	result := columnToPrefix(source)
	fmt.Printf("%s is converted to %s\n", source, result)
}

func columnToPrefix(c string) (pre string) {
	reg := regexp.MustCompile(`([a-z]{4})(\d{4})`)
	result := reg.ReplaceAllString(c, "$1-$2")
	return strings.ToUpper(result)
}
