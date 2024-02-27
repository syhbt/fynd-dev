package matcher

import (
	"strings"

	"github.com/syhbt/fynd/pkg/provider" // replace with your actual project name
)

func MatchDomain(query string) (string, error) {
	var matchedDomain string

	// Get the list of all domains
	listDomain, err := provider.ListAllDomain()
	if err != nil {
		return "", err
	}

	// Match the query with each domain in the list
	for _, domain := range listDomain {
		assetIdentifier := domain.AssetIdentifier
		if strings.Contains(query, assetIdentifier) || strings.Contains(assetIdentifier, query) {
			if isExactTLDMatch(query, assetIdentifier) && (matchedDomain == "" || len(assetIdentifier) < len(matchedDomain)) {
				matchedDomain = assetIdentifier
			}
		}
	}

	return matchedDomain, nil
}

func isExactTLDMatch(query, domain string) bool {
	queryParts := strings.Split(query, ".")
	domainParts := strings.Split(domain, ".")
	return len(queryParts) > 1 && len(domainParts) > 1 && queryParts[len(queryParts)-1] == domainParts[len(domainParts)-1]
}
