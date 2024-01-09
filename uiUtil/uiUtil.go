package uiUtil

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Return color mod for hover/no hover
func IsHoverRec(uiRec rl.Rectangle) rl.Color {
	fmt.Println(uiRec)
	fmt.Println(rl.GetMousePosition())
	fmt.Println("----")
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), uiRec) {

		return rl.DarkGray
	} else {
		return rl.RayWhite
	}
}
