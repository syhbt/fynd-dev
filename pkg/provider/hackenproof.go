package provider

import (
	"encoding/json"
	"os"
)

type HackenProofTarget struct {
	AssetIdentifier string `json:"target"`
	AssetType       string `json:"type"`
}

type HackenProofData struct {
	Name                 string  `json:"name"`
	URL                  string  `json:"url"`
	AllowsDisclosure     bool    `json:"allows_disclosure"`
	ManagedByHackenProof bool    `json:"managed_by_hackenproof"`
	SafeHarbor           string  `json:"safe_harbor"`
	MaxPayout            float64 `json:"max_payout"`
	Targets              struct {
		InScope []HackenProofTarget `json:"in_scope"`
	} `json:"targets"`
}

func GetHackenProofData() ([]HackenProofData, error) {
	file, err := os.Open("data/hackenproof_data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var hackenproof []HackenProofData
	if err := json.NewDecoder(file).Decode(&hackenproof); err != nil {
		return nil, err
	}

	return hackenproof, nil
}

func ListHackenProofDomain() ([]string, error) {
	hackenproofData, err := GetHackenProofData()
	if err != nil {
		return nil, err
	}

	var hackenproofDomain []string
	for _, data := range hackenproofData {
		for _, target := range data.Targets.InScope {
			if target.AssetType == "API" || target.AssetType == "web" || target.AssetType == "web3" {
				hackenproofDomain = append(hackenproofDomain, target.AssetIdentifier)
			}
		}
	}

	return hackenproofDomain, nil
}
