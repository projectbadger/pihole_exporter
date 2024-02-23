package exporter

import (
	"errors"
	"log/slog"

	"github.com/prometheus/client_golang/prometheus"

	"pihole_exporter/config"
	"pihole_exporter/pihole"
)

var (
	errInvalidConfig = errors.New("invalid config")
)

// Exporter collects Pihole metrics from the given address and
// exports them using the Prometheus metrics package.
type Exporter struct {
	// PiHole client for calling the admin API
	Pihole  *pihole.Client
	metrics *metrics
}

// NewExporter returns an initialized Exporter.
func NewExporter(cfg *config.Config) (*Exporter, error) {
	if cfg == nil || cfg.Pihole == nil {
		return nil, errInvalidConfig
	}
	slog.Info("Seting up Pihole exporter", "url", cfg.Pihole.ListenAddress)
	pihole, err := pihole.NewClient(cfg.Pihole)
	if err != nil {
		return nil, err
	}
	return &Exporter{
		Pihole:  pihole,
		metrics: NewMetrics(),
	}, nil
}

// Describe publishes all of the collected PiHole metrics to the
// provided channel by calling the underlying *metrics.Describe(ch).
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	if e.metrics == nil {
		return
	}
	e.metrics.Describe(ch)
}

// Collect collects the metrics from the channel and sends them
// as Prometheus metrics.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	slog.Info("Starting PiHole exporter")
	resp, err := e.Pihole.GetMetrics()
	if err != nil {
		slog.Error("Error getting response from PiHole", "error", err.Error())
		return
	}
	slog.Debug("PiHole metrics response", "metrics", resp)
	e.sendMetric(ch, resp.DomainsBeingBlocked, e.metrics.domainsBeingBlocked)
	e.sendMetric(ch, resp.DNSQueriesToday, e.metrics.dnsQueries)
	e.sendMetric(ch, resp.AdsBlockedToday, e.metrics.adsBlocked)
	e.sendMetric(ch, resp.AdsPercentageToday, e.metrics.adsPercentage)
	e.sendMetric(ch, resp.UniqueDomains, e.metrics.uniqueDomains)
	e.sendMetric(ch, resp.QueriesForwarded, e.metrics.queriesForwarded)
	e.sendMetric(ch, resp.QueriesCached, e.metrics.queriesCached)
	e.sendMetric(ch, resp.ClientsEverSeen, e.metrics.clientsEverSeen)
	e.sendMetric(ch, resp.UniqueClients, e.metrics.uniqueClients)
	e.sendMetric(ch, resp.DNSQueriesAllTypes, e.metrics.dnsQueriesAllTypes)
	e.sendMetric(ch, resp.DNSQueriesAllReplies, e.metrics.dnsQueriesAllReplies)
	e.sendMetric(ch, resp.PrivacyLevel, e.metrics.privacyLevel)
	var status float64 = 0
	if resp.Status == "enabled" {
		status = 1
	}
	e.sendMetric(ch, status, e.metrics.status)
	e.sendMetric(ch, float64(resp.GravityLastUpdated.Absolute), e.metrics.gravityLastUpdated)
	for domain, hits := range resp.TopQueries {
		e.sendMetric(ch, hits, e.metrics.topQueries, domain)
	}
	for domain, hits := range resp.TopAds {
		e.sendMetric(ch, hits, e.metrics.topAds, domain)
	}
	for client, requests := range resp.TopSources {
		e.sendMetric(ch, requests, e.metrics.topSources, client)
	}
	for queryType, requests := range resp.QueryTypes {
		e.sendMetric(ch, requests, e.metrics.queryTypes, queryType)
	}
	for destination, requests := range resp.ForwardDestinations {
		e.sendMetric(ch, requests, e.metrics.forwardDestinations, destination)
	}
	repliesMap := resp.GetRepliesMap()
	for rep, requests := range repliesMap {
		e.sendMetric(ch, requests, e.metrics.replies, rep)
	}

	slog.Info("Pihole exporter finished")
}

// sendMetric sends a new prometheus.Metric to the provided channel.
func (e *Exporter) sendMetric(ch chan<- prometheus.Metric, value float64, desc *prometheus.Desc, labels ...string) {
	ch <- prometheus.MustNewConstMetric(
		desc, prometheus.GaugeValue, value, labels...)
}

// Register registers metrics in the prometheus package.
func (e *Exporter) Register() {
	slog.Debug("Registering metrics.")
	if e == nil {
		return
	}
	prometheus.MustRegister(e)
	slog.Debug("Registered metrics.")
}
