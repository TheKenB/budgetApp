package main

import (
	banner "main/banner"
	gr "main/grid"
	enJson "main/json"
	entry "main/pageEntry"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var res resolution = res720()
	dGrid := gr.NewGrid([2]int{res.x, res.y})
	rl.InitWindow(int32(res.x), int32(res.y), "budgeting")
	rl.SetTargetFPS(60)
	var font rl.Font = rl.LoadFontEx("fonts/Louis George Cafe Bold.ttf", 50, nil)
	var entries []enJson.Entries
	var saved bool = false
	enJson.LoadEntries(&entries)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		banner.DrawBanner((dGrid))
		saved = entry.HandleEntryPageInput(dGrid, font)
		entry.HandleEntryPageResults(dGrid, font, entries)
		if saved {
			enJson.LoadEntries(&entries)
			saved = false
		}
		// //		Draw Grid
		// for i := 0; i < dGrid.Rows; i++ {
		// 	rl.DrawLine(0, int32(i*dGrid.Height), int32(res.x), int32(i*dGrid.Height), rl.Blue)
		// }
		// for j := 0; j < dGrid.Columns; j++ {
		// 	rl.DrawLine(int32(j*dGrid.Width), 0, int32(j*dGrid.Width), int32(res.y), rl.Blue)
		// }

		if rl.IsKeyPressed(rl.KeyZ) {
			res = res1080()
			rl.SetWindowSize(res.x, res.y)
			dGrid = gr.NewGrid([2]int{res.x, res.y})
		}
		if rl.IsKeyPressed(rl.KeyX) {
			res = res720()
			rl.SetWindowSize(res.x, res.y)
			dGrid = gr.NewGrid([2]int{res.x, res.y})
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()

}
