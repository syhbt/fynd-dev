package provider

func ListAllDomain() ([]DomainData, error) {
	var listDomain []DomainData

	domainBugcrowd, err := ListBugcrowdDomain()
	if err != nil {
		return nil, err
	}
	listDomain = append(listDomain, domainBugcrowd...)

	domainFederacy, err := ListFederacyDomain()
	if err != nil {
		return nil, err
	}
	listDomain = append(listDomain, domainFederacy...)

	domainHackenProof, err := ListHackenProofDomain()
	if err != nil {
		return nil, err
	}
	listDomain = append(listDomain, domainHackenProof...)

	domainHackerOne, err := ListHackerOneDomain()
	if err != nil {
		return nil, err
	}
	listDomain = append(listDomain, domainHackerOne...)

	domainIntigriti, err := ListIntigritiDomain()
	if err != nil {
		return nil, err
	}
	listDomain = append(listDomain, domainIntigriti...)

	domainYesWeHack, err := ListYesWeHackDomain()
	if err != nil {
		return nil, err
	}
	listDomain = append(listDomain, domainYesWeHack...)

	return listDomain, nil
}
