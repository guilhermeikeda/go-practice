package main

import "fmt"

type Reader struct {
	teste string
}

func (r *Reader) Close() {
	r.teste = fmt.Sprintf("fechado: %s", r.teste)
}

func main() {
	reader := &Reader{teste: "teste"}
	fmt.Printf("%p\n", reader)
	*reader = Reader{teste: "novo teste"}
	fmt.Printf("%p\n", reader)
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	freeSpots := 0

	flowerbedSize := len(flowerbed)

	for i, slot := range flowerbed {
		if slot == 1 { // filled already
			continue
		}

		prevIndex := i - 1
		prevIsEmpty := (prevIndex < 0) || isEmpty(flowerbed, prevIndex)

		nextIndex := i + 1
		nextIsEmpty := (nextIndex >= flowerbedSize) || isEmpty(flowerbed, nextIndex)

		if prevIsEmpty && nextIsEmpty {
			flowerbed[i] = 1
			freeSpots++
		}
	}

	return freeSpots >= n
}

func isEmpty(flowerbed []int, position int) bool {
	return flowerbed[position] == 0
}
