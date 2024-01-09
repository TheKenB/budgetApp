package main

import (
	banner "main/banner"
	gr "main/grid"
	icons "main/icons"
	enJson "main/json"
	"main/pageEntry"
	entry "main/pageEntry"
	color "main/theme"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var res resolution = res720()
	dGrid := gr.NewGrid([2]int{res.x, res.y})
	//var font rl.Font = rl.LoadFontEx("fonts/Louis George Cafe Bold.ttf", 50, nil)
	var entries []enJson.Entries
	var saved bool = false

	rl.InitWindow(int32(res.x), int32(res.y), "budgeting")
	rl.SetTargetFPS(60)
	pageEntry.LoadTexture()
	enJson.LoadEntries(&entries)
	color.SetMajorFont()
	color.SetMinorFont()
	icons.SetXIcon()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(color.PrimaryColor())
		banner.DrawBanner((dGrid))
		saved = entry.HandleEntryPageInput(dGrid)
		entry.HandleEntryPageResults(dGrid, entries)
		if saved {
			enJson.LoadEntries(&entries)
			saved = false
		}
		//		Draw Grid
		// for i := 0; i < dGrid.Rows; i++ {
		// 	rl.DrawLine(0, int32(i*dGrid.Height), int32(res.x), int32(i*dGrid.Height), rl.Blue)
		// }
		// for j := 0; j < dGrid.Columns; j++ {
		// 	rl.DrawLine(int32(j*dGrid.Width), 0, int32(j*dGrid.Width), int32(res.y), rl.Blue)
		// }

		if rl.IsKeyPressed(rl.KeySemicolon) {
			res = res1080()
			rl.SetWindowSize(res.x, res.y)
			dGrid = gr.NewGrid([2]int{res.x, res.y})
		}
		if rl.IsKeyPressed(rl.KeyComma) {
			res = res720()
			rl.SetWindowSize(res.x, res.y)
			dGrid = gr.NewGrid([2]int{res.x, res.y})
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()

}
