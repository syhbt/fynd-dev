package provider

import (
	"encoding/json"
	"os"
)

type BugcrowdTarget struct {
	Type   string `json:"type"`
	Target string `json:"target"`
}

type BugcrowdData struct {
	Name               string         `json:"name"`
	URL                string         `json:"url"`
	AllowsDisclosure   bool           `json:"allows_disclosure"`
	ManagedByBugcrowd bool           `json:"managed_by_bugcrowd"`
	SafeHarbor         string         `json:"safe_harbor"`
	MaxPayout          int            `json:"max_payout"`
	Targets            struct {
		InScope    []BugcrowdTarget `json:"in_scope"`
		OutOfScope []interface{}    `json:"out_of_scope"`
	} `json:"targets"`
}

func GetBugcrowdData() ([]BugcrowdData, error) {
	file, err := os.Open("data/bugcrowd_data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []BugcrowdData
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
