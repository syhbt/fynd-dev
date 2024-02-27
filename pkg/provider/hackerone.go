package provider

import (
	"encoding/json"
	"os"
)

type HackerOneTarget struct {
	AssetIdentifier string `json:"target"`
	AssetType       string `json:"type"`
}

type HackerOneData struct {
	Name               string  `json:"name"`
	URL                string  `json:"url"`
	AllowsDisclosure   bool    `json:"allows_disclosure"`
	ManagedByHackerOne bool    `json:"managed_by_hackerone"`
	SafeHarbor         string  `json:"safe_harbor"`
	MaxPayout          float64 `json:"max_payout"`
	Targets            struct {
		InScope []HackerOneTarget `json:"in_scope"`
	} `json:"targets"`
}

func GetHackerOneData() ([]HackerOneData, error) {
	file, err := os.Open("data/hackerone_data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var hackerone []HackerOneData
	if err := json.NewDecoder(file).Decode(&hackerone); err != nil {
		return nil, err
	}

	return hackerone, nil
}

func ListHackerOneDomain() ([]DomainData, error) {
	hackeroneData, err := GetHackerOneData()
	if err != nil {
		return nil, err
	}

	var domainHackerOne []DomainData
	for _, data := range hackeroneData {
		for _, target := range data.Targets.InScope {
			if target.AssetType == "website" || target.AssetType == "api" {
				domainHackerOne = append(domainHackerOne, DomainData{URL: data.URL, AssetIdentifier: target.AssetIdentifier})
			}
		}
	}

	return domainHackerOne, nil
}
