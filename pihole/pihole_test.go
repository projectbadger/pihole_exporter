package pihole

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"pihole_exporter/config"
	"pihole_exporter/test_data"
	"testing"
)

func TestClient(t *testing.T) {
	t.Run("NumResults", func(t *testing.T) {
		// client := &Client{
		// 	Username: testData.InputStr[0],
		// 	Password: testData.InputStr[1],
		// }
		// isBasicAuth := client.IsBasicAuth()
		// if isBasicAuth && !testData.ExpectedBool[0] {
		// 	t.Error("returned IsBasicAuth is true when should be false")
		// } else if !isBasicAuth && testData.ExpectedBool[0] {
		// 	t.Error("returned IsBasicAuth is false when should be true")
		// }
	})
	for _, testData := range test_data.IsBasicAuthStr {
		t.Run(testData.Name, func(t *testing.T) {
			// client := &Client{
			// 	Username: testData.InputStr[0],
			// 	Password: testData.InputStr[1],
			// }
			// isBasicAuth := client.IsBasicAuth()
			// if isBasicAuth && !testData.ExpectedBool[0] {
			// 	t.Error("returned IsBasicAuth is true when should be false")
			// } else if !isBasicAuth && testData.ExpectedBool[0] {
			// 	t.Error("returned IsBasicAuth is false when should be true")
			// }
		})
	}
}

func TestGetMetricsURL(t *testing.T) {
	for _, testData := range test_data.MetricsURL {
		t.Run(testData.Name, func(t *testing.T) {
			cfg := &config.Pihole{
				ListenAddress: testData.InputStr[0],
				APIPath:       testData.InputStr[1],
				APIToken:      testData.InputStr[3],
				NumResults:    int64(testData.InputInt[0]),
				TLS: &config.TLS{
					Insecure:      testData.InputBool[0],
					CACertificate: testData.InputStr[2],
				},
			}
			cfg.ListenAddress = getPiholeListenAddress(cfg)

			metricsURL, err := getMetricsURL(cfg)
			if err != nil {
				if testData.ExpectedBool[0] {
					// Expected error
				} else {
					t.Errorf("unexpected error: %s", err)
				}
				return // Test finished
			} else if testData.ExpectedBool[0] {
				// Exšected error, but got nil
				t.Errorf("expected error but got nil")
			}
			if metricsURL.String() != testData.ExpectedStr[0] {
				t.Errorf("metrics URLs do not match: expected: '%s', got: '%s'", testData.ExpectedStr[0], metricsURL.String())
			}
		})
	}
}

