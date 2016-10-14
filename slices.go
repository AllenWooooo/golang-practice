package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)

	for i := range image {
		image[i] = make([]uint8, dx)

		for j := range image[i] {
			image[i][j] = uint8(i * j)
		}
	}

	return image
}

func main() {
	pic.Show(Pic)
}
