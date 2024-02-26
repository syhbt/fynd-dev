package provider

import (
	"encoding/json"
	"os"
)

type HackenProofTarget struct {
	AssetIdentifier string `json:"target"`
	AssetType       string `json:"type"`
}

type HackenProofTargets struct {
	AssetType string `json:"type"`
}

type HackenProofData struct {
	Name                 string `json:"name"`
	URL                  string `json:"url"`
	Archived             bool   `json:"archived"`
	TriagedByHackenproof bool   `json:"triaged_by_hackenproof"`
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
