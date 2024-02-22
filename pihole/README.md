
# pihole

```go
import pihole_exporter/pihole
```

## PiHole

PiHole package provides the PiHole client for making requests to the remote PiHole server.

It returns metrics from the PiHole admin API endpoint.

## Index

- [type Client](#type-client)
  - [NewClient() (Client, error)](#func-newclient-client-error)
  - [GetMetrics() (Metrics, error)](#func-client-getmetrics-metrics-error)
- [type GravityLastUpdated](#type-gravitylastupdated)
- [type Metrics](#type-metrics)
  - [GetRepliesMap()](#func-metrics-getrepliesmap)
- [type TimeRelative](#type-timerelative)


## type [Client](<pihole.go#L51>)

Client holds PiHole client configuration.
It can query the remote PiHole server for data and return it
as *Metrics.
```go
type Client struct {
	// contains filtered or unexported fields
}
```

## func [NewClient() (Client, error)](<pihole.go#L63>)

NewClient returns new *Client from *config.Pihole settings.


```go
func NewClient(cfg *config.Pihole) (*Client, error)
```

## func (*Client) [GetMetrics() (Metrics, error)](<pihole.go#L138>)

GetMetrics sends a request to the remote PiHole API endpoint
and returns the data as *Metrics.


```go
func (client *Client) GetMetrics() (*Metrics, error)
```

## type [GravityLastUpdated](<metrics.go#L63>)

GravityLastUpdated holds gravity_last_updated data from
PiHole statistics
```go
type GravityLastUpdated struct {
	FileExists	bool		`json:"file_exists,omitempty"`
	Absolute	int64		`json:"absolute,omitempty"`
	Relative	TimeRelative	`json:"relative,omitempty"`
}
```

## type [Metrics](<metrics.go#L4>)

Metrics define PiHome Prometheus metrics
```go
type Metrics struct {
	DomainsBeingBlocked	float64			`json:"domains_being_blocked"`
	DNSQueriesToday		float64			`json:"dns_queries_today"`
	AdsBlockedToday		float64			`json:"ads_blocked_today"`
	AdsPercentageToday	float64			`json:"ads_percentage_today"`
	UniqueDomains		float64			`json:"unique_domains,omitempty"`
	QueriesForwarded	float64			`json:"queries_forwarded,omitempty"`
	QueriesCached		float64			`json:"queries_cached,omitempty"`
	ClientsEverSeen		float64			`json:"clients_ever_seen,omitempty"`
	UniqueClients		float64			`json:"unique_clients,omitempty"`
	DNSQueriesAllTypes	float64			`json:"dns_queries_all_types,omitempty"`
	ReplyUNKNOWN		float64			`json:"reply_UNKNOWN,omitempty"`
	ReplyNODATA		float64			`json:"reply_NODATA,omitempty"`
	ReplyNXDOMAIN		float64			`json:"reply_NXDOMAIN,omitempty"`
	ReplyCNAME		float64			`json:"reply_CNAME,omitempty"`
	ReplyIP			float64			`json:"reply_IP,omitempty"`
	ReplyDomain		float64			`json:"reply_DOMAIN,omitempty"`
	ReplyRRNAME		float64			`json:"reply_RRNAME,omitempty"`
	ReplySERVFAIL		float64			`json:"reply_SERVFAIL,omitempty"`
	ReplyREFUSED		float64			`json:"reply_REFUSED,omitempty"`
	ReplyOTHER		float64			`json:"reply_OTHER,omitempty"`
	ReplyDNSSEC		float64			`json:"reply_DNSSEC,omitempty"`
	ReplyNONE		float64			`json:"reply_NONE,omitempty"`
	ReplyBLOB		float64			`json:"reply_BLOB,omitempty"`
	DNSQueriesAllReplies	float64			`json:"dns_queries_all_replies,omitempty"`
	PrivacyLevel		float64			`json:"privacy_level,omitempty"`
	GravityLastUpdated	GravityLastUpdated	`json:"gravity_last_updated,omitempty"`
	Status			string			`json:"status,omitempty"`
	TopQueries		map[string]float64	`json:"top_queries"`
	TopAds			map[string]float64	`json:"top_ads"`
	TopSources		map[string]float64	`json:"top_sources"`
	QueryTypes		map[string]float64	`json:"querytypes,omitempty"`
	ForwardDestinations	map[string]float64	`json:"forward_destinations,omitempty"`
}
```

## func (*Metrics) [GetRepliesMap()](<metrics.go#L41>)

GetRepliesMap returns a map of numbers of replies of
different types


```go
func (m *Metrics) GetRepliesMap() map[string]float64
```

## type [TimeRelative](<metrics.go#L69>)
```go
type TimeRelative struct {
	Days	int	`json:"days"`
	Hours	int	`json:"hours"`
	Minutes	int	`json:"minutes"`
}
```

