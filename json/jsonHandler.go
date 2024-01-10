package jsonHandler

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Entries struct {
	Index       int32
	Description string
	Amount      float32
	Date        string
}

type Indexes struct {
	IndexCount int32
	Data       string
}

func LoadEntries(d *[]Entries) *[]Entries {
	data, err := os.ReadFile("./data/entries.json")

	if err != nil {
		fmt.Println("--No Entry data found")
		return &[]Entries{}
	}

	err = json.Unmarshal(data, &d)
	if err != nil {
		panic(err)
	}
	return d
}

func SaveEntry(newD Entries) {
	data, err := os.ReadFile("./data/entries.json")

	if err != nil {
		fmt.Println("--No Entry data found")
		return
	}

	var records []Entries
	err = json.Unmarshal(data, &records)

	if err != nil {
		panic(err)
	}
	newIndex := GetNextIndex("entries")

	newD.Index = newIndex
	records = append(records, newD)

	b, err := json.MarshalIndent(records, "", " ")

	if err != nil {
		fmt.Printf("Failed to write")
	}
	_ = os.WriteFile("./data/entries.json", b, 0644)
}

func DeleteEntry(newD Entries) {
	data, err := os.ReadFile("./data/entries.json")

	if err != nil {
		fmt.Println("--No Entry data found")
		return
	}
	var records []Entries
	err = json.Unmarshal(data, &records)

	if err != nil {
		panic(err)
	}

	for _, val := range records {
		if newD.Index == val.Index {
			records = append(records[:val.Index], records[val.Index+1:]...)
		}
	}
	newRec, err := json.MarshalIndent(records, "", " ")

	if err != nil {
		panic(err)
	}
	_ = os.WriteFile("./data/entries.json", newRec, 0644)
}

func GetNextIndex(fileName string) int32 {
	var nextIndex int32 = 0
	indCount, err := os.ReadFile("./data/indexes.json")
	if err != nil {
		fmt.Println("--No Index data found")
	}
	var indexes []Indexes
	err = json.Unmarshal(indCount, &indexes)

	if err != nil {
		panic(err)
	}

	for i, val := range indexes {
		if strings.Compare(val.Data, fileName) == 0 {
			indexes[i].IndexCount += 1
			nextIndex = indexes[i].IndexCount
		}
	}

	newRec, err := json.MarshalIndent(indexes, "", " ")

	if err != nil {
		panic(err)
	}
	_ = os.WriteFile("./data/indexes.json", newRec, 0644)

	return nextIndex
}
