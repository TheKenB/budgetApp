package entryError

import (
	"regexp"
	"strconv"
	"strings"
)

func EntryDescErr(desc string) string {
	if len(desc) > 18 || len(desc) == 0 {
		return "Please enter a valid description"
	} else {
		return ""
	}

}

func EntryAmtErr(num string) string {
	if len(num) > 11 || len(num) == 0 {
		return "Please enter a valid amount"
	}
	if num == "" {
		return "Please enter an amount"
	}

	if strings.Count(num, ".") > 1 {
		return "Please use only one decimal"
	}

	tempString := strings.Replace(num, ".", "", -1)
	_, err := strconv.Atoi(tempString)
	if err != nil {
		return "Enter a amount in style of 0.00 or 000"
	}

	return ""
}

func EntryDateErr(date string) string {
	if len(date) == 0 || len(date) < 8 {
		return "Enter date in format MM/dd/yyyy"
	}

	if !regexp.MustCompile(`\d{1,2}\/\d{1,2}\/\d{2,4}`).MatchString(date) {
		return "Enter date in format MM/dd/yyyy"
	}

	return ""
}
