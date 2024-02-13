package searchPage

import (
	gr "main/grid"
	inputs "main/inputs"
	rendEle "main/renderElements"
)

func HandleSearchPageInpts(dGrid gr.DisplayGrid) {
	var height int = dGrid.Height
	var width int = dGrid.Width
	var searchDescription string = "description"
	var searchFromDate string = "from date"
	var searchToDate string = "to date"
	var searchAmount string = "to date"
	var searchBucket string = "to date"

	var searchFromDateRect inputs.TextCollissionLocation = inputs.TextCollissionLocation{Location: inputs.TextInput(float32(gr.GridPosXLeft(3, width)), float32(gr.GridPosYBot(1, height)), width, height, 2), Text: &searchFromDate}
	var searchToDateRect inputs.TextCollissionLocation = inputs.TextCollissionLocation{Location: inputs.TextInput(float32(gr.GridPosXLeft(6, width)), float32(gr.GridPosYBot(1, height)), width, height, 2), Text: &searchToDate}
	var searchDescriptRect inputs.TextCollissionLocation = inputs.TextCollissionLocation{Location: inputs.TextInput(float32(gr.GridPosXLeft(2, width)), float32(gr.GridPosYBot(2, height)), width, height, 3), Text: &searchDescription}
	var searchAmountRect inputs.TextCollissionLocation = inputs.TextCollissionLocation{Location: inputs.TextInput(float32(gr.GridPosXLeft(5, width)), float32(gr.GridPosYBot(2, height)), width, height, 2), Text: &searchAmount}
	var searchBucketRect inputs.TextCollissionLocation = inputs.TextCollissionLocation{Location: inputs.TextInput(float32(gr.GridPosXLeft(7, width)), float32(gr.GridPosYBot(2, height)), width, height, 2), Text: &searchBucket}

	rendEle.DrawInputs(searchFromDateRect, "From Date")
	rendEle.DrawInputs(searchToDateRect, "To Date")
	rendEle.DrawInputs(searchDescriptRect, "Description")
	rendEle.DrawInputs(searchAmountRect, "To Date")
	rendEle.DrawInputs(searchBucketRect, "To Date")
}
