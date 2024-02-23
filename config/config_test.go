package config

import (
	"log/slog"
	"pihole_exporter/test_data"
	"testing"
)

func TestBasicAuth(t *testing.T) {
	for _, testData := range test_data.IsBasicAuthStr {
		t.Run(testData.Name, func(t *testing.T) {
			basicAuthObj := &BasicAuth{
				Username: testData.InputStr[0],
				Password: testData.InputStr[1],
			}
			isBasicAuth := basicAuthObj.IsBasicAuth()
			if isBasicAuth && !testData.ExpectedBool[0] {
				t.Error("returned IsBasicAuth is true when should be false")
			} else if !isBasicAuth && testData.ExpectedBool[0] {
				t.Error("returned IsBasicAuth is false when should be true")
			}
		})
	}
}

func TestTLS(t *testing.T) {
	t.Run("IsTLS", func(t *testing.T) {
		for _, testData := range test_data.TLSIsSet {
			t.Run(testData.Name, func(t *testing.T) {
				tls := &TLS{
					CACertificate: testData.InputStr[0],
					Certificate:   testData.InputStr[1],
					Key:           testData.InputStr[2],
					Insecure:      testData.InputBool[0],
				}
				isTLS := tls.IsTLS()
				if isTLS && !testData.ExpectedBool[0] {
					t.Error("returned IsTLS is true when should be false")
				} else if !isTLS && testData.ExpectedBool[0] {
					t.Error("returned IsTLS is false when should be true")
				}
			})
		}
	})
	t.Run("Web.IsTLS", func(t *testing.T) {
		for _, testData := range test_data.WebTLSIsSet {
			t.Run(testData.Name, func(t *testing.T) {
				tls := &TLS{
					CACertificate: testData.InputStr[0],
					Certificate:   testData.InputStr[1],
					Key:           testData.InputStr[2],
					Insecure:      testData.InputBool[0],
				}
				web := &Web{TLS: tls, ListenAddress: testData.InputStr[3]}
				isSet := web.IsTLS()
				if isSet && !testData.ExpectedBool[0] {
					t.Error("returned IsSet is true when should be false")
				} else if !isSet && testData.ExpectedBool[0] {
					t.Error("returned IsSet is false when should be true")
				}
			})
		}
	})
	t.Run("Pihole.IsTLS", func(t *testing.T) {
		for _, testData := range test_data.PiholeTLSIsSet {
			t.Run(testData.Name, func(t *testing.T) {
				tls := &TLS{
					CACertificate: testData.InputStr[0],
					Certificate:   testData.InputStr[1],
					Key:           testData.InputStr[2],
					Insecure:      testData.InputBool[0],
				}
				pihole := &Pihole{TLS: tls, ListenAddress: testData.InputStr[3]}
				isSet := pihole.IsTLS()
				if isSet && !testData.ExpectedBool[0] {
					t.Error("returned IsSet is true when should be false")
				} else if !isSet && testData.ExpectedBool[0] {
					t.Error("returned IsSet is false when should be true")
				}
			})
		}
	})
}

func TestLog(t *testing.T) {
	t.Run("Level", func(t *testing.T) {
		t.Run("EmptyLogConfig", func(t *testing.T) {
			var l *Log
			level := l.SLogLevel()
			if level != slog.LevelInfo {
				t.Errorf("nil *Log doesn't return level info")
			}
		})
		for _, testData := range test_data.LogLevel {
			t.Run(testData.Name, func(t *testing.T) {
				l := &Log{
					Level: testData.InputStr[0],
				}
				slogLevel := int(l.SLogLevel())
				if slogLevel != testData.ExpectedInt[0] {
					t.Errorf("slog level doesn't match: expected: '%d', got: '%d'", testData.ExpectedInt[0], slogLevel)
				}
			})
		}
	})
}

func TestPihole(t *testing.T) {
	t.Run("APIPath", func(t *testing.T) {
		for _, testData := range test_data.PiholeAPIPath {
			t.Run(testData.Name, func(t *testing.T) {
				p := &Pihole{
					APIPath: testData.InputStr[0],
				}
				apiPath := p.GetAPIPath()
				if apiPath != testData.ExpectedStr[0] {
					t.Errorf("api path doesn't match: expected: '%s', got: '%s'", testData.ExpectedStr[0], apiPath)
				}
			})
		}
	})
	t.Run("BasicAuth", func(t *testing.T) {
		for _, testData := range test_data.IsBasicAuthStr {
			t.Run(testData.Name, func(t *testing.T) {
				p := &Pihole{
					BasicAuth: &BasicAuth{
						Username: testData.InputStr[0],
						Password: testData.InputStr[1],
					}}
				isBasicAuth := p.IsBasicAuth()
				if isBasicAuth && !testData.ExpectedBool[0] {
					t.Error("returned IsBasicAuth is true when should be false")
				} else if !isBasicAuth && testData.ExpectedBool[0] {
					t.Error("returned IsBasicAuth is false when should be true")
				}
			})
		}
	})
	// t.Run("NumResults", func(t *testing.T) {
	// 	for _, testData := range test_data.PiholeNumResults {
	// 		t.Run(testData.Name, func(t *testing.T) {
	// 			p := &Pihole{
	// 				NumResults: int64(testData.InputInt[0]),
	// 			}
	// 			numResults := p.NumResults
	// 			if apiPath != int64(testData.ExpectedInt[0]) {
	// 				t.Errorf("api path doesn't match: expected: '%s', got: '%s'", testData.ExpectedStr[0], apiPath)
	// 			}
	// 		})
	// 	}
	// })
}
