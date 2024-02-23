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
	errMetricsURLIsNil      = errors.New("metrics url is nil")
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

	cfg.ListenAddress = getPiholeListenAddress(cfg)

	metricsURL, err := getMetricsURL(cfg)
	if err != nil {
		return nil, err
	}

	client := getHTTPClient(cfg)

	slog.Debug("Set up PiHole client")

	return &Client{
		client:     client,
		basicAuth:  cfg.BasicAuth,
		metricsURL: metricsURL,
	}, nil
}

// getPiholeListenAddress returns PiHole server listener address.
func getPiholeListenAddress(cfg *config.Pihole) string {
	if cfg.ListenAddress == "" {
		return ""
	}
	if _, err := url.Parse(cfg.ListenAddress); err != nil {
		return ""
	}
	if !strings.HasPrefix(cfg.ListenAddress, "http://") &&
		!strings.HasPrefix(cfg.ListenAddress, "https://") &&
		!strings.HasPrefix(cfg.ListenAddress, "unix://") {
		if cfg.TLS != nil && cfg.IsTLS() {
			return "https://" + cfg.ListenAddress
		} else {
			return "http://" + cfg.ListenAddress
		}
	}
	return cfg.ListenAddress
}

// getMetricsURL returns the whole PiHole server API request URL.
func getMetricsURL(cfg *config.Pihole) (*url.URL, error) {
	metricsURL, err := url.Parse(strings.Trim(cfg.ListenAddress, "/") + cfg.GetAPIPath())

	if err != nil || (metricsURL.Scheme != "http" && metricsURL.Scheme != "https") {
		if err != nil {
			return nil, err
		}
		return nil, errInvalidPiHoleAddress
	}

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

	return metricsURL, nil
}

// getHTTPClient returns a *http.Client, configured by the provided
// *config.Pihole config.
func getHTTPClient(cfg *config.Pihole) *http.Client {
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
	return client
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

func (c *Client) getPiholeMetricsRequest() (*http.Request, error) {
	if c.metricsURL == nil {
		return nil, errMetricsURLIsNil
	}
	piholeRequest, err := http.NewRequest("GET", c.metricsURL.String(), nil)
	if err != nil {
		return nil, err
	}
	c.setupHeaders(piholeRequest)
	return piholeRequest, nil
}

// GetMetrics sends a request to the remote PiHole API endpoint
// and returns the data as *Metrics.
func (client *Client) GetMetrics() (*Metrics, error) {
	slog.Debug("Getting metrics", "url", client.metricsURL.String())

	if client.client == nil {
		return nil, errClientIsNil
	}

	piholeRequest, err := client.getPiholeMetricsRequest()
	if err != nil {
		return nil, err
	}
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
