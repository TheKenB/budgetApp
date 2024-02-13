package searchPage

import (
	gr "main/grid"
	enJson "main/json"
)

var entries []enJson.Entries
var addChange bool = false
var resultChange bool = false
var inital bool = true

func HandleSearchPage(dGrid gr.DisplayGrid) {

	if inital || addChange || resultChange {
		enJson.LoadEntries(&entries)
		ResetLoadStatus()
	}
	HandleSearchPageInpts((dGrid))
	HandleSearchResults(dGrid, entries)
}

func ResetLoadStatus() {
	inital = false
	addChange = false
	resultChange = false
}
