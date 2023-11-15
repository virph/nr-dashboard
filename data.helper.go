package main

import (
	"fmt"
	"strconv"
	"strings"
)

func buildStrData(raw [][]string) []Data {
	result := make([]Data, 0)

	for idx, row := range raw {
		// skip header row
		if idx == 0 {
			continue
		}

		d := Data{}
		if len(row) < 8 {
			continue
		}

		for idx, col := range row {
			switch idx {
			case 0:
				d.PageNumber, _ = strconv.Atoi(col)
			case 1:
				d.PageTitle = col
			case 2:
				d.MetricTitle = col
			case 3:
				d.Attribute.HitCount = col
			case 4:
				d.Attribute.Latency = col
			case 5:
				d.Filter = buildFilterData(col)
			case 6:
				d.ErrorFilter = buildErrorFilterData(col)
			case 7:
				d.StatusField = col
			}
		}

		result = append(result, d)
	}

	return result
}

func buildFilterData(str string) DataFilter {
	result := DataFilter{
		RawString:       str,
		FilterAttribute: make(map[string]DataFilterAttribute),
	}

	for _, fStr := range strings.Split(str, ",") {
		fils := strings.Split(fStr, ":")
		if len(fils) != 2 {
			continue
		}

		keyName := fils[0]
		isNegate := false
		if keyName[0] == '!' {
			isNegate = true
			keyName = keyName[1:]
		}

		temp := DataFilterAttribute{
			Key:      keyName,
			IsNegate: isNegate,
		}
		if _, ok := result.FilterAttribute[keyName]; ok {
			temp = result.FilterAttribute[keyName]
		}
		temp.Values = append(temp.Values, fils[1])
		result.FilterAttribute[keyName] = temp
	}

	filterStrings := []string{}
	for key, filter := range result.FilterAttribute {
		useIn := len(filter.Values) > 1

		format := ""
		if useIn {
			if filter.IsNegate {
				format = "%s NOT IN (%s)"
			} else {
				format = "%s IN (%s)"
			}
		} else {
			if filter.IsNegate {
				format = "%s != %s"
			} else {
				format = "%s = %s"
			}
		}

		str := fmt.Sprintf(format, key, strings.Join(filter.Values, ","))
		filterStrings = append(filterStrings, str)
	}
	result.FilterString = strings.Join(filterStrings, " AND ")
	if len(result.FilterString) > 0 {
		result.FilterString = result.FilterString + " AND "
	}
	return result
}

func buildErrorFilterData(str string) DataErrorFilter {
	result := DataErrorFilter{
		RawString: str,
	}
	errors := strings.Split(str, ",")

	if len(str) == 0 || len(errors) == 0 {
		return result
	}

	format := "error != [%s] AND "
	if len(errors) > 1 {
		format = "error not in (%s) AND "
	}

	result.FilterString = fmt.Sprintf(format, str)

	return result
}
