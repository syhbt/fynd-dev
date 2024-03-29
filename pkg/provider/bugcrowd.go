package provider

import (
	"encoding/json"
	"os"
)

type BugcrowdTarget struct {
	AssetIdentifier string `json:"target"`
	AssetType       string `json:"type"`
}

type BugcrowdTargets struct {
	AssetType string `json:"type"`
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
