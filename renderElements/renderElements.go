package renderElements

import (
	gr "main/grid"
	ui "main/inputs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Draw Inputs
func DrawInputs(headerRec ui.TextCollissionLocation, err, title string) {

	x := headerRec.Location.X
	y := headerRec.Location.Y
	text := headerRec.Text

	rl.DrawText(title, int32(x), int32(y-headerRec.Location.Height), 32, rl.DarkGreen)
	rl.DrawRectangleRec(headerRec.Location, rl.LightGray)
	rl.DrawText(*text, int32(x+15), int32(y), 28, rl.Black)
	rl.DrawText(err, int32(x+10), int32(y+1), 24, rl.Red)
}

// Draw table headers
func DrawResultHeader(x, y, width, height, squares int, title string, end bool) {
	var resultText rl.Rectangle = ui.Button(float32(gr.GridPosXLeft(x, width)), float32(gr.GridPosYBot(y, height-1)), width, height, squares)
	var columnDiv rl.Rectangle = ui.HeaderDivider(float32(gr.GridPosXLeft(x, width)), float32(gr.GridPosYBot(y, height-1)), width, height, squares)
	rl.DrawRectangleRec(resultText, rl.DarkGray)
	if !end {
		rl.DrawRectangleRec(columnDiv, rl.DarkGreen)
	}
	rl.DrawText(title, int32(resultText.X+10), int32(resultText.Y), 28, rl.White)
}
