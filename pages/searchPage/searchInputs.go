package searchPage

import (
	gr "main/grid"
	inputs "main/inputs"
	rendEle "main/renderElements"
)

func HandleSearchPageInpts(dGrid gr.DisplayGrid) {
	var height int = dGrid.Height
	var width int = dGrid.Width
	var tempText string = "Here be a page"
	var entryDescriptRect inputs.TextCollissionLocation = inputs.TextCollissionLocation{Location: inputs.TextInput(float32(gr.GridPosXLeft(1, width)), float32(gr.GridPosYBot(2, height)), width, height, 3), Text: &tempText}
	rendEle.DrawInputs(entryDescriptRect, "Search By Description")

}
