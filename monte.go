// モンテカルロ法 Pure Go版

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num int
	fmt.Scan(&num)

	c := 0

	for i := 0; i < num; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if x*x+y*y <= 1.0 {
			c++
		}
	}

	var p float64
	p = 4.0 * float64(c) / float64(num)

	fmt.Printf("%f\n", p)
}
