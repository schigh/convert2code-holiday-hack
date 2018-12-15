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

// go forward
func fwd(distance float64) {
	fmt.Printf("forward %v\n", distance)
}

// This will draw a blue box
// solution: https://goplay.space/#AnxJDtq-H4X
func main() {
	fmt.Println("draw mode")
	left()
	fmt.Println("color blue")
	// go forward 5
	right()
	// go forward 10
	right()
	// go forward 10
	right()
	// go forward 10
	right()
	// go forward 5
	fmt.Println("say done")
}