func TestGetPiholeMetricsRequest(t *testing.T) {
	t.Run("metricsURLInvalid", func(t *testing.T) {
		var err error

		metricsURL, err := getMetricsURL(&config.Pihole{ListenAddress: "@invalid:/pi.hole"})
		if err == nil || metricsURL != nil {
			t.Errorf("should be invalid: %s", metricsURL)
		}
	})
	t.Run("getPiholeMetricsRequestInvalid", func(t *testing.T) {
		var err error

		metricsURL, err := getMetricsURL(&config.Pihole{ListenAddress: "@invalid:/pi.hole"})
		if err == nil || metricsURL != nil {
			t.Errorf("should be invalid: %s", metricsURL)
		}
	})
	t.Run("ClientMetricsURLInvalid", func(t *testing.T) {
		var err error
		client, err := NewClient(&config.Pihole{ListenAddress: "@invalid:/pi.hole"})
		if err == nil || client != nil {
			t.Errorf("should be invalid: %#v", client)
		}
	})
	t.Run("MetricsURL", func(t *testing.T) {
		for _, testData := range test_data.MetricsURL {
			t.Run(testData.Name, func(t *testing.T) {
				cfg := &config.Pihole{
					ListenAddress: testData.InputStr[0],
					APIPath:       testData.InputStr[1],
					APIToken:      testData.InputStr[3],
					NumResults:    int64(testData.InputInt[0]),
					BasicAuth:     &config.BasicAuth{},
					TLS: &config.TLS{
						Insecure:      testData.InputBool[0],
						CACertificate: testData.InputStr[2],
					},
				}
				cfg.ListenAddress = getPiholeListenAddress(cfg)
				if len(testData.InputStr) > 4 {
					// Has basic auth
					cfg.BasicAuth = &config.BasicAuth{
						Username: testData.InputStr[4],
						Password: testData.InputStr[5],
					}
				}

				var metricsURL *url.URL
				t.Run("metricsURL", func(t *testing.T) {
					var err error
					metricsURL, err = getMetricsURL(cfg)
					if err != nil {
						if testData.ExpectedBool[0] {
							// Expected error
						} else {
							t.Errorf("unexpected error: %s", err)
						}
						return // Test finished
					} else if testData.ExpectedBool[0] {
						// Exšected error, but got nil
						t.Errorf("expected error but got nil")
					}
					if metricsURL.String() != testData.ExpectedStr[0] {
						t.Errorf("metrics URLs do not match: expected: '%s', got: '%s'", testData.ExpectedStr[0], metricsURL.String())
					}
				})
				if metricsURL == nil {
					return
				}

				httpClient := getHTTPClient(cfg)
				client := &Client{
					client:     httpClient,
					metricsURL: metricsURL,
					basicAuth:  cfg.BasicAuth,
				}
				t.Run("client", func(t *testing.T) {
					c, err := NewClient(cfg)
					if metricsURL == nil {
						// Expect invalid client data
						if err == nil {
							t.Errorf("error is nil")
						}
						if testData.ExpectedStr[0] != "" {
							t.Error("expected non-empty metrics URL")
						}
					}
					if c.basicAuth.Username != cfg.BasicAuth.Username {
						t.Errorf("usernames don't match: expected '%s', got '%s'", cfg.BasicAuth.Username, c.basicAuth.Username)
					}
					if c.basicAuth.Password != cfg.BasicAuth.Password {
						t.Errorf("passwords don't match: expected '%s', got '%s'", cfg.BasicAuth.Password, c.basicAuth.Password)
					}
				})
				t.Run("RequestURL", func(t *testing.T) {
					request, err := client.getPiholeMetricsRequest()
					if err != nil {
						if testData.ExpectedBool[1] {
							// Expected error
						} else {
							t.Errorf("unexpected error: %s", err)
						}
						return // Test finished
					} else if testData.ExpectedBool[1] {
						// Exšected error, but got nil
						t.Errorf("expected error 2 but got nil")
					}
					if metricsURL.String() != testData.ExpectedStr[0] {
						t.Errorf("metrics URLs do not match: expected: '%s', got: '%s'", testData.ExpectedStr[1], request.URL.String())
					}
				})
				t.Run("Headers", func(t *testing.T) {
					request, err := client.getPiholeMetricsRequest()
					if err != nil {
						if testData.ExpectedBool[1] {
							// Expected error
						} else {
							t.Errorf("unexpected error: %s", err)
						}
						return // Test finished
					} else if testData.ExpectedBool[1] {
						// Exšected error, but got nil
						t.Errorf("expected error 2 but got nil")
					}
					userAgent := request.Header.Get("User-Agent")
					if userAgent != "pihole-exporter/dev" {
						t.Errorf("user agent not pihole-exporter/dev: '%s'", userAgent)
					}
					if len(testData.InputStr) > 4 {
						authorization := request.Header.Get("Authorization")
						if authorization != testData.ExpectedStr[1] {
							t.Errorf("authorization headers do not match: expected: '%s', got: '%s', request header: %#v", testData.ExpectedStr[1], authorization, request.Header)
						}
					}

					accept := request.Header.Get("Accept")
					if accept != "application/json" {
						t.Errorf("accept not application/json: '%s'", accept)
					}
					contentType := request.Header.Get("Content-Type")
					if contentType != "application/json" {
						t.Errorf("content type not application/json: '%s'", contentType)
					}
					if metricsURL.String() != testData.ExpectedStr[0] {
						t.Errorf("metrics URLs do not match: expected: '%s', got: '%s'", testData.ExpectedStr[0], request.URL.String())
					}
				})
			})
		}
	})
	t.Run("GetMetrics", func(t *testing.T) {
		metrics := &Metrics{
			DomainsBeingBlocked:  1,
			DNSQueriesToday:      2,
			AdsBlockedToday:      3,
			AdsPercentageToday:   4,
			UniqueDomains:        5,
			QueriesForwarded:     6,
			QueriesCached:        7,
			ClientsEverSeen:      8,
			UniqueClients:        9,
			DNSQueriesAllTypes:   10,
			ReplyUNKNOWN:         11,
			ReplyNODATA:          12,
			ReplyNXDOMAIN:        13,
			ReplyCNAME:           14,
			ReplyIP:              15,
			ReplyDOMAIN:          16,
			ReplyRRNAME:          17,
			ReplySERVFAIL:        18,
			ReplyREFUSED:         19,
			ReplyOTHER:           20,
			ReplyDNSSEC:          21,
			ReplyNONE:            22,
			ReplyBLOB:            23,
			DNSQueriesAllReplies: 24,
			PrivacyLevel:         25,
			GravityLastUpdated: GravityLastUpdated{
				Absolute: 1234567890,
			},
		}
		t.Run("Token", func(t *testing.T) {
			cfg := &config.Pihole{
				ListenAddress: "pi.hole", // Will be replaced
				APIPath:       "/admin/api.php",
				APIToken:      "token",
				NumResults:    5,
			}
			cfg1, server := GetTestPiholeServer(cfg, metrics, t)
			client, err := NewClient(cfg1)
			if err != nil {
				t.Errorf("error getting client: %s", err)
			}
			resp, err := client.GetMetrics()
			if err != nil {
				t.Errorf("error getting metrics: %s", err)
			}
			if resp.AdsBlockedToday != metrics.AdsBlockedToday {
				t.Error("metrics don't match")
			}
			server.Close()
		})
		t.Run("TokenBasicAuth", func(t *testing.T) {
			cfg := &config.Pihole{
				ListenAddress: "pi.hole", // Will be replaced
				APIPath:       "/admin/api.php",
				APIToken:      "token",
				NumResults:    5,
				BasicAuth: &config.BasicAuth{
					Username: "user",
					Password: "pass",
				},
			}
			cfg2, server := GetTestPiholeServer(cfg, metrics, t)
			client, err := NewClient(cfg2)
			if err != nil {
				t.Errorf("error getting client: %s", err)
			}
			resp, err := client.GetMetrics()
			if err != nil {
				t.Errorf("error getting metrics: %s", err)
			}
			if resp.AdsBlockedToday != metrics.AdsBlockedToday {
				t.Error("metrics don't match")
			}
			server.Close()
		})
		t.Run("NilClient", func(t *testing.T) {
			cfg := &config.Pihole{
				ListenAddress: "pi.hole", // Will be replaced
				APIPath:       "/admin/api.php",
				APIToken:      "token",
				NumResults:    5,
			}
			cfg1, server := GetTestPiholeServer(cfg, metrics, t)
			client, err := NewClient(cfg1)
			if err != nil {
				t.Errorf("error getting client: %s", err)
			}
			client.client = nil
			_, err = client.GetMetrics()
			if err == nil {
				t.Error("should have error : client is nil")
			}
			client.client = &http.Client{}
			client.metricsURL, _ = url.Parse("")
			_, err = client.GetMetrics()
			if err == nil {
				t.Error("should have error in getPiholeMetricsRequest")
			}
			server.Close()
		})
	})
}

