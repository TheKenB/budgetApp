package pageEntry

import (
	gr "main/grid"
	ui "main/inputs"
	enJson "main/json"
	rendEl "main/renderElements"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var xButton rl.Texture2D

func LoadTexture() {
	xButton = rl.LoadTexture("images/xIcon.png")
}

func HandleEntryPageResults(dGrid gr.DisplayGrid, records []enJson.Entries) {
	var height int = dGrid.Height
	var width int = dGrid.Width
	elapsedTime += rl.GetFrameTime()
	// Draw Header
	//Description Header
	rendEl.DrawResultHeader(1, 5, width, height, 3, "Description", false)

	//Amount Header
	rendEl.DrawResultHeader(4, 5, width, height, 2, "Amount", false)

	//Date Header
	rendEl.DrawResultHeader(6, 5, width, height, 2, "Date", false)

	//Action Header
	rendEl.DrawResultHeader(8, 5, width, height, 2, "Actions", true)

	var headerDivider rl.Rectangle = ui.HorizontalDivider(float32(gr.GridPosXLeft(1, width)), float32(gr.GridPosYTop(6, height-1)), width, height, 9)
	rl.DrawRectangleRec(headerDivider, rl.DarkGreen)

	for i := 0; i < 6; i++ {
		var resultDescTop rl.Rectangle = ui.Button(float32(gr.GridPosXLeft(1, width)), float32(gr.GridPosYTop(6+i, height)), width, height, 3)
		var resultAmtTop rl.Rectangle = ui.Button(float32(gr.GridPosXLeft(4, width)), float32(gr.GridPosYTop(6+i, height)), width, height, 2)
		var resultDateTop rl.Rectangle = ui.Button(float32(gr.GridPosXLeft(6, width)), float32(gr.GridPosYTop(6+i, height)), width, height, 2)
		var resultDescBot rl.Rectangle = ui.Button(float32(gr.GridPosXLeft(1, width)), float32(gr.GridPosYBot(6+i, height)), width, height, 3)
		var resultAmtBot rl.Rectangle = ui.Button(float32(gr.GridPosXLeft(4, width)), float32(gr.GridPosYBot(6+i, height)), width, height, 2)
		var resultDateBot rl.Rectangle = ui.Button(float32(gr.GridPosXLeft(6, width)), float32(gr.GridPosYBot(6+i, height)), width, height, 2)
		var resultActTop rl.Rectangle = ui.Button(float32(gr.GridPosXLeft(8, width)), float32(gr.GridPosYTop(6+i, height)), width, height, 2)
		var resultActBot rl.Rectangle = ui.Button(float32(gr.GridPosXLeft(8, width)), float32(gr.GridPosYBot(6+i, height)), width, height, 2)
		rl.DrawRectangleRec(resultDescBot, rl.LightGray)
		rl.DrawRectangleRec(resultAmtBot, rl.LightGray)
		rl.DrawRectangleRec(resultDateBot, rl.LightGray)
		rl.DrawRectangleRec(resultDescTop, rl.DarkGray)
		rl.DrawRectangleRec(resultAmtTop, rl.DarkGray)
		rl.DrawRectangleRec(resultDateTop, rl.DarkGray)
		rl.DrawRectangleRec(resultActTop, rl.DarkGray)
		rl.DrawRectangleRec(resultActBot, rl.LightGray)
	}

	var rowCount = 0
	for j := 0; j <= 12; j += 2 {
		if j > len(records)-1 {
			break
		}
		var resultActTop rl.Rectangle = ui.Button(float32(gr.GridPosXLeft(8, width)), float32(gr.GridPosYTop(6+rowCount, height)), width, height, 2)
		var resultActBot rl.Rectangle = ui.Button(float32(gr.GridPosXLeft(8, width)), float32(gr.GridPosYBot(6+rowCount, height)), width, height, 2)
		strAmt1 := strconv.FormatFloat(float64(records[j].Amount), 'f', -1, 32)
		rl.DrawText(records[j].Description, int32(gr.GridPosXLeft(1, width)+5), int32(gr.GridPosYTop(6+rowCount, height)), 28, rl.Black)
		rl.DrawText(strAmt1, int32(gr.GridPosXLeft(4, width))+5, int32(gr.GridPosYTop(6+rowCount, height)), 28, rl.Black)
		rl.DrawText(records[j].Date, int32(gr.GridPosXLeft(6, width))+5, int32(gr.GridPosYTop(6+rowCount, height)), 28, rl.Black)
		//rl.DrawTextureEx(xButton, rl.Vector2{X: resultActTop.X, Y: resultActBot.Y}, 0, 100, rl.RayWhite)

		if j+1 <= len(records)-1 {
			strAmt2 := strconv.FormatFloat(float64(records[j+1].Amount), 'f', -1, 32)
			rl.DrawText(records[j+1].Description, int32(gr.GridPosXLeft(1, width)+5), int32(gr.GridPosYBot(6+rowCount, height)), 28, rl.Black)
			rl.DrawText(strAmt2, int32(gr.GridPosXLeft(4, width))+5, int32(gr.GridPosYBot(6+rowCount, height)), 28, rl.Black)
			rl.DrawText(records[j+1].Date, int32(gr.GridPosXLeft(6, width))+5, int32(gr.GridPosYBot(6+rowCount, height)), 28, rl.Black)
			rl.DrawTextureRec(xButton, resultActBot, rl.Vector2{X: resultActTop.X, Y: resultActBot.Y}, rl.RayWhite)
		}
		rowCount++
	}
}
