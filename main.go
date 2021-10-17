package main

import (
	"fmt"
	"math/rand"
)

func main() {
	bins := Binning(10, 0, 1)

	idx := 0
	for {
		bins.Fill(rand.Float64())
		idx += 1

		if idx > 100 {
			fmt.Printf("bins: %#v\n", bins)
			fmt.Println("")
		}
	}
}
