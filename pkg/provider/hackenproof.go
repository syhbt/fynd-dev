package provider

import (
	"encoding/json"
	"os"
)

type HackenproofTarget struct {
	Target      string `json:"target"`
	Type        string `json:"type"`
	Instruction string `json:"instruction"`
	Severity    string `json:"severity"`
	Reward      string `json:"reward"`
}

type HackenproofTargets struct {
	InScope    []HackenproofTarget `json:"in_scope"`
	OutOfScope []HackenproofTarget `json:"out_of_scope"`
}

type HackenproofData struct {
	ID                   string             `json:"id"`
	Name                 string             `json:"name"`
	URL                  string             `json:"url"`
	Archived             bool               `json:"archived"`
	TriagedByHackenproof bool               `json:"triaged_by_hackenproof"`
	Targets              HackenproofTargets `json:"targets"`
}

func GetHackenproofData() ([]HackenproofData, error) {
	file, err := os.Open("data/hackenproof_data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []HackenproofData
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
