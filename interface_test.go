package GolangProgramPattern

import "testing"

func TestInterface(t *testing.T) {
	d1 := Country{"USA"}
	d2 := City{"Los Angeles"}
	PrintStr(d1)
	//Country = USA
	PrintStr(d2)
	//City = Los Angeles
}
