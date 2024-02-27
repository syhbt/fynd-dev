package provider

import (
	"encoding/json"
	"os"
)

type BugcrowdTarget struct {
	AssetIdentifier string `json:"target"`
	AssetType       string `json:"type"`
}

type BugcrowdData struct {
	Name              string  `json:"name"`
	URL               string  `json:"url"`
	AllowsDisclosure  bool    `json:"allows_disclosure"`
	ManagedByBugcrowd bool    `json:"managed_by_bugcrowd"`
	SafeHarbor        string  `json:"safe_harbor"`
	MaxPayout         float64 `json:"max_payout"`
	Targets           struct {
		InScope []BugcrowdTarget `json:"in_scope"`
	} `json:"targets"`
}

type DomainData struct {
	URL             string
	AssetIdentifier string
}

func GetBugcrowdData() ([]BugcrowdData, error) {
	file, err := os.Open("data/bugcrowd_data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var bugcrowd []BugcrowdData
	if err := json.NewDecoder(file).Decode(&bugcrowd); err != nil {
		return nil, err
	}

	return bugcrowd, nil
}

func ListBugcrowdDomain() ([]DomainData, error) {
	bugcrowdData, err := GetBugcrowdData()
	if err != nil {
		return nil, err
	}

	var domainBugcrowd []DomainData
	for _, data := range bugcrowdData {
		for _, target := range data.Targets.InScope {
			if target.AssetType == "website" || target.AssetType == "api" {
				domainBugcrowd = append(domainBugcrowd, DomainData{URL: data.URL, AssetIdentifier: target.AssetIdentifier})
			}
		}
	}

	return domainBugcrowd, nil
}
