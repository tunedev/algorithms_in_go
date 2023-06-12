package main

import "math"

func main() {

}

// Problem Statement
//    Given two crystal balls that will break if dropped from high enough distance,
//    Determine the exact spot in which it will break in the most optimized way

func findOptimallyDistanceWhereTwoCrystalBallsBrks(breaks []bool) int {
	n := len(breaks)
	jumpedDistance := int(math.Floor(math.Sqrt(float64(n))))

	i := jumpedDistance

	for ; i < n; i += jumpedDistance {
		if breaks[i] {
			break
		}
	}

	i -= jumpedDistance

	for j := 0; j < jumpedDistance && i < n; j++ {
		i++
		if breaks[i] {
			return i
		}
	}
	return -1
}
