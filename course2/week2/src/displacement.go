package main

import (
	"fmt"
)

func GenDisplaceFn(a, vo, so float64) func(t float64) float64 {
	fn := func(t float64) float64 {
		return so + vo*t + ((0.5) * t * a * a)
	}
	return fn
}

func main() {

	var acceleration float64
	var initialVelocity float64
	var initialDisplacement float64
	fmt.Printf("\nenter the acceleration\n")
	fmt.Scan(&acceleration)
	fmt.Printf("\ninitial velocity\n")
	fmt.Scan(&initialVelocity)
	fmt.Printf("\ninitial displacement\n")
	fmt.Scan(&initialDisplacement)
	ComputeDisplacement := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	fmt.Printf("\ndistance travelled after 3 sec : %f", ComputeDisplacement(3))
	fmt.Printf("\ndistance travelled after 5 sec : %f", ComputeDisplacement(5))

	var time float64
	fmt.Printf("\nEnter time :")
	fmt.Scan(&time)
	fmt.Printf("\ndistance travelled after %f sec : %f", time, ComputeDisplacement(time))

}
