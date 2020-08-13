package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/inconsolata"
	"golang.org/x/image/math/fixed"
)

const (
	exitOK = iota
	exitFail
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// copy from https://stackoverflow.com/a/31832326
func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	for _, arg := range args {
		img := image.NewRGBA(image.Rect(0, 0, 300, 100))
		addLabel(img, 20, 30, arg)

		f, err := os.Create(fmt.Sprintf("out-%s.png", randStringRunes(8)))
		if err != nil {
			fmt.Println("failed to create output file: ", err)
			return exitFail
		}
		defer f.Close()

		if err := png.Encode(f, img); err != nil {
			fmt.Println("failed to save image: ", err)
			return exitFail
		}
	}

	return exitOK
}

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{R: 200, G: 100, B: 0, A: 255}
	point := fixed.Point26_6{
		X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: inconsolata.Regular8x16,
		Dot:  point,
	}
	d.DrawString(label)
}
