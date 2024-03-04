package provider

import (
	"encoding/json"
	"os"
)

type FederacyTarget struct {
	Target string `json:"target"`
	Type   string `json:"type"`
}

type FederacyTargets struct {
	InScope    []FederacyTarget `json:"in_scope"`
	OutOfScope []FederacyTarget `json:"out_of_scope"`
}

type FederacyData struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	OffersAwards bool            `json:"offers_awards"`
	URL          string          `json:"url"`
	Targets      FederacyTargets `json:"targets"`
}

func GetFederacyData() ([]FederacyData, error) {
	file, err := os.Open("data/federacy_data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []FederacyData
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
