package pageEntry

import (
	gr "main/grid"
	ui "main/inputs"
	enJson "main/json"
	rendEl "main/renderElements"
	color "main/theme"
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
	rl.DrawRectangleRec(headerDivider, color.MinorAColor())

	// Draw row backdrop
	for i := 0; i < 6; i++ {
		rendEl.DrawResultRowBackdrop(1, 6, i, 3, width, height)
		rendEl.DrawResultRowBackdrop(4, 6, i, 2, width, height)
		rendEl.DrawResultRowBackdrop(6, 6, i, 2, width, height)
		rendEl.DrawResultRowBackdrop(8, 6, i, 2, width, height)
	}

	// Draw Row Text Value
	var rowCount = 0
	for j := 0; j <= 12; j += 2 {
		if j > len(records)-1 {
			break
		}
		strAmt1 := strconv.FormatFloat(float64(records[j].Amount), 'f', -1, 32)
		strAmt2 := ""
		if j+1 <= len(records)-1 {
			strAmt2 = strconv.FormatFloat(float64(records[j+1].Amount), 'f', -1, 32)
		}
		rendEl.DrawResultRowText(records[j].Description, strAmt2, 1, 6, width, height, rowCount, 28)
		rendEl.DrawResultRowText(strAmt1, strAmt2, 4, 6, width, height, rowCount, 28)
		rendEl.DrawResultRowText(records[j].Date, strAmt2, 6, 6, width, height, rowCount, 28)
		rowCount++
	}
}
