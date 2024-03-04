package provider

import (
	"encoding/json"
	"os"
)

type YesWeHackTarget struct {
	Target string `json:"target"`
	Type   string `json:"type"`
}

type YesWeHackTargets struct {
	InScope    []YesWeHackTarget `json:"in_scope"`
	OutOfScope []YesWeHackTarget `json:"out_of_scope"`
}

type YesWeHackData struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Public    bool             `json:"public"`
	Disabled  bool             `json:"disabled"`
	Managed   *bool            `json:"managed"`
	MinBounty int              `json:"min_bounty"`
	MaxBounty int              `json:"max_bounty"`
	Targets   YesWeHackTargets `json:"targets"`
}

func GetYesWeHackData() ([]YesWeHackData, error) {
	file, err := os.Open("yeswehack_data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []YesWeHackData
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
