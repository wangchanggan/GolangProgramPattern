package GolangProgramPattern

import (
	"fmt"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	var list = []string{"Hao", "Chen", "MegaEase"}

	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Printf("%v\n", x)
	//["HAO", "CHEN", "MEGAEASE"]

	y := MapStrToInt(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", y)
	//[3, 4, 8]
}
