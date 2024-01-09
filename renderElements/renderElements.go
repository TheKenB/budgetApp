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

// Draw back row backdrop
func DrawResultRowBackdrop(x, y, rowNum, squares, width, height int) {
	topRec := ui.Button(float32(gr.GridPosXLeft(x, width)), float32(gr.GridPosYTop(y+rowNum, height)), width, height, squares)
	botRec := ui.Button(float32(gr.GridPosXLeft(x, width)), float32(gr.GridPosYBot(y+rowNum, height)), width, height, squares)
	rl.DrawRectangleRec(topRec, rl.LightGray)
	rl.DrawRectangleRec(botRec, rl.Gray)
}

// Draw row result text
func DrawResultRowText(value1, value2 string, x, y, width, height, rowCount, fontSize int) {
	rl.DrawText(value1, int32(gr.GridPosXLeft(x, width)+5), int32(gr.GridPosYTop(y+rowCount, height)), int32(fontSize), rl.Black)
	rl.DrawText(value2, int32(gr.GridPosXLeft(x, width)+5), int32(gr.GridPosYBot(y+rowCount, height)), int32(fontSize), rl.Black)

}
