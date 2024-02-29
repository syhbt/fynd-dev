package matcher

import (
	"strings"

	"github.com/syhbt/fynd/pkg/provider"
	// Ganti "your-package-path/provider" dengan path package yang sesuai
)

// MatchURLAssetIdentifier mencocokan input dengan data asset_identifier yang asset_type nya "URL"
func MatchURLAssetIdentifier(input string, hackerOneData []provider.HackerOneData, bugcrowdData []provider.BugcrowdData, intigritiData []provider.IntigritiData) bool {
	// Cek data dari HackerOne
	for _, d := range hackerOneData {
		for _, target := range d.Targets.InScope {
			if target.AssetType == "URL" && len(input) >= 5 && strings.Contains(target.AssetIdentifier, input) {
				return true
			}
		}
	}

	// Cek data dari Bugcrowd
	for _, d := range bugcrowdData {
		for _, target := range d.Targets.InScope {
			if target.Type == "URL" && len(input) >= 5 && strings.Contains(target.Target, input) {
				return true
			}
		}
	}

	// Cek data dari Intigriti
	for _, d := range intigritiData {
		for _, target := range d.Targets.InScope {
			if target.Type == "URL" && len(input) >= 5 && strings.Contains(target.Endpoint, input) {
				return true
			}
		}
	}

	return false
}
