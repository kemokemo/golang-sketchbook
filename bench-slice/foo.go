package foo

// copy from 'https://mattn.kaoriya.net/software/lang/go/20161019124907.htm'

import "fmt"

func makeSlice(n int) []string {
	var r []string
	for i := 0; i < n; i++ {
		r = append(r, fmt.Sprintf("%03d だよーん", i))
	}
	return r
}
