package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/syhbt/fynd/pkg/matcher"
	"github.com/syhbt/fynd/pkg/provider"
)

func main() {
	var queries []string

	// Check if there is piped input
	if stat, _ := os.Stdin.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			queries = append(queries, scanner.Text())
		}
	} else {
		flag.Usage()
		os.Exit(1)
	}

	for _, query := range queries {
		if matchAndPrint(query) {
			continue
		}
		fmt.Printf("%s\x1b[35m not found\x1b[0m\n", query)
	}
}

// matchAndPrint matches the query with data from providers and displays the result
func matchAndPrint(query string) bool {
	hackerOneData, err := provider.GetHackerOneData()
	if err != nil {
		log.Fatal(err)
	}

	bugcrowdData, err := provider.GetBugcrowdData()
	if err != nil {
		log.Fatal(err)
	}

	intigritiData, err := provider.GetIntigritiData()
	if err != nil {
		log.Fatal(err)
	}

	if matcher.MatchURLAssetIdentifier(query, hackerOneData, bugcrowdData, intigritiData) {
		fmt.Printf("%s", query)
		for _, data := range hackerOneData {
			for _, target := range data.Targets.InScope {
				if strings.Contains(target.AssetIdentifier, query) && target.AssetType == "URL" {
					green := color.New(color.FgGreen).SprintFunc()
					cyan := color.New(color.FgCyan).SprintFunc()

					fmt.Printf("%s ", query)
					fmt.Printf("%s %s\n", green(target.AssetIdentifier), cyan(data.URL))
					return true
				}
			}
		}
	}
	return false
}
