package main

import "fmt"

const (
	sides = 4
)

// This will draw a blue box
// https://goplay.space/#mjA9HADYTQd
func main() {
	degrees := float64(360)/float64(sides)

	fmt.Println("draw mode")
	color("blue")

	for i := 0; i < sides; i++ {
		right(degrees)
		fwd(5)
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
