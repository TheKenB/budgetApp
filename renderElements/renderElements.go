package renderElements

import (
	gr "main/grid"
	ui "main/inputs"
	color "main/theme"
	uiUtil "main/uiUtil"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Draw Inputs
func DrawInputs(headerRec ui.TextCollissionLocation, title string) {

	x := headerRec.Location.X
	y := headerRec.Location.Y
	text := headerRec.Text
	color.DrawMajorText(title, int32(x), int32(y-headerRec.Location.Height), 32, color.MinorAColor())
	rl.DrawRectangleRec(headerRec.Location, color.MinorCColor())
	color.DrawMinorText(*text, int32(x+15), int32(y), 28, rl.Black)
}

// Draw Input Error
func DrawInputErr(x, y, width, height int, err string, col rl.Color, lowPos bool) {
	if !lowPos {
		color.DrawMajorText(err, int32(gr.GridPosXLeft(x, width)), int32(gr.GridPosYTop(y, height)), 22, col)
	} else {
		color.DrawMajorText(err, int32(gr.GridPosXLeft(x, width)), int32(gr.GridPosYBot(y, height)), 22, col)
	}
}

// Draw table headers
func DrawResultHeader(x, y, width, height, squares int, title string, end bool) {
	var resultText rl.Rectangle = ui.Button(float32(gr.GridPosXLeft(x, width)), float32(gr.GridPosYBot(y, height-1)), width, height, squares)
	var columnDiv rl.Rectangle = ui.HeaderDivider(float32(gr.GridPosXLeft(x, width)), float32(gr.GridPosYBot(y, height-1)), width, height, squares)
	rl.DrawRectangleRec(resultText, color.SecondaryColor())
	if !end {
		rl.DrawRectangleRec(columnDiv, color.MinorAColor())
	}
	color.DrawMajorText(title, int32(resultText.X+10), int32(resultText.Y), 28, color.MinorCColor())
}

// Draw back row backdrop
func DrawResultRowBackdrop(x, y, rowNum, squares, width, height int) {
	topRec := ui.Button(float32(gr.GridPosXLeft(x, width)), float32(gr.GridPosYTop(y+rowNum, height)), width, height, squares)
	botRec := ui.Button(float32(gr.GridPosXLeft(x, width)), float32(gr.GridPosYBot(y+rowNum, height)), width, height, squares)
	rl.DrawRectangleRec(topRec, color.MinorCColor())
	rl.DrawRectangleRec(botRec, color.MinorBColor())
}

// Draw row result text
func DrawResultRowText(value1, value2 string, x, y, width, height, rowCount, fontSize int, doubleDraw bool) {
	color.DrawMinorText(value1, int32(gr.GridPosXLeft(x, width)+5), int32(gr.GridPosYTop(y+rowCount, height)), int32(fontSize), rl.Black)
	if doubleDraw {
		color.DrawMinorText(value2, int32(gr.GridPosXLeft(x, width)+5), int32(gr.GridPosYBot(y+rowCount, height)), int32(fontSize), rl.Black)
	}
}

// Draw row action button
func DrawResultAction(icon rl.Texture2D, x, y, width, height, rowCount int, doubleDraw bool) {

	posVectorTop := rl.Vector2{X: float32(gr.GridPosXLeft(x, width) + 5), Y: float32(gr.GridPosYTop(y+rowCount, height) + 5)}
	posVectorBot := rl.Vector2{X: float32(gr.GridPosXLeft(x, width) + 5), Y: float32(gr.GridPosYBot(y+rowCount, height) + 5)}
	topHoverState := uiUtil.IsHoverRec(rl.Rectangle{X: posVectorTop.X, Y: posVectorTop.Y, Width: float32(width / 4), Height: float32(height / 2)})
	botHoverState := uiUtil.IsHoverRec(rl.Rectangle{X: posVectorBot.X, Y: posVectorBot.Y, Width: float32(width / 4), Height: float32(height / 2)})
	rl.DrawTextureEx(icon, posVectorTop, 0, 0.5, uiUtil.HoverBright(topHoverState))
	if doubleDraw {
		rl.DrawTextureEx(icon, posVectorBot, 0, 0.5, uiUtil.HoverBright(botHoverState))
	}
}
