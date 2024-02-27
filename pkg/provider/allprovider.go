package provider

import (
	"fmt"
)

type AllDomainData struct {
	DomainData            []DomainData
	FederacyDomainData    []FederacyDomainData
	HackenProofDomainData []HackenProofDomainData
	HackerOneDomainData   []HackerOneDomainData
	IntigritiDomainData   []IntigritiDomainData
	YesWeHackDomainData   []YesWeHackDomainData
}

func ListAllDomain() (AllDomainData, error) {
	var listDomain AllDomainData
	var err error

	listDomain.DomainData, err = ListBugcrowdDomain()
	if err != nil {
		return AllDomainData{}, fmt.Errorf("error getting Bugcrowd domains: %v", err)
	}

	listDomain.FederacyDomainData, err = ListFederacyDomain()
	if err != nil {
		return AllDomainData{}, fmt.Errorf("error getting Federacy domains: %v", err)
	}

	listDomain.HackenProofDomainData, err = ListHackenProofDomain()
	if err != nil {
		return AllDomainData{}, fmt.Errorf("error getting HackenProof domains: %v", err)
	}

	listDomain.HackerOneDomainData, err = ListHackerOneDomain()
	if err != nil {
		return AllDomainData{}, fmt.Errorf("error getting HackerOne domains: %v", err)
	}

	listDomain.IntigritiDomainData, err = ListIntigritiDomain()
	if err != nil {
		return AllDomainData{}, fmt.Errorf("error getting Intigriti domains: %v", err)
	}

	listDomain.YesWeHackDomainData, err = ListYesWeHackDomain()
	if err != nil {
		return AllDomainData{}, fmt.Errorf("error getting YesWeHack domains: %v", err)
	}

	return listDomain, nil
}
