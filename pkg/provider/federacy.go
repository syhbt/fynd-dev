package provider

import (
	"encoding/json"
	"os"
)

type FederacyTarget struct {
	AssetIdentifier string `json:"target"`
	AssetType       string `json:"type"`
}

type FederacyTargets struct {
	AssetType string `json:"type"`
}

type FederacyData struct {
	Name         string `json:"name"`
	OfferRewards bool   `json:"offer_rewards"`
	URL          string `json:"url"`
	Targets      struct {
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
