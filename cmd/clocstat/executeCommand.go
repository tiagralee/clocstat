package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func execute(params string, verbose bool) ClocResult {
	app := "cloc"
	args := strings.Fields(strings.Replace(params, "'", "", -1))
	currentDir, _ := os.Getwd()
	//prints out the raw cloc output
	if verbose {
		cmd := exec.Command(app, args[0:]...)
		cmd.Dir = currentDir
		out, err := cmd.Output()
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s", string(out))
	}

	args = append([]string{"--json"}, args...)
	cmd := exec.Command(app, args[0:]...)
	cmd.Dir = currentDir
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	var result ClocResult
	json.Unmarshal(out, &result)

	return result
}
