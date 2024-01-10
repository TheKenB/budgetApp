package pageEntry

import (
	errHand "main/errorHandle"
	gr "main/grid"
	ui "main/inputs"
	enJson "main/json"
	rendEle "main/renderElements"
	color "main/theme"
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
var addColor = color.MinorAColor()
var descErr string = ""
var amtErr string = ""
var dateErr string = ""

func HandleEntryPageInput(dGrid gr.DisplayGrid) bool {
	var height int = dGrid.Height
	var width int = dGrid.Width
	var inputRects []ui.TextCollissionLocation
	var inTextBox bool = false
	elapsedTime += rl.GetFrameTime()
	// Draw Inputs
	//Description Field
	var entryDescriptRect ui.TextCollissionLocation = ui.TextCollissionLocation{Location: ui.TextInput(float32(gr.GridPosXLeft(1, width)), float32(gr.GridPosYBot(2, height)), width, height, 3), Text: &descText}
	rendEle.DrawInputs(entryDescriptRect, "Entry Description")
	rendEle.DrawInputErr(1, 3, width, height, descErr, color.DangerColor(), false)

	//Amount Field
	var amountRect ui.TextCollissionLocation = ui.TextCollissionLocation{Location: ui.TextInput(float32(gr.GridPosXLeft(5, width)), float32(gr.GridPosYBot(2, height)), width, height, 2), Text: &amtText}
	rendEle.DrawInputs(amountRect, "Amount")
	rendEle.DrawInputErr(5, 3, width, height, amtErr, color.DangerColor(), false)

	//Date Field
	var dateRect ui.TextCollissionLocation = ui.TextCollissionLocation{Location: ui.TextInput(float32(gr.GridPosXLeft(7, width)), float32(gr.GridPosYBot(2, height)), width, height, 2), Text: &dateText}
	rendEle.DrawInputs(dateRect, "Date")
	rendEle.DrawInputErr(7, 3, width, height, dateErr, color.DangerColor(), true)

	//Add Button
	var addRect = ui.Button(float32(gr.GridPosXLeft(10, width)), float32(gr.GridPosYBot(2, height)), width, height, 1)

	rl.DrawRectangleRec(addRect, addColor)
	color.DrawMajorText("Add", int32(gr.GridPosTextXCent(10, width)), int32(gr.GridPosYBot(2, height)), 32, color.MinorCColor())

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
	var saved bool = false
	saved = HandleAddButton(addRect)

	if !inTextBox && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		curText.Active = false
	}

	if !inTextBox {
		rl.SetMouseCursor(rl.MouseCursorArrow)
	}
	return saved
}

func ClearInputs() {
	descText = ""
	amtText = ""
	dateText = ""
}

func HandleAddButton(rec rl.Rectangle) bool {
	//Add button hover color indicator
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rec) {
		addColor = color.MinorBColor()
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			descErr = errHand.EntryDescErr(descText)
			amtErr = errHand.EntryAmtErr(amtText)
			dateErr = errHand.EntryDateErr(dateText)
			if len(descErr) == 0 && len(amtErr) == 0 && len(dateErr) == 0 {
				amt, err := strconv.ParseFloat(amtText, 32)
				if err != nil {
					panic(err)
				}
				var newEntry enJson.Entries = enJson.Entries{Description: descText, Amount: float32(amt), Date: dateText, Index: 0}
				enJson.SaveEntry(newEntry)
				ClearInputs()
				return true
			}
		}
	} else {
		addColor = color.MinorAColor()
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

				var textBuffer rl.Vector2 = rl.MeasureTextEx(color.FontMajor, *box.Text, 32, 2)
				if blinking {
					color.DrawMajorText("_", int32(box.Location.X)+int32(textBuffer.X)+15, box.Location.ToInt32().Y, 28, rl.Black)
				}
			}
		}
	}
}
