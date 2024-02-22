package pihole

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	// "github.com/prometheus/common/log"

	"log/slog"
	"net/http"
	"net/url"

	"pihole_exporter/config"
)

const (
	// Request Accept header value
	acceptHeader = "application/json"
	// Request Content-Type header value
	mediaType = "application/json"
)

var (
	userAgent = fmt.Sprintf("pihole-exporter/%s", config.Version)
	// Errors
	errClientIsNil          = errors.New("client is nil")
	errInvalidPiHoleAddress = errors.New("invalid pihole address")
	// Pihole API query variables
	piholeAPIQueryVariables = []string{
		"summaryRaw",
		"overTimeData",
		"topItems",
		"recentItems",
		"getQueryTypes",
		"getForwardDestinations",
		"getQuerySources",
		"jsonForceObject",
	}
)

// Client holds PiHole client configuration.
// It can query the remote PiHole server for data and return it
// as *Metrics.
type Client struct {
	client     *http.Client
	metricsURL *url.URL
	basicAuth  *config.BasicAuth
}

const (
	// Default number of different results returned.
	defaultQueryResults = 30
)

// NewClient returns new *Client from *config.Pihole settings.
func NewClient(cfg *config.Pihole) (*Client, error) {
	slog.Debug("Setting up PiHole client")

	if !strings.HasPrefix(cfg.Listen, "http://") &&
		!strings.HasPrefix(cfg.Listen, "https://") &&
		!strings.HasPrefix(cfg.Listen, "unix://") {
		if cfg.TLS != nil && cfg.TLS.CACertificate != "" {
			cfg.Listen = "https://" + cfg.Listen
		}
	}

	metricsURL, err := url.Parse(strings.Trim(cfg.Listen, "/") + cfg.GetAPIPath())
	metricsQuery := metricsURL.Query()
	for _, queryVar := range piholeAPIQueryVariables {
		if cfg.NumResults > 0 {
			metricsQuery.Set(queryVar, strconv.FormatInt(cfg.NumResults, 10))
		} else {
			metricsQuery.Set(queryVar, strconv.Itoa(defaultQueryResults))
		}
	}
	if cfg.APIToken != "" {
		metricsQuery.Set("auth", cfg.APIToken)
	}
	metricsURL.RawQuery = metricsQuery.Encode()

	if err != nil || (metricsURL.Scheme != "http" && metricsURL.Scheme != "https") {
		if err != nil {
			return nil, err
		}
		return nil, errInvalidPiHoleAddress
	}

	client := http.DefaultClient

	if cfg.TLS != nil {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: cfg.TLS.Insecure,
			},
		}

		if cfg.TLS.CACertificate != "" {
			caCert, err := os.ReadFile(cfg.TLS.CACertificate)
			if err != nil {
				slog.Error("Error reading CA certificate", "error", err)
			}
			caCertPool := x509.NewCertPool()
			caCertPool.AppendCertsFromPEM(caCert)
			tr.TLSClientConfig.RootCAs = caCertPool
		}

		client = &http.Client{Transport: tr}
	}

	slog.Debug("Set up PiHole client")

	return &Client{
		client:     client,
		basicAuth:  cfg.BasicAuth,
		metricsURL: metricsURL,
	}, nil
}

// setupHeaders sets up Content-Type, Accept and User-Agent headers
// in the provided request. If basic authentication is enabled,
// it also adds it to the request.
func (c *Client) setupHeaders(request *http.Request) {
	request.Header.Add("Content-Type", mediaType)
	request.Header.Add("Accept", acceptHeader)
	request.Header.Add("User-Agent", userAgent)

	if c.basicAuth.IsBasicAuth() {
		request.SetBasicAuth(c.basicAuth.Username, c.basicAuth.Password)
	}
}

// GetMetrics sends a request to the remote PiHole API endpoint
// and returns the data as *Metrics.
func (client *Client) GetMetrics() (*Metrics, error) {
	slog.Debug("Getting metrics", "url", client.metricsURL.String())

	if client.client == nil {
		return nil, errClientIsNil
	}

	piholeRequest, err := http.NewRequest("GET", client.metricsURL.String(), nil)
	if err != nil {
		return nil, err
	}
	client.setupHeaders(piholeRequest)
	resp, err := client.client.Do(piholeRequest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var metrics Metrics

	dec := json.NewDecoder(bytes.NewBuffer(body))
	if err := dec.Decode(&metrics); err != nil {
		return nil, err
	}

	return &metrics, nil
}
