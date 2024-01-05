package banner

import (
	gr "main/grid"
	ui "main/inputs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawBanner(dGrid gr.DisplayGrid) {
	var height int = dGrid.Height
	var width int = dGrid.Width
	rl.DrawRectangleRec(ui.TextInput(float32(gr.GridPosXLeft(0, width)), float32(gr.GridPosYTop(0, height)), width, height, 13), rl.DarkGreen)
}
