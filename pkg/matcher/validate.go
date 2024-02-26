package matcher

import "strings"

func MatchDomain(query string, listDomain []string) string {
	var matchedDomain string

	// Mencocokkan query dengan setiap domain dalam list
	for _, domain := range listDomain {
		if strings.Contains(query, domain) || strings.Contains(domain, query) {
			if isExactTLDMatch(query, domain) && (matchedDomain == "" || len(domain) < len(matchedDomain)) {
				matchedDomain = domain
			}
		}
	}

	return matchedDomain
}

func isExactTLDMatch(query, domain string) bool {
	queryParts := strings.Split(query, ".")
	domainParts := strings.Split(domain, ".")
	return len(queryParts) > 1 && len(domainParts) > 1 && queryParts[len(queryParts)-1] == domainParts[len(domainParts)-1]
}
