package matcher

import (
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/syhbt/fynd/pkg/provider"
)

func MatchAndPrintHackerOne(query string) bool {
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

	federacyData, err := provider.GetFederacyData()
	if err != nil {
		log.Fatal(err)
	}

	hackenproofData, err := provider.GetHackenproofData()
	if err != nil {
		log.Fatal(err)
	}

	yeswehackData, err := provider.GetYesWeHackData()
	if err != nil {
		log.Fatal(err)
	}

	if MatchDomain(query, bugcrowdData, federacyData, hackenproofData, hackerOneData, intigritiData, yeswehackData) {
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

func MatchAndPrintBugcrowd(query string) bool {
	bugcrowdData, err := provider.GetBugcrowdData()
	if err != nil {
		log.Fatal(err)
	}

	hackerOneData, err := provider.GetHackerOneData()
	if err != nil {
		log.Fatal(err)
	}

	intigritiData, err := provider.GetIntigritiData()
	if err != nil {
		log.Fatal(err)
	}

	federacyData, err := provider.GetFederacyData()
	if err != nil {
		log.Fatal(err)
	}

	hackenproofData, err := provider.GetHackenproofData()
	if err != nil {
		log.Fatal(err)
	}

	yeswehackData, err := provider.GetYesWeHackData()
	if err != nil {
		log.Fatal(err)
	}

	if MatchDomain(query, bugcrowdData, federacyData, hackenproofData, hackerOneData, intigritiData, yeswehackData) {
		for _, data := range bugcrowdData {
			for _, target := range data.Targets.InScope {
				if strings.Contains(target.Target, query) && target.Type == "website" {
					green := color.New(color.FgGreen).SprintFunc()
					cyan := color.New(color.FgCyan).SprintFunc()

					fmt.Printf("%s ", query)
					fmt.Printf("%s %s\n", green(target.Target), cyan(data.URL))
					return true
				}
			}
		}
	}
	return false
}

func MatchAndPrintFederacy(query string) bool {
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

	federacyData, err := provider.GetFederacyData()
	if err != nil {
		log.Fatal(err)
	}

	hackenproofData, err := provider.GetHackenproofData()
	if err != nil {
		log.Fatal(err)
	}

	yeswehackData, err := provider.GetYesWeHackData()
	if err != nil {
		log.Fatal(err)
	}

	if MatchDomain(query, bugcrowdData, federacyData, hackenproofData, hackerOneData, intigritiData, yeswehackData) {
		for _, data := range federacyData {
			for _, target := range data.Targets.InScope {
				if strings.Contains(target.Target, query) && target.Type == "website" {
					green := color.New(color.FgGreen).SprintFunc()
					cyan := color.New(color.FgCyan).SprintFunc()

					fmt.Printf("%s ", query)
					fmt.Printf("%s %s\n", green(target.Target), cyan(data.URL))
					return true
				}
			}
		}
	}
	return false
}

func MatchAndPrintHackenproof(query string) bool {
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

	federacyData, err := provider.GetFederacyData()
	if err != nil {
		log.Fatal(err)
	}

	hackenproofData, err := provider.GetHackenproofData()
	if err != nil {
		log.Fatal(err)
	}

	yeswehackData, err := provider.GetYesWeHackData()
	if err != nil {
		log.Fatal(err)
	}

	if MatchDomain(query, bugcrowdData, federacyData, hackenproofData, hackerOneData, intigritiData, yeswehackData) {
		for _, data := range hackenproofData {
			for _, target := range data.Targets.InScope {
				if strings.Contains(target.Target, query) && target.Type == "Web" {
					green := color.New(color.FgGreen).SprintFunc()
					cyan := color.New(color.FgCyan).SprintFunc()

					fmt.Printf("%s ", query)
					fmt.Printf("%s %s\n", green(target.Target), cyan(data.URL))
					return true
				}
			}
		}
	}
	return false
}

func MatchAndPrintIntigriti(query string) bool {
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

	federacyData, err := provider.GetFederacyData()
	if err != nil {
		log.Fatal(err)
	}

	hackenproofData, err := provider.GetHackenproofData()
	if err != nil {
		log.Fatal(err)
	}

	yeswehackData, err := provider.GetYesWeHackData()
	if err != nil {
		log.Fatal(err)
	}

	if MatchDomain(query, bugcrowdData, federacyData, hackenproofData, hackerOneData, intigritiData, yeswehackData) {
		for _, data := range intigritiData {
			for _, target := range data.Targets.InScope {
				if strings.Contains(target.Endpoint, query) && target.Type == "url" {
					green := color.New(color.FgGreen).SprintFunc()
					cyan := color.New(color.FgCyan).SprintFunc()

					fmt.Printf("%s ", query)
					fmt.Printf("%s %s\n", green(target.Endpoint), cyan(data.URL))
					return true
				}
			}
		}
	}
	return false
}

func MatchAndPrintYesWeHack(query string) bool {
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

	federacyData, err := provider.GetFederacyData()
	if err != nil {
		log.Fatal(err)
	}

	hackenproofData, err := provider.GetHackenproofData()
	if err != nil {
		log.Fatal(err)
	}

	yeswehackData, err := provider.GetYesWeHackData()
	if err != nil {
		log.Fatal(err)
	}

	if MatchDomain(query, bugcrowdData, federacyData, hackenproofData, hackerOneData, intigritiData, yeswehackData) {
		for _, data := range yeswehackData {
			for _, target := range data.Targets.InScope {
				if strings.Contains(target.Target, query) && target.Target == "web-application" {
					green := color.New(color.FgGreen).SprintFunc()
					cyan := color.New(color.FgCyan).SprintFunc()

					fmt.Printf("%s ", query)
					fmt.Printf("%s %s\n", green(target.Target), cyan(fmt.Sprintf("https://yeswehack.com/%s", data.ID)))
					return true
				}
			}
		}
	}
	return false
}

func MatchAndPrintChaos(query string) bool {
	chaosDatas, err := provider.GetChaosData()
	if err != nil {
		log.Fatal(err)
	}

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

	federacyData, err := provider.GetFederacyData()
	if err != nil {
		log.Fatal(err)
	}

	hackenproofData, err := provider.GetHackenproofData()
	if err != nil {
		log.Fatal(err)
	}

	yeswehackData, err := provider.GetYesWeHackData()
	if err != nil {
		log.Fatal(err)
	}

	if MatchDomain(query, bugcrowdData, federacyData, hackenproofData, hackerOneData, intigritiData, yeswehackData) {
		for _, data := range chaosDatas.Programs {
			for _, domain := range data.Domains {
				if strings.Contains(domain, query) {
					green := color.New(color.FgGreen).SprintFunc()
					cyan := color.New(color.FgCyan).SprintFunc()

					fmt.Printf("%s ", query)
					fmt.Printf("%s %s\n", green(domain), cyan(data.URL))
					return true
				}
			}
		}
	}

	return false
}
