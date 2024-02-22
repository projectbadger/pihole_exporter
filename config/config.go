package config

import (
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	// CLI banner
	banner = "pihole_exporter - %s\n"
)

var (
	// Application version
	Version = "dev"
	// Application name
	PackageName = "pihole_exporter"
	// Errors
	errFileIsDir = errors.New("file is directory")
)

// Config holds all of the configuration options.
type Config struct {
	// Metrics server configuration
	Web *Web `yaml:"web,omitempty"`
	// PiHole server configuration
	Pihole *Pihole `yaml:"pihole,omitempty"`
	// Logging configuration
	Log *Log `yaml:"log,omitempty"`
	// Verbose logging
	Debug bool `yaml:"-"`
	// Print version and exit
	Version bool `yaml:"-"`
	// Use a configuration file
	Config string `yaml:"-"`
}

// NewConfig returns a new *Config.
// It parses application flags and applies them to the struct.
// If a config file is used, it unmarshals the config file and
// applies CLI flags on top.
//
// CLI flags override config file values.
func NewConfig() (*Config, error) {
	cfg := &Config{
		Web:    &Web{TLS: &TLS{}, BasicAuth: &BasicAuth{}},
		Pihole: &Pihole{TLS: &TLS{}, BasicAuth: &BasicAuth{}},
		Log:    &Log{},
	}
	cfg.parseFlags()

	if cfg.Config == "" {
		return cfg, setupSLog(cfg)
	}

	f, err := readConfigFile(cfg.Config)
	if err != nil {
		slog.Error("error reading config file", "config", err.Error())
		return nil, err
	}
	cfgFromFile, err := yamlToConfig(f)
	if err != nil {
		slog.Error("error unmarshaling config file", "config", err.Error())
		return nil, err
	}
	// Web
	if cfgFromFile.Web == nil {
		cfgFromFile.Web = &Web{}
	}
	if cfgFromFile.Web.BasicAuth == nil {
		cfgFromFile.Web.BasicAuth = &BasicAuth{}
	}
	if cfg.Web.ListenAddress != "" {
		cfgFromFile.Web.ListenAddress = cfg.Web.ListenAddress
	}
	if cfg.Web.BasicAuth.IsBasicAuth() {
		cfgFromFile.Web.BasicAuth = cfg.Web.BasicAuth
	}
	if cfg.Web.MetricsPath != "" {
		cfgFromFile.Web.MetricsPath = cfg.Web.MetricsPath
	}
	if cfgFromFile.Web.TLS == nil {
		cfgFromFile.Web.TLS = &TLS{}
	}
	if cfg.Web.TLS.CACertificate != "" {
		cfgFromFile.Web.TLS.CACertificate = cfg.Web.TLS.CACertificate
	}
	if cfg.Web.TLS.Certificate != "" {
		cfgFromFile.Web.TLS.Certificate = cfg.Web.TLS.Certificate
	}
	if cfg.Web.TLS.Key != "" {
		cfgFromFile.Web.TLS.Key = cfg.Web.TLS.Key
	}
	// Pihole
	if cfgFromFile.Pihole == nil {
		cfgFromFile.Pihole = &Pihole{}
	}
	if cfg.Pihole.Listen != "" {
		cfgFromFile.Pihole.Listen = cfg.Pihole.Listen
	}
	if cfg.Pihole.BasicAuth.IsBasicAuth() {
		cfgFromFile.Pihole.BasicAuth = cfg.Pihole.BasicAuth
	}
	if cfgFromFile.Pihole.TLS == nil {
		cfgFromFile.Pihole.TLS = &TLS{}
	}
	if cfg.Pihole.TLS.CACertificate != "" {
		cfgFromFile.Pihole.TLS.CACertificate = cfg.Pihole.TLS.CACertificate
	}
	if cfg.Pihole.TLS.Certificate != "" {
		cfgFromFile.Pihole.TLS.Certificate = cfg.Pihole.TLS.Certificate
	}
	if cfg.Pihole.TLS.Key != "" {
		cfgFromFile.Pihole.TLS.Key = cfg.Pihole.TLS.Key
	}
	if cfgFromFile.Pihole.NumResults < 1 ||
		(cfg.Pihole.NumResults > 0 && cfg.Pihole.NumResults != 30) {
		cfgFromFile.Pihole.TLS.Key = cfg.Pihole.TLS.Key
	}
	// Log
	if cfgFromFile.Log == nil {
		cfgFromFile.Log = &Log{}
	}
	if cfg.Log.Bare {
		cfgFromFile.Log.Bare = cfg.Log.Bare
	}
	if cfg.Log.Output != "" {
		cfgFromFile.Log.Output = cfg.Log.Output
	}
	if cfg.Log.Level != "" {
		cfgFromFile.Log.Level = cfg.Log.Level
	}

	if err := setupSLog(cfg); err != nil {
		slog.Error("Error setting up logging", "error", err)
		return nil, err
	}

	return cfgFromFile, nil
}

