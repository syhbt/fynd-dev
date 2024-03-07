package parser

import (
	"encoding/json"
	"os"
)

type FyndData struct {
	Provider string   `json:"provider"`
	Name     string   `json:"name"`
	URL      string   `json:"url"`
	Domains  []string `json:"domains"`
}

func GetFyndData() ([]FyndData, error) {
	jsonFile, err := os.Open("data/fynd_data.json") // Corrected file name
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	var fyndData []FyndData
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&fyndData)
	if err != nil {
		return nil, err
	}

	return fyndData, nil
}
