package main

import (
	"fmt"
	"strconv"
	"strings"
)

const columnLenTitle = 30
const columnLen = 12

func generateReport(compareGroup []string, rawClocMap map[string]ClocResult, columns ...string) {
	writeTitle(compareGroup)
	//build title
	rowBuilder("", columns...)
	var totalFiles int = 0
	for _, item := range compareGroup {
		if clocResult, ok := rawClocMap[item]; ok {
			totalFiles += clocResult.Sum.NFiles
		}
	}
	for _, item := range compareGroup {
		if clocResult, ok := rawClocMap[item]; ok {
			totalLines := clocResult.Sum.Code + clocResult.Sum.Blank + clocResult.Sum.Comment
			title := item
			nFile := strconv.Itoa(clocResult.Sum.NFiles)
			lineOfCode := strconv.Itoa(clocResult.Sum.Code)
			lineOfBlank := strconv.Itoa(clocResult.Sum.Blank)
			lineOfComment := strconv.Itoa(clocResult.Sum.Comment)
			codePercentage := "0"
			blankPercentage := "0"
			commentPercentage := "0"
			filePercentage := "0"
			if totalFiles > 0 {
				filePercentage = strconv.Itoa(clocResult.Sum.NFiles * 100 / totalFiles)
			}
			if totalLines > 0 {
				codePercentage = strconv.Itoa(clocResult.Sum.Code * 100 / totalLines)
				blankPercentage = strconv.Itoa(clocResult.Sum.Blank * 100 / totalLines)
				commentPercentage = strconv.Itoa(clocResult.Sum.Comment * 100 / totalLines)
			}
			var columnValues []string
			for _, c := range columns {
				cValue := strings.TrimSpace(c)
				switch cValue {
				case "code":
					columnValues = append(columnValues, lineOfCode)
				case "code%":
					columnValues = append(columnValues, codePercentage)
				case "blank":
					columnValues = append(columnValues, lineOfBlank)
				case "blank%":
					columnValues = append(columnValues, blankPercentage)
				case "comment":
					columnValues = append(columnValues, lineOfComment)
				case "comment%":
					columnValues = append(columnValues, commentPercentage)
				case "files":
					columnValues = append(columnValues, nFile)
				case "files%":
					columnValues = append(columnValues, filePercentage)
				default:
					// Do nothing
				}

			}

			rowBuilder(title, columnValues...)
		}
	}
}

func writeTitle(compareItems []string) {
	fmt.Printf("\n#Compare %s\n\n", strings.Join(compareItems, ", "))
}

func columnBuilder(columnName string, columnLen int) string {
	return strings.Repeat(" ", columnLen-len(columnName)) + columnName
}

func titleColunmBuilder(columnName string, columnLen int) string {
	return columnName + strings.Repeat(" ", columnLen-len(columnName))
}

func rowBuilder(title string, columns ...string) {

	titleColunm := titleColunmBuilder(title, columnLenTitle)
	seperaterLength := columnLenTitle
	columnsValue := ""
	for _, c := range columns {
		columnsValue += columnBuilder(strings.TrimSpace(c), columnLen)
		seperaterLength += columnLen
	}
	fmt.Printf("%s%s\n", titleColunm, columnsValue)
	fmt.Printf("%s\n", strings.Repeat("-", seperaterLength))
}