// parseFlags sets the flags using flag package, parses them and
// checks for version and help flags.
func (cfg *Config) parseFlags() {
	if cfg == nil {
		panic("config is nil")
	}
	// General
	flag.BoolVar(&cfg.Version, "version", false, "print version and exit")
	flag.BoolVar(&cfg.Debug, "debug", false, "print verbose debugging information")
	flag.StringVar(&cfg.Config, "config", "", "Path to a config file")
	// Web
	flag.StringVar(&cfg.Web.ListenAddress, "web.listen-address", ":9311", "Metrics server listener address")
	flag.StringVar(&cfg.Web.MetricsPath, "web.metrics-path", "/metrics", "URL path on which metrics are exposed")
	flag.StringVar(&cfg.Web.TLS.Certificate, "web.tls.certificate", "", "Metrics server TLS certificate")
	flag.StringVar(&cfg.Web.TLS.Key, "web.tls.key", "", "Metrics server TLS private key")
	flag.StringVar(&cfg.Web.BasicAuth.Username, "web.basic-auth.username", "", "Web server basic auth username")
	flag.StringVar(&cfg.Web.BasicAuth.Password, "web.basic-auth.password", "", "Web server basic auth password")
	// Pihole
	flag.StringVar(&cfg.Pihole.Listen, "pihole.listen-address", "", "Pihole endpoint URL")
	flag.StringVar(&cfg.Pihole.TLS.CACertificate, "pihole.tls.ca-certificate", "", "CA certificate to trust when connecting to Pihole")
	// flag.StringVar(&cfg.Pihole.TLS.Certificate, "pihole.tls.certificate", "", "Pihole mTLS certificate")
	// flag.StringVar(&cfg.Pihole.TLS.Key, "pihole.tls.key", "", "Pihole mTLS private key")
	flag.StringVar(&cfg.Pihole.BasicAuth.Username, "pihole.basic-auth.username", "", "Pihole basic auth username")
	flag.StringVar(&cfg.Pihole.BasicAuth.Password, "pihole.basic-auth.password", "", "Pihole basic auth password")
	flag.StringVar(&cfg.Pihole.APIToken, "pihole.api-token", "", "Pihole API token for authentication")
	flag.Int64Var(&cfg.Pihole.NumResults, "pihole.num-results", 30, "Number of results returned for each query\n(top domains, top queries...)")
	// Log
	flag.StringVar(&cfg.Log.Format, "log.format", "", "Logging format: [ text | json ]\ndefault: text")
	flag.StringVar(&cfg.Log.Output, "log.output", "", "Logging output: [ stdout | /path/to/file.log ]\ndefault: stdout")
	flag.StringVar(&cfg.Log.Level, "log.level", "info", "Logging level: [ info | warn | error | debug ]\ndefault: info")
	flag.BoolVar(&cfg.Log.Bare, "log.bare", false, "Hide log level and timestamps when logging")
	flag.Usage = func() {
		fmt.Printf(banner, Version)
		flag.PrintDefaults()
	}

	flag.Parse()
	if cfg.Version {
		fmt.Printf("%s", Version)
		os.Exit(0)
	}

	if cfg.Web.ListenAddress == "" {
		cfg.printHelpAndExit("Pihole listen address must be defined", 1)
	}
}

// PrintHelpAndExit prints a customizable message along with
// application help and exits with the provided code.
func (cfg *Config) printHelpAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Println(message)
		// slog.Info(message)
	}
	flag.Usage()
	os.Exit(exitCode)
}

// Web holds PiHole exporter metrics server configuration data.
type Web struct {
	// PiHole exporter metrics server listen address
	ListenAddress string `yaml:"listen_address,omitempty"`
	// PiHole exporter URL path for serving metrics
	MetricsPath string `yaml:"metrics_path,omitempty"`
	// TLS config for PiHole exporter metrics server connection
	TLS *TLS `yaml:"tls,omitempty"`
	// Basic auth config for PiHole exporter
	// metrics server connection
	BasicAuth *BasicAuth `yaml:"basic_auth,omitempty"`
}

