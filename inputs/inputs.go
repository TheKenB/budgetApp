package inputs

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextCollissionLocation struct {
	Location rl.Rectangle
	Text     *string
}

type MenuTextCollissionLocation struct {
	Location rl.Rectangle
	List     []string
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

func HorizontalDivider(posX, posY float32, width, height, squares int) rl.Rectangle {
	var sWidth float32 = float32((width)) * float32(squares)
	var sHeight float32 = float32((height / 10))
	var input rl.Rectangle = rl.NewRectangle(posX, posY, float32(sWidth), float32(sHeight))
	return input
}

func HeaderDivider(posX, posY float32, width, height, squares int) rl.Rectangle {
	var sWidth float32 = float32((width)) / 14
	var sHeight float32 = float32((height / 2))
	var endX float32 = posX + (float32(width) * float32(squares)) - float32(width/10)
	var input rl.Rectangle = rl.NewRectangle(endX, posY, float32(sWidth), float32(sHeight))
	return input
}

func ResultLine(posX, posY float32, width, height, squares int) rl.Rectangle {
	var sWidth float32 = float32((width)) * float32(squares)
	var sHeight float32 = float32(height)
	var input rl.Rectangle = rl.NewRectangle(posX, posY, float32(sWidth), float32(sHeight))
	return input
}

func DropDown(posX, posY float32, width, height, squares int) rl.Rectangle {
	var sWidth float32 = float32((width)) * float32(squares)
	var menHeight float32 = float32(height * 6)
	var dropMen rl.Rectangle = rl.NewRectangle(posX, posY+float32(height), float32(sWidth), float32(menHeight))
	return dropMen
}
