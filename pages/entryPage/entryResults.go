package pageEntry

import (
	gr "main/grid"
	icons "main/icons"
	ui "main/inputs"
	enJson "main/json"
	rendEl "main/renderElements"
	color "main/theme"
	"main/uiUtil"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var xButton rl.Texture2D

func LoadTexture() {
	xButton = rl.LoadTexture("images/xIcon.png")
}

func HandleEntryPageResults(dGrid gr.DisplayGrid, records []enJson.Entries) bool {
	var height int = dGrid.Height
	var width int = dGrid.Width
	var listChange bool = false
	elapsedTime += rl.GetFrameTime()
	// Draw Header
	//Description Header
	rendEl.DrawResultHeader(1, 4, width, height, 3, "Description", false)

	//Amount Header
	rendEl.DrawResultHeader(4, 4, width, height, 2, "Amount", false)

	//Date Header
	rendEl.DrawResultHeader(6, 4, width, height, 2, "Date", false)

	//Action Header
	rendEl.DrawResultHeader(8, 4, width, height, 2, "Actions", true)

	var headerDivider rl.Rectangle = ui.HorizontalDivider(float32(gr.GridPosXLeft(1, width)), float32(gr.GridPosYTop(5, height-1)), width, height, 9)
	rl.DrawRectangleRec(headerDivider, color.MinorAColor())

	// Draw row backdrop
	for i := 0; i < 6; i++ {
		rendEl.DrawResultRowBackdrop(1, 5, i, 3, width, height)
		rendEl.DrawResultRowBackdrop(4, 5, i, 2, width, height)
		rendEl.DrawResultRowBackdrop(6, 5, i, 2, width, height)
		rendEl.DrawResultRowBackdrop(8, 5, i, 2, width, height)
	}

	// Draw Row Text Value
	var rowCount = 0
	for j := 0; j <= 12; j += 2 {
		if j > len(records)-1 {
			break
		}
		strAmt1 := strconv.FormatFloat(float64(records[j].Amount), 'f', -1, 32)
		strAmt2 := ""
		strDesc2 := ""
		strDate2 := ""
		doubleDraw := false
		if j+1 <= len(records)-1 {
			strAmt2 = strconv.FormatFloat(float64(records[j+1].Amount), 'f', -1, 32)
			strDesc2 = records[j+1].Description
			strDate2 = records[j+1].Date
			doubleDraw = true
		}
		rendEl.DrawResultRowText(records[j].Description, strDesc2, 1, 5, width, height, rowCount, 28, doubleDraw)
		rendEl.DrawResultRowText(strAmt1, strAmt2, 4, 5, width, height, rowCount, 28, doubleDraw)
		rendEl.DrawResultRowText(records[j].Date, strDate2, 6, 5, width, height, rowCount, 28, doubleDraw)

		xButton := icons.XButtonTexture()

		rendEl.DrawResultAction(xButton, 8, 5, width, height, rowCount, doubleDraw)
		posVectorTop := rl.Vector2{X: float32(gr.GridPosXLeft(8, width) + 5), Y: float32(gr.GridPosYTop(5+rowCount, height) + 5)}
		posVectorBot := rl.Vector2{X: float32(gr.GridPosXLeft(8, width) + 5), Y: float32(gr.GridPosYBot(5+rowCount, height) + 5)}
		listChange = DeleteRecordCheck(width, height, rowCount, posVectorTop, records[j])
		if !listChange && doubleDraw {
			listChange = DeleteRecordCheck(width, height, rowCount, posVectorBot, records[j+1])
		}
		if listChange {
			return listChange
		}
		rowCount++
	}
	return listChange
}

func DeleteRecordCheck(width, height, rowCount int, vectorLoc rl.Vector2, record enJson.Entries) bool {
	recDeleted := false
	topHoverState := uiUtil.IsHoverRec(rl.Rectangle{X: vectorLoc.X, Y: vectorLoc.Y, Width: float32(width / 4), Height: float32(height / 2)})
	if topHoverState && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			enJson.DeleteEntry(record)
			recDeleted = true
		}
	}
	return recDeleted
}
