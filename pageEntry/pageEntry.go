package pageEntry

import (
	gr "main/grid"
	ui "main/inputs"
	enJson "main/json"
	"strconv"
	"unicode/utf8"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var descText string = ""
var amtText string = ""
var dateText string = ""
var curText ui.ActiveText = ui.ActiveText{Active: false, Pos: [2]int{0, 0}}
var blinking bool = false
var elapsedTime float32 = 0
var addColor = rl.DarkGreen

func HandleEntryPageInput(dGrid gr.DisplayGrid, font rl.Font) bool {
	var height int = dGrid.Height
	var width int = dGrid.Width
	var inputRects []ui.TextCollissionLocation
	var inTextBox bool = false
	elapsedTime += rl.GetFrameTime()
	// Draw Inputs
	//Description Field
	rl.DrawText("Entry Description", int32(gr.GridPosXLeft(1, width)), int32(gr.GridPosYTop(3, height)), 32, rl.DarkGreen)
	var entryDescriptRect ui.TextCollissionLocation = ui.TextCollissionLocation{Location: ui.TextInput(float32(gr.GridPosXLeft(1, width)), float32(gr.GridPosYBot(3, height)), width, height, 3), Text: &descText}
	rl.DrawRectangleRec(entryDescriptRect.Location, rl.LightGray)
	rl.DrawTextEx(font, descText, rl.Vector2{X: entryDescriptRect.Location.X + 15, Y: entryDescriptRect.Location.Y}, 28, 2, rl.Black)

	//Amount Field
	rl.DrawText("Amount", int32(gr.GridPosXLeft(5, width)), int32(gr.GridPosYTop(3, height)), 32, rl.DarkGreen)
	var amountRect ui.TextCollissionLocation = ui.TextCollissionLocation{Location: ui.TextInput(float32(gr.GridPosXLeft(5, width)), float32(gr.GridPosYBot(3, height)), width, height, 2), Text: &amtText}
	rl.DrawRectangleRec(amountRect.Location, rl.LightGray)
	rl.DrawTextEx(font, amtText, rl.Vector2{X: amountRect.Location.X + 15, Y: amountRect.Location.Y}, 28, 2, rl.Black)

	//Date Field
	rl.DrawText("Date", int32(gr.GridPosXLeft(7, width)), int32(gr.GridPosYTop(3, height)), 32, rl.DarkGreen)
	var dateRect ui.TextCollissionLocation = ui.TextCollissionLocation{Location: ui.TextInput(float32(gr.GridPosXLeft(7, width)), float32(gr.GridPosYBot(3, height)), width, height, 2), Text: &dateText}
	rl.DrawRectangleRec(dateRect.Location, rl.LightGray)
	rl.DrawTextEx(font, dateText, rl.Vector2{X: dateRect.Location.X + 15, Y: dateRect.Location.Y}, 28, 2, rl.Black)

	//Add Button
	var addRect = ui.Button(float32(gr.GridPosXLeft(10, width)), float32(gr.GridPosYBot(3, height)), width, height, 1)

	rl.DrawRectangleRec(addRect, addColor)
	rl.DrawText("Add", int32(gr.GridPosTextXCent(10, width)), int32(gr.GridPosYBot(3, height)), 32, rl.White)

	// Check if user is in input boxes
	inputRects = append(inputRects, entryDescriptRect, amountRect, dateRect)

	for _, rect := range inputRects {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rect.Location) {
			inTextBox = true
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				rl.SetMouseCursor((rl.MouseCursorIBeam))
				curText = ui.ActiveText{Active: true, Pos: [2]int{int(rect.Location.X), int(rect.Location.Y)}}
			}
		}
	}
	HandleInputTyping(inputRects)
	var saved bool = HandleAddButton(addRect)

	if !inTextBox && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		curText.Active = false
	}

	if !inTextBox {
		rl.SetMouseCursor(rl.MouseCursorArrow)
	}
	return saved
}

