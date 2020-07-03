package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func fileScanner(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var list []string
	for scanner.Scan() {
		lineText := scanner.Text()
		if lineText == "" || strings.HasPrefix(lineText, "#") {
			continue
		}
		list = append(list, lineText)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return list
}
