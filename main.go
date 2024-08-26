package main

import (
	"fmt"

	test "github.com/dineshs90/demo2"
)

func main() {

	fmt.Println("Test")
	s1 := test.Bark()
	s2 := test.Barks()

	fmt.Println(s1)
	fmt.Println(s2)

}
