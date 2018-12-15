package main

import (
	"fmt"
	"math/rand"
)

const (
	sides = 16
)

// This will draw a blue box
// https://goplay.space/#7G-boCrOBOF
func main() {
	degrees := float64(360)/float64(sides)

	fmt.Println("draw mode")

	for i := 0; i < sides; i++ {
		randomColor()
		right(degrees)
		fwd(1)
	}
}

// turn left n degrees
func left(degrees float64) {
	fmt.Printf("left %v\n", degrees)
}

// turn right n degrees
func right(degrees float64) {
	fmt.Printf("right %v\n", degrees)
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

// set a random color
func randomColor() {
	red := float64(rand.Intn(256))
	green := float64(rand.Intn(256))
	blue := float64(rand.Intn(256))

	colorRGBA(red, green, blue, 1)
}
