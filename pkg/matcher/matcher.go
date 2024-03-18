package matcher

import (
	"strings"

	"github.com/syhbt/fynd/pkg/parser"
)

type MatchResult struct {
	Query  string
	Domain string
	URL    string
	Found  bool
}

type DomainData struct {
	Domain string
	URL    string
}

func MatchDomains(query []string, data []parser.FyndData, results chan MatchResult) {
    domainMap := make(map[string]DomainData)
    for _, d := range data {
        for _, domain := range d.Domains {
            domainMap[domain] = DomainData{Domain: domain, URL: d.URL}
        }
    }

    for _, q := range query {
        if domainData, ok := domainMap[q]; ok {
            results <- MatchResult{
                Query:  q,
                Domain: domainData.Domain,
                URL:    domainData.URL,
            }
        } else {
            found := false
            for domain, domainData := range domainMap {
                if isSubdomain(q, domain) || isSubdomain(domain, q) {
                    results <- MatchResult{
                        Query:  q,
                        Domain: domainData.Domain,
                        URL:    domainData.URL,
                        Found:  true,
                    }
                    found = true
                    break
                }
            }
            if !found {
                results <- MatchResult{
                    Query: q,
                    Found: false,
                }
            }
        }
    }

    close(results)
}

func isSubdomain(domain1, domain2 string) bool {
	parts1 := strings.Split(domain1, ".")
	parts2 := strings.Split(domain2, ".")

	if len(parts1) < len(parts2) {
		parts1, parts2 = parts2, parts1
	}

	// If the shorter domain has only two parts, the longer domain must have more than two parts
	// and the last two parts of both domains must be the same
	if len(parts2) == 2 && len(parts1) > 2 {
		if parts1[len(parts1)-1] != parts2[len(parts2)-1] || parts1[len(parts1)-2] != parts2[len(parts2)-2] {
			return false
		}
	} else if len(parts2) < 2 {
		return false
	}

	for i, j := len(parts1)-1, len(parts2)-1; j >= 0; i, j = i-1, j-1 {
		if parts1[i] != parts2[j] {
			return false
		}
	}

	return true
}
