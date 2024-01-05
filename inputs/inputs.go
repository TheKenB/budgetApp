package inputs

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextCollissionLocation struct {
	Location rl.Rectangle
	Text     *string
}

type ActiveText struct {
	Active bool
	Pos    [2]int
}

func TextInput(posX, posY float32, width, height, squares int) rl.Rectangle {
	var sWidth float32 = float32((width))*float32(squares) - 20
	var sHeight float32 = float32((height / 2))
	var input rl.Rectangle = rl.NewRectangle(posX, posY, float32(sWidth), float32(sHeight))
	return input
}

func Button(posX, posY float32, width, height, squares int) rl.Rectangle {
	var sWidth float32 = float32((width)) * float32(squares)
	var sHeight float32 = float32((height / 2))
	var input rl.Rectangle = rl.NewRectangle(posX, posY, float32(sWidth), float32(sHeight))
	return input
}
