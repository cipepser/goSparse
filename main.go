package main

import (
	"fmt"

	"github.com/james-bowman/sparse"
)

func main() {
	r := 4
	c := 4
	data := []float64{6, 5, 3, 1, 4, 2}

	ia := []int{3, 2, 1, 0, 2, 1}
	ja := []int{0, 2, 3, 2, 1, 1}

	A := sparse.NewCOO(r, c, ia, ja, data)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Printf("%.0f", A.At(i, j))
			if j < c-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}

}
