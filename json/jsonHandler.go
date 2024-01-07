package jsonHandler

import (
	"encoding/json"
	"fmt"
	"os"
)

type Entries struct {
	Description string
	Amount      float32
	Date        string
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
	records = append(records, newD)

	b, err := json.MarshalIndent(records, "", " ")

	if err != nil {
		fmt.Printf("Failed to write")
	}
	_ = os.WriteFile("./data/entries.json", b, 0644)
}
