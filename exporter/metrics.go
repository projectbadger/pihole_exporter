package exporter

import "github.com/prometheus/client_golang/prometheus"

const (
	// Metrics namespace
	namespace = "pihole"
	// Metrics names
	metricDomainsBeingBlockedName  = "domains_being_blocked"
	metricDNSQueriesTodayName      = "dns_queries_today"
	metricAdsBlockedName           = "ads_blocked"
	metricAdsPercentageName        = "ads_percentage"
	metricUniqueDomainsName        = "unique_domains"
	metricQueriesForwardedName     = "queries_forwarded"
	metricQueriesCachedName        = "queries_cached"
	metricClientsEverSeenName      = "clients_ever_seen"
	metricUniqueClientsName        = "unique_clients"
	metricDNSQueriesAllTypesName   = "dns_queries_all_types"
	metricDNSQueriesAllRepliesName = "dns_queries_all_replies"
	metricPrivacyLevelName         = "privacy_level"
	// metricStatusName               = "status"
	metricTopQueriesName          = "top_queries"
	metricTopAdsName              = "top_ads"
	metricTopSourcesName          = "top_sources"
	metricForwardDestinationsName = "forward_destinations"
	metricQueryTypesName          = "query_types"
	metricRepliesName             = "replies"
	metricGravityLastUpdatedName  = "gravity_last_updated"
	// Metrics descriptions
	metricDomainsBeingBlockedDescription  = "Domains being blocked today"
	metricDNSQueriesTodayDescription      = "DNS queries today"
	metricAdsBlockedDescription           = "Ads blocked today"
	metricAdsPercentageDescription        = "Ads percentage today"
	metricUniqueDomainsDescription        = "Unique domains"
	metricQueriesForwardedDescription     = "Forwarded queries"
	metricQueriesCachedDescription        = "Cached queries"
	metricClientsEverSeenDescription      = "Number of different clients ever"
	metricUniqueClientsDescription        = "Number of unique clients"
	metricDNSQueriesAllTypesDescription   = "DNS queries all types"
	metricDNSQueriesAllRepliesDescription = "DNS queries all replies"
	metricPrivacyLevelDescription         = "Privacy level"
	// metricStatusDescription               = "Status"
	metricTopQueriesDescription          = "Top queries"
	metricTopAdsDescription              = "Top ads"
	metricTopSourcesDescription          = "Top sources"
	metricForwardDestinationsDescription = "Forward destinations"
	metricQueryTypesDescription          = "DNS query types"
	metricRepliesDescription             = "DNS Replies"
	metricGravityLastUpdatedDescription  = "gravity_last_updated"
)

// metrics holds all the prometheus.Desc metrics descriptions.
type metrics struct {
	domainsBeingBlocked  *prometheus.Desc
	dnsQueries           *prometheus.Desc
	adsBlocked           *prometheus.Desc
	adsPercentage        *prometheus.Desc
	uniqueDomains        *prometheus.Desc
	queriesForwarded     *prometheus.Desc
	queriesCached        *prometheus.Desc
	clientsEverSeen      *prometheus.Desc
	uniqueClients        *prometheus.Desc
	dnsQueriesAllTypes   *prometheus.Desc
	dnsQueriesAllReplies *prometheus.Desc
	privacyLevel         *prometheus.Desc
	// status               *prometheus.Desc
	topQueries          *prometheus.Desc
	topAds              *prometheus.Desc
	topSources          *prometheus.Desc
	forwardDestinations *prometheus.Desc
	queryTypes          *prometheus.Desc
	replies             *prometheus.Desc
	gravityLastUpdated  *prometheus.Desc
}

// NewMetrics returns a new *metrics struct holding all of
// the prometheus.Desc metrics descriptions.
func NewMetrics() *metrics {
	return &metrics{
		domainsBeingBlocked: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricDomainsBeingBlockedName),
			metricDomainsBeingBlockedDescription,
			nil, nil,
		),
		dnsQueries: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricDNSQueriesTodayName),
			metricDNSQueriesTodayDescription,
			nil, nil,
		),
		adsBlocked: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricAdsBlockedName),
			metricAdsBlockedDescription,
			nil, nil,
		),
		adsPercentage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricAdsPercentageName),
			metricAdsPercentageDescription,
			nil, nil,
		),
		uniqueDomains: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricUniqueDomainsName),
			metricUniqueDomainsDescription,
			nil, nil,
		),
		queriesForwarded: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricQueriesForwardedName),
			metricQueriesForwardedDescription,
			nil, nil,
		),
		queriesCached: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricQueriesCachedName),
			metricQueriesCachedDescription,
			nil, nil,
		),
		clientsEverSeen: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricClientsEverSeenName),
			metricClientsEverSeenDescription,
			nil, nil,
		),
		uniqueClients: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricUniqueClientsName),
			metricUniqueClientsDescription,
			nil, nil,
		),
		dnsQueriesAllTypes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricDNSQueriesAllTypesName),
			metricDNSQueriesAllTypesDescription,
			nil, nil,
		),
		dnsQueriesAllReplies: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricDNSQueriesAllRepliesName),
			metricDNSQueriesAllRepliesDescription,
			nil, nil,
		),
		privacyLevel: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricPrivacyLevelName),
			metricPrivacyLevelDescription,
			nil, nil,
		),
		gravityLastUpdated: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricGravityLastUpdatedName),
			metricGravityLastUpdatedDescription,
			nil, nil,
		),
		// status: prometheus.NewDesc(
		// 	prometheus.BuildFQName(namespace, "", metricStatusName),
		// 	metricStatusDescription,
		// 	nil, nil,
		// ),
		topQueries: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricTopQueriesName),
			metricTopQueriesDescription,
			[]string{"domain"}, nil,
		),
		topAds: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricTopAdsName),
			metricTopAdsDescription,
			[]string{"domain"}, nil,
		),
		topSources: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricTopSourcesName),
			metricTopSourcesDescription,
			[]string{"client"}, nil,
		),
		forwardDestinations: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricForwardDestinationsName),
			metricForwardDestinationsDescription,
			[]string{"destination"}, nil,
		),
		queryTypes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricQueryTypesName),
			metricQueryTypesDescription,
			[]string{"type"}, nil,
		),
		replies: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", metricRepliesName),
			metricRepliesDescription,
			[]string{"reply"}, nil,
		),
	}
}

// Describe publishes all of the collected PiHole metrics to the
// provided channel.
func (m *metrics) Describe(ch chan<- *prometheus.Desc) {
	if m == nil {
		return
	}
	ch <- m.domainsBeingBlocked
	ch <- m.dnsQueries
	ch <- m.adsBlocked
	ch <- m.adsPercentage
	ch <- m.uniqueDomains
	ch <- m.queriesForwarded
	ch <- m.queriesCached
	ch <- m.clientsEverSeen
	ch <- m.uniqueClients
	ch <- m.dnsQueriesAllTypes
	ch <- m.dnsQueriesAllReplies
	ch <- m.privacyLevel
	ch <- m.gravityLastUpdated
	// ch <- m.status
	ch <- m.topQueries
	ch <- m.topAds
	ch <- m.topSources
	ch <- m.queryTypes
	ch <- m.forwardDestinations
}
