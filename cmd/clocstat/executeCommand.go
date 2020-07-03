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
			fmt.Print(err)
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
	e := json.Unmarshal(out, &result)
	if e != nil {
		log.Print(e)
	}
	return result
}
