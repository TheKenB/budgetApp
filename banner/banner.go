package banner

import (
	gr "main/grid"
	ui "main/inputs"
	pgEntry "main/pages/entryPage"
	pgSearch "main/pages/searchPage"
	color "main/theme"
	util "main/uiUtil"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var entryRec rl.Rectangle
var searchRec rl.Rectangle

type pages struct {
	name string
}

var pageList []pages = []pages{
	{name: "entry"},
	{name: "search"},
}

var curPage pages = pageList[0]

func DrawBanner(dGrid gr.DisplayGrid) {
	var height int = dGrid.Height
	var width int = dGrid.Width
	entryRec = ui.Button(float32(gr.GridPosXLeft(0, width)), float32(gr.GridPosYTop(0, height)), width-5, height, 1)
	searchRec = ui.Button(float32(gr.GridPosXLeft(1, width)), float32(gr.GridPosYTop(0, height)), width-5, height, 1)
	rl.DrawRectangleRec(ui.TextInput(float32(gr.GridPosXLeft(0, width)), float32(gr.GridPosYTop(0, height)), width, height, 13), color.SecondaryColor())
	color.DrawMajorText("Entry", int32(gr.GridPosTextXCent(0, width)), int32(gr.GridPosYTop(0, height)), 32, HandleBannerHover(entryRec))
	color.DrawMajorText("Search", int32(gr.GridPosTextXCent(1, width)), int32(gr.GridPosYTop(0, height)), 32, HandleBannerHover(searchRec))
	HandlePageSelect(entryRec)
	HandlePageSelect(searchRec)
	HandlePageRender(dGrid)

}

func HandleBannerHover(rec rl.Rectangle) rl.Color {
	if util.IsHoverRec(rec) {
		return color.MinorCColor()
	} else {
		return color.MinorBColor()
	}
}

func HandlePageSelect(rec rl.Rectangle) {
	if util.IsHoverRec(rec) && rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		switch rec {
		case entryRec:
			curPage = pageList[0]
		case searchRec:
			curPage = pageList[1]
		}
	}
}

func HandlePageRender(dGrid gr.DisplayGrid) {
	switch curPage {
	case pageList[0]:
		pgEntry.HandlePageEntry(dGrid)
	case pageList[1]:
		pgSearch.HandleSearchPage(dGrid)
	}
}
