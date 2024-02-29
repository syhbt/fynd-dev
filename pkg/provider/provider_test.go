package provider

import (
	"testing"
)

func TestListAllDomains(t *testing.T) {
	domains, err := ListAllDomains()
	if err != nil {
		t.Errorf("ListAllDomains() error = %v", err)
		return
	}

	if len(domains) == 0 {
		t.Error("ListAllDomains() should return at least one domain, got 0")
	}

	t.Logf("ListAllDomains() returned %d domains", len(domains))
}
