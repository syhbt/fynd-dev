package provider

import (
	"encoding/json"
	"os"
)

type IntigritiTarget struct {
	AssetIdentifier string `json:"endpoint"`
	AssetType       string `json:"type"`
}

type IntigritiTargets struct {
	AssetType string `json:"asset_type"`
}

type IntigritiData struct {
	Name            string `json:"name"`
	CompanyHandle   string `json:"company_handle"`
	URL             string `json:"url"`
	Status          string `json:"status"`
	Confidentiality string `json:"confidentiality"`
	Targets         struct {
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
