package grid

import (
	"math"
)

type DisplayGrid struct {
	Rows    int
	Columns int
	Width   int
	Height  int
}

func NewGrid(res [2]int) DisplayGrid {
	grid := DisplayGrid{
		Rows:    12,
		Columns: 12,
		Height:  int(math.Floor(float64(res[1]) / float64(12))),
		Width:   int(math.Floor(float64(res[0]) / float64(12))),
	}
	return grid
}

func GridPosXLeft(pos, width int) int {
	return pos * width
}

func GridPosTextXCent(pos, width int) int {
	return (pos * width) + width/4
}

func GridPosYTop(pos, height int) int {
	return pos * height
}

func GridPosYBot(pos, height int) int {
	return (pos * height) + height/2
}
