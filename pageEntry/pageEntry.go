package pageEntry

import (
	gr "main/grid"
	ui "main/inputs"
	"unicode/utf8"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var font rl.Font = rl.LoadFontEx("budgetApp/fonts/CreatoDisplay-Black.otf", 28, nil)
var descText string = ""
var amtText string = ""
var dateText string = ""
var curText ui.ActiveText = ui.ActiveText{Active: false, Pos: [2]int{0, 0}}
var blinking bool = false
var elapsedTime float32 = 0

func HandleEntryPage(dGrid gr.DisplayGrid) {

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
	rl.DrawTextEx(font, descText, rl.Vector2{X: entryDescriptRect.Location.X + 10, Y: entryDescriptRect.Location.Y}, 28, 2, rl.Black)

	//Amount Field
	rl.DrawText("Amount", int32(gr.GridPosXLeft(5, width)), int32(gr.GridPosYTop(3, height)), 32, rl.DarkGreen)
	var amountRect ui.TextCollissionLocation = ui.TextCollissionLocation{Location: ui.TextInput(float32(gr.GridPosXLeft(5, width)), float32(gr.GridPosYBot(3, height)), width, height, 2), Text: &amtText}
	rl.DrawRectangleRec(amountRect.Location, rl.LightGray)
	rl.DrawText(amtText, amountRect.Location.ToInt32().X+10, amountRect.Location.ToInt32().Y, 28, rl.Black)

	//Date Field
	rl.DrawText("Date", int32(gr.GridPosXLeft(7, width)), int32(gr.GridPosYTop(3, height)), 32, rl.DarkGreen)
	var dateRect ui.TextCollissionLocation = ui.TextCollissionLocation{Location: ui.TextInput(float32(gr.GridPosXLeft(7, width)), float32(gr.GridPosYBot(3, height)), width, height, 2), Text: &dateText}
	rl.DrawRectangleRec(dateRect.Location, rl.LightGray)
	rl.DrawText(dateText, dateRect.Location.ToInt32().X+10, dateRect.Location.ToInt32().Y, 28, rl.Black)

	//Add Button
	var addRect = ui.Button(float32(gr.GridPosXLeft(10, width)), float32(gr.GridPosYBot(3, height)), width, height, 1)
	rl.DrawRectangleRec(addRect, rl.DarkGreen)
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

	if !inTextBox && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		curText.Active = false
	}

	if !inTextBox {
		rl.SetMouseCursor(rl.MouseCursorArrow)
	}

	//Check if in typing mode and update text
	if curText.Active {
		for _, box := range inputRects {
			// If current collission is a input box
			if int(box.Location.X) == curText.Pos[0] && int(box.Location.Y) == curText.Pos[1] {
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
