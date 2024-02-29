package matcher

import (
	"strings"

	"github.com/syhbt/fynd/pkg/provider" // Replace with your module path
)

// MatchDomains matches each query against the list of domains.
// It returns a map containing only the matched domains.
func MatchDomains(dataQuery []string) (map[string]string, error) {
	matchedDomain := make(map[string]string)

	// Call the ListAllDomains function
	listDomain, err := provider.ListAllDomains()
	if err != nil {
		return nil, err
	}

	for _, query := range dataQuery {
		for _, domain := range listDomain {
			if strings.HasPrefix(domain, "*.") {
				wildcard := strings.TrimPrefix(domain, "*.")
				if strings.HasSuffix(query, wildcard) {
					matchedDomain[query] = query
					break
				}
			} else if strings.Contains(domain, query) {
				if query == domain || strings.HasSuffix(domain, "."+query) {
					matchedDomain[query] = query
				}
				break
			}
		}
	}

	return matchedDomain, nil
}
