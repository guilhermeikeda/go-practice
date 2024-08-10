package main

func main() {
	gain := []int{-5, 1, 5, 0, -7}
	largestAltitudeAlt(gain)
}

// 43 min
func largestAltitudeAlt(gain []int) int {
	highest := 0

	acc := 0
	for _, currGain := range gain {
		acc += currGain
		if acc > highest {
			highest = acc
		}
	}
	return highest
}

func largestAltitude(gain []int) int {
	highest := 0

	points := []int{0}
	for i, currGain := range gain {
		point := points[i] + currGain
		points = append(points, point)
		if point > highest {
			highest = point
		}
	}

	return highest
}