func TestGetRepliesMap(t *testing.T) {
	var m *Metrics
	rMapNil := m.GetRepliesMap()
	if rMapNil != nil {
		t.Error("map should be nil")
	}
	m = &Metrics{
		ReplyUNKNOWN:  1.1,
		ReplyNODATA:   1.1,
		ReplyNXDOMAIN: 1.1,
		ReplyCNAME:    1.1,
		ReplyIP:       1.1,
		ReplyDOMAIN:   1.1,
		ReplyRRNAME:   1.1,
		ReplySERVFAIL: 1.1,
		ReplyREFUSED:  1.1,
		ReplyOTHER:    1.1,
		ReplyDNSSEC:   1.1,
		ReplyNONE:     1.1,
		ReplyBLOB:     1.1,
	}
	rMap := m.GetRepliesMap()
	if val, ok := rMap["UNKNOWN"]; !ok || val != 1.1 {
		t.Error("UNKNOWN invalid")
	}
	if val, ok := rMap["NODATA"]; !ok || val != 1.1 {
		t.Error("NODATA invalid")
	}
	if val, ok := rMap["NXDOMAIN"]; !ok || val != 1.1 {
		t.Error("NXDOMAIN invalid")
	}
	if val, ok := rMap["CNAME"]; !ok || val != 1.1 {
		t.Error("CNAME invalid")
	}
	if val, ok := rMap["IP"]; !ok || val != 1.1 {
		t.Error("IP invalid")
	}
	if val, ok := rMap["DOMAIN"]; !ok || val != 1.1 {
		t.Error("DOMAIN invalid")
	}
	if val, ok := rMap["RRNAME"]; !ok || val != 1.1 {
		t.Error("RRNAME invalid")
	}
	if val, ok := rMap["SERVFAIL"]; !ok || val != 1.1 {
		t.Error("SERVFAIL invalid")
	}
	if val, ok := rMap["OTHER"]; !ok || val != 1.1 {
		t.Error("OTHER invalid")
	}
	if val, ok := rMap["DNSSEC"]; !ok || val != 1.1 {
		t.Error("DNSSEC invalid")
	}
	if val, ok := rMap["NONE"]; !ok || val != 1.1 {
		t.Error("NONE invalid")
	}
	if val, ok := rMap["BLOB"]; !ok || val != 1.1 {
		t.Error("BLOB invalid")
	}
}

func GetTestPiholeServer(cfg *config.Pihole, response *Metrics, t *testing.T) (*config.Pihole, *httptest.Server) {
	if cfg.APIPath == "" {
		cfg.APIPath = "/admin/api.php"
	}
	if response == nil {
		response = &Metrics{
			DomainsBeingBlocked: 500,
			DNSQueriesToday:     500,
			AdsBlockedToday:     500,
			AdsPercentageToday:  5,
		}
	}
	router := http.NewServeMux()
	router.HandleFunc(cfg.APIPath, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if cfg.IsBasicAuth() {
			user, pass, ok := r.BasicAuth()
			if !ok {
				t.Error("basic auth not ok")
			}
			if user != cfg.BasicAuth.Username || pass != cfg.BasicAuth.Password {
				t.Error("basic auth username or password mismatch")
			}
		}
		resp, _ := json.Marshal(response)
		w.Write(resp)
	}))
	server := httptest.NewServer(router)
	cfg.ListenAddress = server.URL
	return cfg, server
}
