package main

import (
	"fmt"
)

// turn left
func left() {
	fmt.Println("left")
}

// turn right
func right() {
	fmt.Println("right")
}
// This will draw a blue box
// incomplete: https://goplay.space/#MRQB1wvpVZs
// solution: https://goplay.space/#4afbasdPvvf
func main() {
	fmt.Println("draw mode")
	// turn left
	fmt.Println("color blue")
	fmt.Println("forward 5")
	// turn right
	fmt.Println("forward 10")
	// turn right
	fmt.Println("forward 10")
	// turn right
	fmt.Println("forward 10")
	// turn right
	fmt.Println("forward 5")
	fmt.Println("say done")
}
