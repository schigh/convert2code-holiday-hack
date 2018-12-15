package main

import (
	"fmt"
	"math/rand"
)

// This will draw a blue box
// https://goplay.space/#vmpdfY2M63M
func main() {
	fmt.Println("draw mode")
	color("blue")

	for i := 0; i < 20; i++ {
		randomTurn()
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

func randomTurn() {
	leftOrRight := rand.Intn(2) % 2
	degrees := float64(rand.Intn(360))

	if leftOrRight == 0 {
		right(degrees)
	} else {
		left(degrees)
	}
}
