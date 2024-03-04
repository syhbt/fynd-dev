package matcher

import (
	"strings"

	"github.com/syhbt/fynd/pkg/provider"
)

func MatchDomain(input string, bugcrowdData []provider.BugcrowdData, federacyData []provider.FederacyData, hackenproofData []provider.HackenproofData, hackerOneData []provider.HackerOneData, intigritiData []provider.IntigritiData, yeswehackData []provider.YesWeHackData) bool {
	// Check data from Bugcrowd
	for _, d := range bugcrowdData {
		for _, target := range d.Targets.InScope {
			if (target.Type == "api" || target.Type == "website") && len(input) >= 5 && strings.Contains(target.Target, input) {
				return true
			}
		}
	}

	// Check data from Federacy
	for _, d := range federacyData {
		for _, target := range d.Targets.InScope {
			if (target.Type == "api" || target.Type == "webiste") && len(input) >= 5 && strings.Contains(target.Target, input) {
				return true
			}
		}
	}

	// Check data from Hackenproof
	for _, d := range hackenproofData {
		for _, target := range d.Targets.InScope {
			if (target.Type == "API" || target.Type == "Web" || target.Type == "Web3") && len(input) >= 5 && strings.Contains(target.Target, input) {
				return true
			}
		}
	}

	// Check data from HackerOne
	for _, d := range hackerOneData {
		for _, target := range d.Targets.InScope {
			if (target.AssetType == "URL" || target.AssetType == "WILDCARD") && len(input) >= 5 && strings.Contains(target.AssetIdentifier, input) {
				return true
			}
		}
	}

	// Check data from Intigriti
	for _, d := range intigritiData {
		for _, target := range d.Targets.InScope {
			if (target.Type == "url" || target.Type == "other") && len(input) >= 5 && strings.Contains(target.Endpoint, input) {
				return true
			}
		}
	}

	// Check data from Yeswehack
	for _, d := range yeswehackData {
		for _, target := range d.Targets.InScope {
			if (target.Type == "api" || target.Type == "web-application") && len(input) >= 5 && strings.Contains(target.Target, input) {
				return true
			}
		}
	}

	return false
}
