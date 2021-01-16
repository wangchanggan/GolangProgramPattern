package GolangProgramPattern

import (
	"fmt"
	"testing"
)

func TestReduce(t *testing.T) {
	var list = []string{"Hao", "Chen", "MegaEase"}

	x := Reduce(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", x)
	// 15
}
