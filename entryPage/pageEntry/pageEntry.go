package pageEntry

import (
	dg "main/grid"
	enJson "main/json"
)

var entries []enJson.Entries
var addChange bool = false
var resultChange bool = false
var inital bool = true

func HandlePageEntry(grid dg.DisplayGrid) bool {
	if inital || addChange || resultChange {
		enJson.LoadEntries(&entries)
		ResetLoadStatus()
	}
	addChange = HandleEntryPageInput(grid)
	resultChange = HandleEntryPageResults(grid, entries)

	if addChange || resultChange {
		return true
	} else {
		return false
	}
}

func ResetLoadStatus() {
	inital = false
	addChange = false
	resultChange = false
}
