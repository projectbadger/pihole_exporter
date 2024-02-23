package exporter

import (
	"pihole_exporter/config"
	"testing"
)

func TestNewMetrics(t *testing.T) {
	var m *metrics
	m.Describe(nil)
	m = NewMetrics()
	if m.domainsBeingBlocked == nil ||
		m.dnsQueries == nil ||
		m.adsBlocked == nil ||
		m.adsPercentage == nil ||
		m.uniqueDomains == nil ||
		m.queriesForwarded == nil ||
		m.queriesCached == nil ||
		m.clientsEverSeen == nil ||
		m.uniqueClients == nil ||
		m.dnsQueriesAllTypes == nil ||
		m.dnsQueriesAllReplies == nil ||
		m.privacyLevel == nil ||
		m.status == nil ||
		m.topQueries == nil ||
		m.topAds == nil ||
		m.topSources == nil ||
		m.forwardDestinations == nil ||
		m.queryTypes == nil ||
		m.replies == nil ||
		m.gravityLastUpdated == nil {
		t.Error("metric is nil")
	}
}

func TestExporter(t *testing.T) {
	t.Run("NewExporter", func(t *testing.T) {
		var cfg *config.Config
		_, err := NewExporter(cfg)
		if err == nil {
			t.Errorf("unexpected nil error")
		}
		cfg = &config.Config{Pihole: &config.Pihole{
			ListenAddress: "pi.hole", // Will be replaced
			APIPath:       "/admin/api.php",
			APIToken:      "token",
			NumResults:    5,
			BasicAuth: &config.BasicAuth{
				Username: "user",
				Password: "pass",
			},
		}}
		e, err := NewExporter(cfg)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		if e.metrics == nil {
			t.Error("metrics struct is nil")
		}
		// e.Describe(nil)
	})
}
