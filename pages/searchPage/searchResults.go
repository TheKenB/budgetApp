package searchPage

import (
	gr "main/grid"
	"main/icons"
	ui "main/inputs"
	enJson "main/json"
	rendEl "main/renderElements"
	color "main/theme"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// icons "main/icons"

// "main/uiUtil"

func HandleSearchResults(dGrid gr.DisplayGrid, records []enJson.Entries) bool {
	var height int = dGrid.Height
	var width int = dGrid.Width
	var listChange bool = false
	// Draw Header
	//Description Header
	rendEl.DrawResultHeader(1, 4, width, height, 2, "Description", false)

	//Bucket Header
	rendEl.DrawResultHeader(3, 4, width, height, 2, "Buckets", false)

	//Amount Header
	rendEl.DrawResultHeader(5, 4, width, height, 2, "Amount", false)

	//Date Header
	rendEl.DrawResultHeader(7, 4, width, height, 2, "Date", false)

	//Action Header
	rendEl.DrawResultHeader(9, 4, width, height, 2, "Actions", true)

	var headerDivider rl.Rectangle = ui.HorizontalDivider(float32(gr.GridPosXLeft(1, width)), float32(gr.GridPosYTop(5, height-1)), width, height, 10)
	rl.DrawRectangleRec(headerDivider, color.MinorAColor())

	// Draw row backdrop
	for i := 0; i < 6; i++ {
		rendEl.DrawResultRowBackdrop(1, 5, i, 3, width, height)
		rendEl.DrawResultRowBackdrop(3, 5, i, 2, width, height)
		rendEl.DrawResultRowBackdrop(5, 5, i, 2, width, height)
		rendEl.DrawResultRowBackdrop(7, 5, i, 2, width, height)
		rendEl.DrawResultRowBackdrop(9, 5, i, 2, width, height)
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
		strBucket2 := ""
		strDate2 := ""
		doubleDraw := false
		if j+1 <= len(records)-1 {
			strAmt2 = strconv.FormatFloat(float64(records[j+1].Amount), 'f', -1, 32)
			strDesc2 = records[j+1].Description
			strDate2 = records[j+1].Date
			strBucket2 = records[j+1].Bucket
			doubleDraw = true
		}
		xButton := icons.XButtonTexture()
		rendEl.DrawResultRowText(records[j].Description, strDesc2, 1, 5, width, height, rowCount, 28, doubleDraw)
		rendEl.DrawResultRowText(records[j].Bucket, strBucket2, 3, 5, width, height, rowCount, 28, doubleDraw)
		rendEl.DrawResultRowText(strAmt1, strAmt2, 5, 5, width, height, rowCount, 28, doubleDraw)
		rendEl.DrawResultRowText(records[j].Date, strDate2, 7, 5, width, height, rowCount, 28, doubleDraw)
		rendEl.DrawResultAction(xButton, 9, 5, width, height, rowCount, doubleDraw)

		//deleteVectorTop := rl.Vector2{X: float32(gr.GridPosXLeft(9, width) + 5), Y: float32(gr.GridPosYTop(5+rowCount, height) + 5)}
		//deleteVectorBot := rl.Vector2{X: float32(gr.GridPosXLeft(9, width) + 5), Y: float32(gr.GridPosYBot(5+rowCount, height) + 5)}
		// listChange = DeleteRecordCheck(width, height, rowCount, deleteVectorTop, records[j])
		// if !listChange && doubleDraw {
		// 	listChange = DeleteRecordCheck(width, height, rowCount, deleteVectorBot, records[j+1])
		// }
		// if listChange {
		// 	return listChange
		// }
		rowCount++
	}

	return listChange
}
