package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
)

const (
	exitCodeOK int = iota
	exitCodeFailed
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	width := 300
	height := 300

	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Circle(width/2, height/2, 100)
	canvas.Text(width/2, height/2+5, "Hello, SVG!", "text-anchor:middle;font-size:30px;fill:white")
	canvas.End()

	return exitCodeOK
}
