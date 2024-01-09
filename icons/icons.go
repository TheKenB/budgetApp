package icons

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var xIcon rl.Texture2D

func SetXIcon() {
	xIcon = rl.LoadTexture("images/xIcon.png")
}

func XButtonTexture() rl.Texture2D {
	return xIcon
}