func HandleEntryPageResults(dGrid gr.DisplayGrid, font rl.Font, records []enJson.Entries) {
	var height int = dGrid.Height
	var width int = dGrid.Width
	//	var inputRects []ui.TextCollissionLocation

	elapsedTime += rl.GetFrameTime()
	// Draw Header
	//Description Header
	var entryDescriptRect rl.Rectangle = ui.Button(float32(gr.GridPosXLeft(1, width)), float32(gr.GridPosYBot(5, height-1)), width, height, 3)
	var descColumnDiv rl.Rectangle = ui.HeaderDivider(float32(gr.GridPosXLeft(1, width)), float32(gr.GridPosYBot(5, height-1)), width, height, 3)
	rl.DrawRectangleRec(entryDescriptRect, rl.DarkGray)
	rl.DrawRectangleRec(descColumnDiv, rl.DarkGreen)
	rl.DrawText("Description", int32(entryDescriptRect.X+10), int32(entryDescriptRect.Y), 28, rl.White)

	//Amount Header
	var amountRect rl.Rectangle = ui.Button(float32(gr.GridPosXLeft(4, width)), float32(gr.GridPosYBot(5, height-1)), width, height, 2)
	var amountColumnDiv rl.Rectangle = ui.HeaderDivider(float32(gr.GridPosXLeft(4, width)), float32(gr.GridPosYBot(5, height-1)), width, height, 2)
	rl.DrawRectangleRec(amountRect, rl.DarkGray)
	rl.DrawRectangleRec(amountColumnDiv, rl.DarkGreen)
	rl.DrawText("Amount", int32(amountRect.X+10), int32(amountRect.Y), 28, rl.White)

	//Date Header
	var dateRect rl.Rectangle = ui.Button(float32(gr.GridPosXLeft(6, width)), float32(gr.GridPosYBot(5, height-1)), width, height, 2)
	var datecolumnDiv rl.Rectangle = ui.HeaderDivider(float32(gr.GridPosXLeft(6, width)), float32(gr.GridPosYBot(5, height-1)), width, height-1, 2)
	rl.DrawRectangleRec(dateRect, rl.DarkGray)
	rl.DrawRectangleRec(datecolumnDiv, rl.DarkGreen)
	rl.DrawText("Date", int32(dateRect.X+10), int32(dateRect.Y), 28, rl.White)

	//Action Header
	var actionRect rl.Rectangle = ui.Button(float32(gr.GridPosXLeft(8, width)), float32(gr.GridPosYBot(5, height-1)), width, height, 2)
	rl.DrawRectangleRec(actionRect, rl.DarkGray)
	rl.DrawText("Actions", int32(actionRect.X+10), int32(actionRect.Y), 28, rl.White)

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
		strAmt1 := strconv.FormatFloat(float64(records[j].Amount), 'f', -1, 32)
		rl.DrawTextEx(font, records[j].Description, rl.Vector2{X: float32(gr.GridPosXLeft(1, width) + 5), Y: float32(gr.GridPosYTop(6+rowCount, height))}, 28, 2, rl.Black)
		rl.DrawTextEx(font, strAmt1, rl.Vector2{X: float32(gr.GridPosXLeft(4, width)) + 5, Y: float32(gr.GridPosYTop(6+rowCount, height))}, 28, 2, rl.Black)
		rl.DrawTextEx(font, records[j].Date, rl.Vector2{X: float32(gr.GridPosXLeft(6, width)) + 5, Y: float32(gr.GridPosYTop(6+rowCount, height))}, 28, 2, rl.Black)
		if j+1 <= len(records)-1 {
			strAmt2 := strconv.FormatFloat(float64(records[j+1].Amount), 'f', -1, 32)
			rl.DrawTextEx(font, records[j+1].Description, rl.Vector2{X: float32(gr.GridPosXLeft(1, width) + 5), Y: float32(gr.GridPosYBot(6+rowCount, height))}, 28, 2, rl.Black)
			rl.DrawTextEx(font, strAmt2, rl.Vector2{X: float32(gr.GridPosXLeft(4, width)) + 5, Y: float32(gr.GridPosYBot(6+rowCount, height))}, 28, 2, rl.Black)
			rl.DrawTextEx(font, records[j+1].Date, rl.Vector2{X: float32(gr.GridPosXLeft(6, width)) + 5, Y: float32(gr.GridPosYBot(6+rowCount, height))}, 28, 2, rl.Black)
		}
		rowCount++
	}
}

func ClearInputs() {
	descText = ""
	amtText = ""
	dateText = ""
}

func HandleAddButton(rec rl.Rectangle) bool {
	//Add button hover color indicator
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rec) {
		addColor = rl.Lime
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			amt, err := strconv.ParseFloat(amtText, 32)
			if err != nil {
				panic(err)
			}
			var newEntry enJson.Entries = enJson.Entries{Description: descText, Amount: float32(amt), Date: dateText}
			enJson.SaveEntry(newEntry)
			ClearInputs()
			return true
		}
	} else {
		addColor = rl.DarkGreen
		return false
	}
	return false
}

func HandleInputTyping(recs []ui.TextCollissionLocation) {
	//Check if in typing mode and update text
	var indexMod = 0
	var tabbing = false
	if curText.Active {
		for i, box := range recs {

			// If current collission is a input box
			if int(box.Location.X) == curText.Pos[0] && int(box.Location.Y) == curText.Pos[1] && !tabbing {
				//Skip to next box
				if rl.IsKeyReleased(rl.KeyTab) {
					if i < len(recs)-1 {
						indexMod = i + 1
					}
					curText.Pos = [2]int{int(recs[indexMod].Location.X), int(recs[indexMod].Location.Y)}
					tabbing = true
					continue
				}

				numCharacters := utf8.RuneCountInString(*box.Text)
				key := rl.GetCharPressed()

				if elapsedTime >= 0.5 {
					blinking = !blinking
					elapsedTime = 0
				}
				// Check if more characters have been pressed on the same frame
				for key > 0 {
					// NOTE: Only allow keys in range [32..125]
					if (key >= 32) && (key <= 125) && (numCharacters < 40) {
						*box.Text = *box.Text + string(key)
					}
					key = rl.GetCharPressed() // Check next character in the queue
				}

				if rl.IsKeyPressed(rl.KeyBackspace) {
					if numCharacters != 0 {
						tempString := *box.Text
						*box.Text = tempString[:numCharacters-1]
					}
				}

				var textBuffer int32 = rl.MeasureText(*box.Text, 28) + 15
				if blinking {
					rl.DrawText("_", box.Location.ToInt32().X+textBuffer, box.Location.ToInt32().Y, 28, rl.Black)
				}
			}
		}
	}
}
