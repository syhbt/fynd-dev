package provider

func ListAllDomains() ([]string, error) {

	bugcrowdDomains, err := ListBugcrowdDomain()
	if err != nil {
		return nil, err
	}

	federacyDomains, err := ListFederacyDomain()
	if err != nil {
		return nil, err
	}

	hackenproofDomains, err := ListHackenProofDomain()
	if err != nil {
		return nil, err
	}

	hackeroneDomains, err := ListHackerOneDomain()
	if err != nil {
		return nil, err
	}
	intigritiDomains, err := ListIntigritiDomain()
	if err != nil {
		return nil, err
	}
	yeswehackDomains, err := ListYesWeHackDomain()
	if err != nil {
		return nil, err
	}

	// Menggabungkan domain dari HackerOne dan Bugcrowd
	listDomain := append(bugcrowdDomains, federacyDomains...)
	listDomain = append(listDomain, hackenproofDomains...)
	listDomain = append(listDomain, hackeroneDomains...)
	listDomain = append(listDomain, intigritiDomains...)
	listDomain = append(listDomain, yeswehackDomains...)

	return listDomain, nil
}
