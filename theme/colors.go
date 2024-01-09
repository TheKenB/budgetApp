package colors

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var fontMajor rl.Font
var fontMinor rl.Font

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

// grey green
func MinorCColor() rl.Color {
	return rl.NewColor(213, 221, 213, 255)
}

func DrawMajor(text string, posX int32, posY int32, fontSize int32, col rl.Color) {
	rl.DrawTextEx(fontMajor, text, rl.Vector2{X: float32(posX), Y: float32(posY)}, float32(fontSize), 3, col)
}

func DrawMinor(text string, posX int32, posY int32, fontSize int32, col rl.Color) {
	rl.DrawTextEx(fontMinor, text, rl.Vector2{X: float32(posX), Y: float32(posY)}, float32(fontSize), 3, col)
}

func SetMajorFont() {
	fontMajor = rl.LoadFontEx("fonts/RobotoSlabMedium.ttf", 42, nil)
}

func SetMinorFont() {
	fontMinor = rl.LoadFontEx("fonts/RobotoSlabLight.ttf", 42, nil)
}
