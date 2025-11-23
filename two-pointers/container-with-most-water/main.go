package main

func main() {
}

func maxArea(height []int) int {
	// area= largura X altura
	// largura = iX - iY
	// altura min(height[iX], height[iY])

	// começa no primeiro e no segundo
	// calcula area, move segundo ponteiro até chegar no final calculando as áreas e mantendo a maior

	if len(height) <= 1 {
		return 0
	}

	var maxArea int
	var ptrA int = 0
	var ptrB int = len(height) - 1 // last element

	for ptrA != ptrB {
		width := ptrB - ptrA
		heightx := min(height[ptrB], height[ptrA])
		area := heightx * width

		if area > maxArea {
			maxArea = area
		}

		if height[ptrA] < height[ptrB] {
			ptrA++
		} else {
			ptrB--
		}
	}

	return maxArea
}
