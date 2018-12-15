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

// use a color by name or hex code
func color(name string) {
	fmt.Printf("color %s\n", name)
}

// use a color by its red, green, blue, and alpha channels
func colorRGBA(red, green, blue, alpha float64) {
	fmt.Printf("color rgba(%v,%v,%v,%v)\n", red, green, blue, alpha)
}

// This will draw a blue box
// solution: https://goplay.space/#qAWo1IIPSDn
func main() {
	fmt.Println("draw mode")
	left()
	// set color to 'blue'
	fwd(5)
	right()
	// set color to '#ffcc33'
	fwd(10)
	right()
	// use colorRGBA to set color to purple
	fwd(10)
	right()
	// set color to 'red'
	fwd(10)
	right()
	// set color to '#0000ff'
	fwd(5)
	fmt.Println("say done")
}
