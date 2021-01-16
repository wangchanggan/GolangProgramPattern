package GolangProgramPattern

import "testing"

func TestDecorator(t *testing.T) {
	decorator(Hello)("Hello, World!")

	/*Started
	Hello, World!
	Done*/
}
