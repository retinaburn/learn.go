package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	list := [][]uint8{}
	for i := 0; i < dy; i++ {
		row := []uint8{}
		for j := 0; j < dx; j++ {
			row = append(row, uint8(0))
		}
		list = append(list, row)
	}
	return list

}

func main() {
	list := Pic(256, 256)
	fmt.Println(list)
	pic.Show(Pic)
}
