package GolangProgramPattern

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	var intset = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := Filter(intset, func(n int) bool {
		return n%2 == 1
	})
	fmt.Printf("%v\n", out)
	//[1 3 5 7 9]

	out = Filter(intset, func(n int) bool {
		return n > 5
	})
	fmt.Printf("%v\n", out)
	//[6 7 8 9 10]
}
