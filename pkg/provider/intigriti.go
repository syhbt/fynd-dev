package provider

import (
	"encoding/json"
	"os"
)

type IntigritiTarget struct {
	AssetIdentifier string `json:"endpoint"`
	AssetType       string `json:"type"`
}

type IntigritiData struct {
	Name               string  `json:"name"`
	URL                string  `json:"url"`
	AllowsDisclosure   bool    `json:"allows_disclosure"`
	ManagedByIntigriti bool    `json:"managed_by_intigriti"`
	SafeHarbor         string  `json:"safe_harbor"`
	MaxPayout          float64 `json:"max_payout"`
	Targets            struct {
		InScope []IntigritiTarget `json:"in_scope"`
	} `json:"targets"`
}

func GetIntigritiData() ([]IntigritiData, error) {
	file, err := os.Open("data/intigriti_data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var intigriti []IntigritiData
	if err := json.NewDecoder(file).Decode(&intigriti); err != nil {
		return nil, err
	}

	return intigriti, nil
}

func ListIntigritiDomain() ([]string, error) {
	intigritiData, err := GetIntigritiData()
	if err != nil {
		return nil, err
	}

	var intigritiDomain []string
	for _, data := range intigritiData {
		for _, target := range data.Targets.InScope {
			if target.AssetType == "other" || target.AssetType == "url" {
				intigritiDomain = append(intigritiDomain, target.AssetIdentifier)
			}
		}
	}

	return intigritiDomain, nil
}
