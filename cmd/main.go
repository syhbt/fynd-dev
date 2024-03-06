package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/syhbt/fynd/pkg/parser" // replace with your module path
)

func main() {
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

	scanner := bufio.NewScanner(os.Stdin)
	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	for scanner.Scan() {
		query := scanner.Text()
		query = strings.TrimSpace(query)

		found := false
		var url string
		for domain, u := range domainToURL {
			if strings.Contains(domain, query) {
				url = u
				found = true
				break
			}
		}

		if found {
			fmt.Printf("%s %s %s\n", query, green(query), cyan(url))
		} else {
			fmt.Printf("%s %s\n", query, magenta("Not Found"))
		}
	}
}
