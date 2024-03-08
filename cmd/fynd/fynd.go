package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/syhbt/fynd/internal"
	"github.com/syhbt/fynd/pkg/matcher"
	"github.com/syhbt/fynd/pkg/parser"

	"github.com/fatih/color"
	"github.com/gosuri/uilive"
)

var onlyFound bool
var outputFile string
var silent bool

func init() {
	flag.BoolVar(&onlyFound, "of", false, "Only display results if found")
	flag.BoolVar(&onlyFound, "only-found", false, "Only display results if found")
	flag.StringVar(&outputFile, "o", "", "Output file to save matched queries")
	flag.BoolVar(&silent, "silent", false, "Don't display the banner")
	flag.Parse()
}

func main() {
	internal.DisplayBanner(silent)
	yellow := color.New(color.FgYellow).SprintFunc()

	start := time.Now()

	writer := uilive.New()
	writer.Start()

	go func() {
		if !silent {
			for {
				elapsed := time.Since(start)
				timer := elapsed.Round(time.Second)
				fmt.Fprintf(writer, "[%s] [%s] Please wait ...\n\n", yellow("Matching Up"), timer)
				time.Sleep(time.Second)
			}
		}
	}()
	data, err := parser.GetFyndData()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var query []string
	for scanner.Scan() {
		query = append(query, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	results := matcher.MatchDomains(query, data)

	writer.Stop()

	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()

	var file *os.File
	if outputFile != "" {
		file, err = os.Create(outputFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}

	for _, result := range results {
		if result.Found {
			fmt.Printf("%s %s %s\n", result.Query, green(result.Domain), cyan(result.URL))
			if file != nil {
				fmt.Fprintf(file, "%s\n", result.Query)
			}
		} else if !onlyFound {
			fmt.Printf("%s %s\n", result.Query, magenta("Not Found"))
		}
	}
}
