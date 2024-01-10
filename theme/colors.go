package colors

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var FontMajor rl.Font
var FontMinor rl.Font

// Dark Green
func PrimaryColor() rl.Color {
	return rl.NewColor(27, 39, 39, 255)
}

// Muted Dark Green
func SecondaryColor() rl.Color {
	return rl.NewColor(60, 81, 72, 255)
}

// Green
func MinorAColor() rl.Color {
	return rl.NewColor(107, 142, 78, 255)
}

// Light Green
func MinorBColor() rl.Color {
	return rl.NewColor(178, 197, 178, 255)
}

// Gray green
func MinorCColor() rl.Color {
	return rl.NewColor(213, 221, 213, 255)
}

// Red danger
func DangerColor() rl.Color {
	return rl.NewColor(204, 51, 0, 255)
}

func DrawMajorText(text string, posX int32, posY int32, fontSize int32, col rl.Color) {
	rl.DrawTextEx(FontMajor, text, rl.Vector2{X: float32(posX), Y: float32(posY)}, float32(fontSize), 3, col)
}

func DrawMinorText(text string, posX int32, posY int32, fontSize int32, col rl.Color) {
	rl.DrawTextEx(FontMinor, text, rl.Vector2{X: float32(posX), Y: float32(posY)}, float32(fontSize), 3, col)
}

func SetMajorFont() {
	FontMajor = rl.LoadFontEx("fonts/RobotoSlabMedium.ttf", 42, nil)
}

func SetMinorFont() {
	FontMinor = rl.LoadFontEx("fonts/RobotoSlabLight.ttf", 42, nil)
}
