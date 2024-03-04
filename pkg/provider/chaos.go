package provider

import (
	"encoding/json"
	"os"
)

type Program struct {
	Name    string   `json:"name"`
	URL     string   `json:"url"`
	Bounty  bool     `json:"bounty"`
	Swag    bool     `json:"swag"`
	Domains []string `json:"domains"`
}

type ChaosData struct {
	Programs []Program `json:"programs"`
}

func GetChaosData() (ChaosData, error) {
	file, err := os.Open("data/chaos_data.json")
	if err != nil {
		return ChaosData{}, err
	}
	defer file.Close()

	var data ChaosData
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return ChaosData{}, err
	}

	return data, nil
}
