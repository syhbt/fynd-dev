package matcher

import (
	"errors"

	"github.com/syhbt/fynd/pkg/provider"
)

func GetURLByDomain(matchedDomain string) (string, error) {
	var url string

	bugcrowdData, err := provider.GetBugcrowdData()
	if err != nil {
		return "", err
	}

	federacyData, err := provider.GetFederacyData()
	if err != nil {
		return "", err
	}

	hackenproofData, err := provider.GetHackenProofData()
	if err != nil {
		return "", err
	}

	hackerOneData, err := provider.GetHackerOneData()
	if err != nil {
		return "", err
	}

	intigritiData, err := provider.GetIntigritiData()
	if err != nil {
		return "", err
	}
	yeswehackData, err := provider.GetYesWeHackData()
	if err != nil {
		return "", err
	}

	// Cari di Bugcrowd Data
	for _, data := range bugcrowdData {
		for _, target := range data.Targets.InScope {
			if target.AssetIdentifier == matchedDomain {
				url = data.URL
				return url, nil
			}
		}
	}

	// Cari di Bugcrowd Data
	for _, data := range federacyData {
		for _, target := range data.Targets.InScope {
			if target.AssetIdentifier == matchedDomain {
				url = data.URL
				return url, nil
			}
		}
	}

	// Cari di Bugcrowd Data
	for _, data := range hackenproofData {
		for _, target := range data.Targets.InScope {
			if target.AssetIdentifier == matchedDomain {
				url = data.URL
				return url, nil
			}
		}
	}

	// Cari di HackerOne Data
	for _, data := range hackerOneData {
		for _, target := range data.Targets.InScope {
			if target.AssetIdentifier == matchedDomain {
				url = data.URL
				return url, nil
			}
		}
	}

	// Cari di Bugcrowd Data
	for _, data := range intigritiData {
		for _, target := range data.Targets.InScope {
			if target.AssetIdentifier == matchedDomain {
				url = data.URL
				return url, nil
			}
		}
	}

	// Cari di Bugcrowd Data
	for _, data := range yeswehackData {
		for _, target := range data.Targets.InScope {
			if target.AssetIdentifier == matchedDomain {
				url = data.URL
				return url, nil
			}
		}
	}
	return "", errors.New("URL tidak ditemukan untuk domain yang cocok")
}
