package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/syhbt/fynd/pkg/parser" // replace with your module path
)

func main() {
	onlyFound := flag.Bool("of", false, "Only show found results")
	outputFile := flag.String("o", "", "Output file for matched queries")
	flag.BoolVar(onlyFound, "only-found", false, "Only show found results")
	flag.Parse()

	fyndData, err := parser.GetFyndData()
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	domainToURL := make(map[string]string)
	for _, data := range fyndData {
		for _, domain := range data.Domains {
			domainToURL[domain] = data.URL
		}
	}

	var file *os.File
	if *outputFile != "" {
		file, err = os.Create(*outputFile)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		defer file.Close()
	}

	scanner := bufio.NewScanner(os.Stdin)
	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	for scanner.Scan() {
		query := scanner.Text()
		query = strings.TrimSpace(query)

		found := false
		var url string
		var domain string
		for d, u := range domainToURL {
			if strings.Contains(d, query) && len(query) >= 7 {
				domain = d
				url = u
				found = true
				break
			}
		}

		if found {
			fmt.Printf("%s %s %s\n", query, green(domain), cyan(url))
			if file != nil {
				_, err = file.WriteString(query + "\n")
				if err != nil {
					fmt.Println("Error writing to output file:", err)
					return
				}
			}
		} else if !*onlyFound {
			fmt.Printf("%s %s\n", query, magenta("Not Found"))
		}
	}
}
