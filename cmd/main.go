package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/syhbt/fynd/pkg/matcher"
)

func main() {
	var queries []string
	var chaos bool

	flag.BoolVar(&chaos, "chaos", false, "Use the Chaos provider")
	flag.Parse()

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
		if chaos {
			if matcher.MatchAndPrintChaos(query) {
				continue
			}
		} else {
			if matcher.MatchAndPrintHackerOne(query) ||
				matcher.MatchAndPrintBugcrowd(query) ||
				matcher.MatchAndPrintFederacy(query) ||
				matcher.MatchAndPrintHackenproof(query) ||
				matcher.MatchAndPrintIntigriti(query) ||
				matcher.MatchAndPrintYesWeHack(query) {
				continue
			}
		}
		magenta := color.New(color.FgMagenta).SprintFunc()
		fmt.Printf("%s %s\n", query, magenta("not found"))
	}
}
