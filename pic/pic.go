package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for x := range pic {
		pic[x] = make([]uint8, dx)
	}

	for y, row := range pic {
		for x := range row {
			row[x] = uint8((x * y) / (x + y))
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
