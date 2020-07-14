package main

import (
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "arguments",
		Usage: "arguments example",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: "config",
				Usage: "configuration file for .clocstat",
			},
			&cli.StringFlag{
				Name:  "verbose",
				Value: "verbose",
				Usage: "prints out the raw cloc output for each comparison",
			},
		},
		Action: func(c *cli.Context) error {
			fileName := c.String("config")
			verboseString := c.String("verbose")
			verbose := (verboseString == "yes" || verboseString == "true")
			linesFromConfigFile := fileScanner(fileName)
			commandMap := make(map[string]string)
			for _, item := range linesFromConfigFile {
				if strings.HasPrefix(item, "!compare") {
					compareItems := strings.Split(strings.Split(item, ":")[1], ",")
					m := make(map[string]ClocResult)
					for i := range compareItems {
						compareItems[i] = strings.TrimSpace(compareItems[i])
						m[compareItems[i]] = execute(commandMap[compareItems[i]], verbose)
					}
					if len(strings.Split(item, ":")) > 2 {
						generateReport(compareItems, m, strings.Split(strings.Split(item, ":")[2], ",")...)
					} else {
						generateReport(compareItems, m, []string{"files", "code", "code%"}...)
					}
				}
				s := strings.Split(item, ":")
				commandMap[s[0]] = strings.TrimSpace(s[1])
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
