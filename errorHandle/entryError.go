package entryError

import (
	"regexp"
	"strconv"
	"strings"
)

func EntryDescErr(desc string) string {
	if len(desc) > 18 {
		return "Max length 18"
	} else if len(desc) == 0 {
		return "Desc required"
	} else {
		return ""
	}
}

func EntryBucketErr(bucket string) string {
	if len(bucket) > 11 {
		return "Max length 12"
	} else if len(bucket) == 0 {
		return "Bucket required"
	} else {
		return ""
	}
}

func EntryAmtErr(num string) string {
	if len(num) > 11 {
		return "Max length 12"
	} else if len(num) == 0 {
		return "Amount required"
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
		return "Format of 0.00 or 000"
	}
	return ""
}

func EntryDateErr(date string) string {
	if len(date) == 0 || len(date) < 8 {
		return "Format of MM/dd/yyyy"
	}

	if !regexp.MustCompile(`\d{1,2}\/\d{1,2}\/\d{2,4}`).MatchString(date) {
		return "Format of MM/dd/yyyy"
	}

	return ""
}
