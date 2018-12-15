package main

import "fmt"

const (
	TargetX = 10
	TargetY = 10
)

type direction int

const (
	North direction = iota
	East
	South
	West
)

type position struct {
	X int
	Y int
	Heading direction
}

// the gopher's position
var gopherPosition position

// https://goplay.space/#XoutAIuKl4j
func main() {
	fmt.Println("draw mode")
	fmt.Println("color blue")

	for {
		xDifference := TargetX - gopherPosition.X
		yDifference := TargetY - gopherPosition.Y

		// if our gopher is at their destination, we're done!
		if xDifference == 0 && yDifference == 0 {
			break
		}

		// if the X difference is less than 0, we need to go west
		// if it is greater than 0, we need to go east
		// if it's zero, we don't move west or east!
		if xDifference < 0 {
			faceWest()
			fmt.Println("forward 1")
			gopherPosition.X--
		} else if xDifference > 0 {
			faceEast()
			fmt.Println("forward 1")
			gopherPosition.X++
		}

		// if the Y difference is less than 0, we need to go south
		// if it is greater than 0, we need to go north
		// if it's zero, we don't move north or south!
		if yDifference < 0 {
			faceSouth()
			fmt.Println("forward 1")
			gopherPosition.Y--
		} else if yDifference > 0 {
			faceNorth()
			fmt.Println("forward 1")
			gopherPosition.Y++
		}
	}
}

// point the gopher north
func faceNorth() {
	switch gopherPosition.Heading {
	case South:
		fmt.Println("right 180")
	case East:
		fmt.Println("left 90")
	case West:
		fmt.Println("right 90")
	}
	gopherPosition.Heading = North
}

// point the gopher east
func faceEast() {
	switch gopherPosition.Heading {
	case West:
		fmt.Println("right 180")
	case South:
		fmt.Println("left 90")
	case North:
		fmt.Println("right 90")
	}
	gopherPosition.Heading = East
}

// point the gopher south
func faceSouth() {
	switch gopherPosition.Heading {
	case North:
		fmt.Println("right 180")
	case West:
		fmt.Println("left 90")
	case East:
		fmt.Println("right 90")
	}
	gopherPosition.Heading = South
}

// point the goper west
func faceWest() {
	switch gopherPosition.Heading {
	case East:
		fmt.Println("right 180")
	case North:
		fmt.Println("left 90")
	case South:
		fmt.Println("right 90")
	}
	gopherPosition.Heading = West
}
