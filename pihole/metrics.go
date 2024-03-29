package pihole

// Metrics define PiHome Prometheus metrics
type Metrics struct {
	DomainsBeingBlocked  float64            `json:"domains_being_blocked"`
	DNSQueriesToday      float64            `json:"dns_queries_today"`
	AdsBlockedToday      float64            `json:"ads_blocked_today"`
	AdsPercentageToday   float64            `json:"ads_percentage_today"`
	UniqueDomains        float64            `json:"unique_domains,omitempty"`
	QueriesForwarded     float64            `json:"queries_forwarded,omitempty"`
	QueriesCached        float64            `json:"queries_cached,omitempty"`
	ClientsEverSeen      float64            `json:"clients_ever_seen,omitempty"`
	UniqueClients        float64            `json:"unique_clients,omitempty"`
	DNSQueriesAllTypes   float64            `json:"dns_queries_all_types,omitempty"`
	ReplyUNKNOWN         float64            `json:"reply_UNKNOWN,omitempty"`
	ReplyNODATA          float64            `json:"reply_NODATA,omitempty"`
	ReplyNXDOMAIN        float64            `json:"reply_NXDOMAIN,omitempty"`
	ReplyCNAME           float64            `json:"reply_CNAME,omitempty"`
	ReplyIP              float64            `json:"reply_IP,omitempty"`
	ReplyDOMAIN          float64            `json:"reply_DOMAIN,omitempty"`
	ReplyRRNAME          float64            `json:"reply_RRNAME,omitempty"`
	ReplySERVFAIL        float64            `json:"reply_SERVFAIL,omitempty"`
	ReplyREFUSED         float64            `json:"reply_REFUSED,omitempty"`
	ReplyOTHER           float64            `json:"reply_OTHER,omitempty"`
	ReplyDNSSEC          float64            `json:"reply_DNSSEC,omitempty"`
	ReplyNONE            float64            `json:"reply_NONE,omitempty"`
	ReplyBLOB            float64            `json:"reply_BLOB,omitempty"`
	DNSQueriesAllReplies float64            `json:"dns_queries_all_replies,omitempty"`
	PrivacyLevel         float64            `json:"privacy_level,omitempty"`
	GravityLastUpdated   GravityLastUpdated `json:"gravity_last_updated,omitempty"`
	Status               string             `json:"status,omitempty"`
	TopQueries           map[string]float64 `json:"top_queries"`
	TopAds               map[string]float64 `json:"top_ads"`
	TopSources           map[string]float64 `json:"top_sources"`
	QueryTypes           map[string]float64 `json:"querytypes,omitempty"`
	ForwardDestinations  map[string]float64 `json:"forward_destinations,omitempty"`
}

// GetRepliesMap returns a map of numbers of replies of
// different types
func (m *Metrics) GetRepliesMap() map[string]float64 {
	if m == nil {
		return nil
	}
	replies := make(map[string]float64)
	replies["UNKNOWN"] = m.ReplyUNKNOWN
	replies["NODATA"] = m.ReplyNODATA
	replies["NXDOMAIN"] = m.ReplyNXDOMAIN
	replies["CNAME"] = m.ReplyCNAME
	replies["IP"] = m.ReplyIP
	replies["DOMAIN"] = m.ReplyDOMAIN
	replies["RRNAME"] = m.ReplyRRNAME
	replies["SERVFAIL"] = m.ReplySERVFAIL
	replies["REFUSED"] = m.ReplyREFUSED
	replies["OTHER"] = m.ReplyOTHER
	replies["DNSSEC"] = m.ReplyDNSSEC
	replies["NONE"] = m.ReplyNONE
	replies["BLOB"] = m.ReplyBLOB
	return replies
}

// GravityLastUpdated holds gravity_last_updated data from
// PiHole statistics
type GravityLastUpdated struct {
	FileExists bool         `json:"file_exists,omitempty"`
	Absolute   int64        `json:"absolute,omitempty"`
	Relative   TimeRelative `json:"relative,omitempty"`
}

type TimeRelative struct {
	Days    int `json:"days"`
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
}
