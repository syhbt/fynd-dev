package provider

import (
	"encoding/json"
	"os"
)

type IntigritiBounty struct {
	Value    int    `json:"value"`
	Currency string `json:"currency"`
}

type IntigritiTarget struct {
	Type        string `json:"type"`
	Endpoint    string `json:"endpoint"`
	Description string `json:"description"`
	Impact      string `json:"impact"`
}

type IntigritiData struct {
	ID                   string          `json:"id"`
	Name                 string          `json:"name"`
	CompanyHandle        string          `json:"company_handle"`
	Handle               string          `json:"handle"`
	URL                  string          `json:"url"`
	Status               string          `json:"status"`
	ConfidentialityLevel string          `json:"confidentiality_level"`
	MinBounty            IntigritiBounty `json:"min_bounty"`
	MaxBounty            IntigritiBounty `json:"max_bounty"`
	Targets              struct {
		InScope []IntigritiTarget `json:"in_scope"`
	} `json:"targets"`
}

func GetIntigritiData() ([]IntigritiData, error) {
	file, err := os.Open("data/intigriti_data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []IntigritiData
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
