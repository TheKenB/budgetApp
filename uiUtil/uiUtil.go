package uiUtil

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Return color mod for hover/no hover
func IsHoverRec(uiRec rl.Rectangle) bool {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), uiRec) {
		return true
	} else {
		return false
	}
}

func HoverBright(isHover bool) rl.Color {
	if isHover {
		return rl.RayWhite
	} else {
		return rl.Gray
	}
}
