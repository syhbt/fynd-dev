package provider

import (
	"encoding/json"
	"os"
)

type YesWeHackTarget struct {
	AssetIdentifier string `json:"target"`
	AssetType       string `json:"type"`
}

type YesWeHackData struct {
	Name               string  `json:"name"`
	URL                string  `json:"url"`
	AllowsDisclosure   bool    `json:"allows_disclosure"`
	ManagedByYesWeHack bool    `json:"managed_by_yeswehack"`
	SafeHarbor         string  `json:"safe_harbor"`
	MaxPayout          float64 `json:"max_payout"`
	Targets            struct {
		InScope []YesWeHackTarget `json:"in_scope"`
	} `json:"targets"`
}

func GetYesWeHackData() ([]YesWeHackData, error) {
	file, err := os.Open("data/yeswehack_data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var yeswehack []YesWeHackData
	if err := json.NewDecoder(file).Decode(&yeswehack); err != nil {
		return nil, err
	}

	return yeswehack, nil
}

func ListYesWeHackDomain() ([]string, error) {
	yeswehackData, err := GetYesWeHackData()
	if err != nil {
		return nil, err
	}

	var yeswehackDomain []string
	for _, data := range yeswehackData {
		for _, target := range data.Targets.InScope {
			if target.AssetType == "api" || target.AssetType == "other" || target.AssetType == "web-application" {
				yeswehackDomain = append(yeswehackDomain, target.AssetIdentifier)
			}
		}
	}

	return yeswehackDomain, nil
}
