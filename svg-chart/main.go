package main

import (
	"fmt"
	"os"

	chart "github.com/wcharczuk/go-chart"
)

const (
	exitCodeOK int = iota
	exitCodeFailed
)

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:      "The XAxis",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "The YAxis",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name:    "A test series",
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{0.1, 0.2, 1.0, 0.3, 0.0},
			},
		},
	}

	// this is necessary to render the chart.ContinuousSeries's name.
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}

	f, err := os.Create("out.svg")
	if err != nil {
		fmt.Println("failed to create out.svg: ", err)
		return exitCodeFailed
	}
	defer f.Close()

	err = graph.Render(chart.SVG, f)
	if err != nil {
		fmt.Println("failed to render the graph: ", err)
		return exitCodeFailed
	}

	return exitCodeOK
}
