package provider

import (
	"encoding/json"
	"os"
)

type FederacyTarget struct {
	AssetIdentifier string `json:"target"`
	AssetType       string `json:"type"`
}

type FederacyData struct {
	Name              string  `json:"name"`
	URL               string  `json:"url"`
	AllowsDisclosure  bool    `json:"allows_disclosure"`
	ManagedByFederacy bool    `json:"managed_by_federacy"`
	SafeHarbor        string  `json:"safe_harbor"`
	MaxPayout         float64 `json:"max_payout"`
	Targets           struct {
		InScope []FederacyTarget `json:"in_scope"`
	} `json:"targets"`
}

func GetFederacyData() ([]FederacyData, error) {
	file, err := os.Open("data/federacy_data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var federacy []FederacyData
	if err := json.NewDecoder(file).Decode(&federacy); err != nil {
		return nil, err
	}

	return federacy, nil
}

func ListFederacyDomain() ([]DomainData, error) {
	federacyData, err := GetFederacyData()
	if err != nil {
		return nil, err
	}

	var domainFederacy []DomainData
	for _, data := range federacyData {
		for _, target := range data.Targets.InScope {
			if target.AssetType == "website" || target.AssetType == "api" {
				domainFederacy = append(domainFederacy, DomainData{URL: data.URL, AssetIdentifier: target.AssetIdentifier})
			}
		}
	}

	return domainFederacy, nil
}
