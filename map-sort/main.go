package main

import (
	"fmt"
	"sort"
)

//Food is food struct to eat. Which one do you like?
type Food struct {
	Name  string
	Price int
}

func main() {
	foods := make([]Food, 4)
	foods[0] = Food{Name: "みかん", Price: 150}
	foods[1] = Food{Name: "バナナ", Price: 100}
	foods[2] = Food{Name: "りんご", Price: 120}
	foods[3] = Food{Name: "ぶどう", Price: 200}

	sort.Slice(foods, func(i, j int) bool {
		return foods[i].Price < foods[j].Price
	})

	for _, food := range foods {
		fmt.Printf("%+v\n", food)
	}
}
