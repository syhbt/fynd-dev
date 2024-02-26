package provider

import (
	"encoding/json"
	"os"
)

type YesWeHackTarget struct {
	AssetIdentifier string `json:"target"`
	AssetType       string `json:"type"`
}

type YesWeHackTargets struct {
	AssetType string `json:"asset_type"`
}

type YesWeHackData struct {
	IntigritiData string  `json:"id"`
	Name          string  `json:"name"`
	Public        bool    `json:"public"`
	Disabled      bool    `json:"disabled"`
	Managed       bool    `json:"managed"`
	MinBounty     float64 `json:"min_bounty"`
	MaxBounty     float64 `json:"max_bounty"`
	Targets       struct {
		InScope []YesWeHackTarget `json:"in_scope"`
	} `json:"targets"`
	MaxSeverity string `json:"max_severity"`
	URL         string `json:"url"`
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
