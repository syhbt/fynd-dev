package main

import (
	"bufio"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/syhbt/fynd/pkg/matcher"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		query := scanner.Text()
		matchedDomains, err := matcher.MatchDomains([]string{query})
		if err != nil {
			log.Fatal(err)
		}

		for matchedDomain, _ := range matchedDomains {
			color.New(color.FgWhite).Printf("%s ", query)
			color.New(color.FgYellow).Printf("%s ", matchedDomain)
			color.New(color.FgGreen).Printf("%s\n", "Found!")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
