package provider

import (
	"encoding/json"
	"os"
)

type HackerOneTarget struct {
	AssetIdentifier string `json:"asset_identifier"`
	AssetType       string `json:"asset_type"`
}

type HackerOneTargets struct {
	AssetType string `json:"asset_type"`
}

type HackerOneData struct {
	Name                              string  `json:"name"`
	Website                           string  `json:"website"`
	SubmissionState                   string  `json:"submission_state"`
	OffersBounties                    bool    `json:"offers_bounties"`
	OffersSwag                        bool    `json:"offers_swag"`
	AverageTimeToBountyAwarded        float64 `json:"average_time_to_bounty_awarded"`
	AverageTimeToFirstProgramResponse float64 `json:"average_time_to_first_program_response"`
	Targets                           struct {
		InScope []HackerOneTarget `json:"in_scope"`
	} `json:"targets"`
	MaxSeverity string `json:"max_severity"`
	URL         string `json:"url"`
}

func GetHackerOneData() ([]HackerOneData, error) {
	file, err := os.Open("data/hackerone_data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var hackerone []HackerOneData
	if err := json.NewDecoder(file).Decode(&hackerone); err != nil {
		return nil, err
	}

	return hackerone, nil
}