// IsTLS return whether TLS is enabled.
func (cfg *Web) IsTLS() bool {
	if cfg == nil {
		return false
	}
	return (cfg.TLS != nil &&
		cfg.TLS.Certificate != "" &&
		cfg.TLS.Key != "") &&
		!strings.HasPrefix(cfg.ListenAddress, "unix://") &&
		!strings.HasPrefix(cfg.ListenAddress, "http://")
}

// IsBasicAuth return whether basic auth is enabled.
func (cfg *Web) IsBasicAuth() bool {
	if cfg == nil {
		return false
	}
	return cfg.BasicAuth.IsBasicAuth()
}

// Pihole holds PiHole server configuration data.
type Pihole struct {
	// PiHole server address
	Listen string `yaml:"listen,omitempty"`
	// TLS config for PiHole server connection
	TLS *TLS `yaml:"tls,omitempty"`
	// Basic auth config for PiHole server connection
	BasicAuth *BasicAuth `yaml:"basic_auth,omitempty"`
	// PiHole admin API key
	APIToken string `yaml:"api_token,omtiempty"`
	// PiHole admin API URL path
	APIPath string `yaml:"api_path,omtiempty"`
	// Number of returned results for each query.
	// Default 30.
	NumResults int64 `yaml:"num_results,omitempty"`
}

// IsTLS return whether TLS is enabled.
func (cfg *Pihole) IsTLS() bool {
	if cfg == nil {
		return false
	}
	return (cfg.TLS != nil &&
		(cfg.TLS.CACertificate != "" ||
			cfg.TLS.Insecure)) &&
		(!strings.HasPrefix(cfg.Listen, "unix://") &&
			!strings.HasPrefix(cfg.Listen, "http://"))
}

const (
	// defaultAPIPath is the default PiHole API endpoint URL path.
	defaultAPIPath = "/admin/api.php"
)

// GetAPIPath returns PiHole API path. Defaults to `/admin/api.php`.
func (cfg *Pihole) GetAPIPath() string {
	if cfg == nil || cfg.APIPath == "" {
		return defaultAPIPath
	}
	return cfg.APIPath
}

// IsBasicAuth return whether basic auth is enabled.
func (cfg *Pihole) IsBasicAuth() bool {
	if cfg == nil {
		return false
	}
	return cfg.BasicAuth.IsBasicAuth()
}

// TLS holds TLS configuration.
type TLS struct {
	// Path to a root CA certificate
	CACertificate string `yaml:"ca_certificate.omitempty"`
	// Path to a server or client certificate
	Certificate string `yaml:"certificate.omitempty"`
	// Path to a certificate private key
	Key string `yaml:"key.omitempty"`
	// Do not verify TLS certificate validity
	Insecure bool `yaml:"insecure.omitempty"`
}

// BasicAuth holds basic auth configuration.
// Basic auth applies if both fields are not empty.
type BasicAuth struct {
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
}

// Log holds logging configuration.
type Log struct {
	Format string `yaml:"format,omitempty"`
	Bare   bool   `yaml:"bare,omitempty"`
	Level  string `yaml:"level,omitempty"`
	Output string `yaml:"output,omitempty"`
}

// SLogLevel returns logging level in slog format.
// Default is level info.
func (l *Log) SLogLevel() slog.Level {
	if l == nil {
		return slog.LevelInfo
	}
	switch l.Level {
	case "", "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	case "debug":
		return slog.LevelDebug
	case "none":
		return 0
	}
	return slog.LevelInfo
}

// IsBasicAuth returns true if both username and password is set,
func (a *BasicAuth) IsBasicAuth() bool {
	return a != nil &&
		a.Username != "" &&
		a.Password != ""
}

// yamlToConfig unmarshals a JSON byte array into a new Config.
func yamlToConfig(configBytes []byte) (c *Config, err error) {
	c = &Config{}
	err = yaml.Unmarshal(configBytes, c)
	return
}

// readConfigFile reads a file from the provided file path.
func readConfigFile(path string) (content []byte, err error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, errFileIsDir
	}
	content, err = os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %s", err.Error())
	}
	return
}
