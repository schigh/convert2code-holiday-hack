package main

import (
	"fmt"
	"math"
)

const (
	TargetX = -5
	TargetY = -5
)

// https://goplay.space/#53f1iEF2a-l
func main() {
	fmt.Println("draw mode")
	fmt.Println("color orange")
	fmt.Println("width 4")

	// calculate the slope between the points represented by 0,0 (the origin) and TargetX, TargetY
	// https://www.chilimath.com/lessons/intermediate-algebra/slope-formula/
	// slope is calculated as: (x2 - x1) / (y2 - y1)
	// x1 and y1 are always zero, so x2 - x1 is always x2, and y2 - y1 is always y2
	slope := float64(TargetX) / float64(TargetY)
	fmt.Printf("slope: %v\n", slope)

	// now get the distance between the two points. this is a little trickier
	// https://www.chilimath.com/lessons/intermediate-algebra/distance-formula/
	// distance is defined as √(x2 - x1)^2 + (y2 - y1)^2
	// x1 and y1 are always zero, so x2 - x1 is always x2, and y2 - y1 is always y2
	distance := math.Sqrt(float64(TargetX*TargetX) + float64(TargetY*TargetY))
	fmt.Printf("distance: %v\n", distance)

	// calculate angle
	angle := math.Atan2(float64(TargetX), float64(TargetY)) * (180 / math.Pi)

	if angle < 0 {
		fmt.Printf("left %v\n", math.Abs(angle))
	} else {
		fmt.Printf("right %v\n", angle)
	}
	fmt.Printf("forward %v\n", distance)
}
