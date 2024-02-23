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
	cfg.parseFlags(os.Args[0], os.Args[1:])
	cfg.validate()

	if cfg.Config == "" {
		return cfg, setupSLog(cfg)
	}

	cfgFromFile, err := readConfigFile(cfg.Config)
	if err != nil {
		slog.Error("error reading config file", "config", err.Error())
		return nil, err
	}

	cfgFromFile = overwriteConfig(cfgFromFile, cfg)

	if err := setupSLog(cfg); err != nil {
		slog.Error("Error setting up logging", "error", err)
		return nil, err
	}

	return cfgFromFile, nil
}

// parseFlags sets the flags using flag package, parses them and
// checks for version and help flags.
func (cfg *Config) parseFlags(name string, args []string) {
	if cfg == nil {
		panic("config is nil")
	}
	flagSet := flag.NewFlagSet(name, flag.ExitOnError)
	// General
	flagSet.BoolVar(&cfg.Version, "version", false, "print version and exit")
	flagSet.BoolVar(&cfg.Debug, "debug", false, "print verbose debugging information")
	flagSet.StringVar(&cfg.Config, "config", "", "Path to a config file")
	// Web
	flagSet.StringVar(&cfg.Web.ListenAddress, "web.listen-address", ":9311", "Metrics server listener address")
	flagSet.StringVar(&cfg.Web.MetricsPath, "web.metrics-path", "/metrics", "URL path on which metrics are exposed")
	flagSet.StringVar(&cfg.Web.TLS.Certificate, "web.tls.certificate", "", "Metrics server TLS certificate")
	flagSet.StringVar(&cfg.Web.TLS.Key, "web.tls.key", "", "Metrics server TLS private key")
	flagSet.StringVar(&cfg.Web.BasicAuth.Username, "web.basic-auth.username", "", "Web server basic auth username")
	flagSet.StringVar(&cfg.Web.BasicAuth.Password, "web.basic-auth.password", "", "Web server basic auth password")
	// Pihole
	flagSet.StringVar(&cfg.Pihole.ListenAddress, "pihole.listen-address", "", "Pihole endpoint URL")
	flagSet.StringVar(&cfg.Pihole.TLS.CACertificate, "pihole.tls.ca-certificate", "", "CA certificate to trust when connecting to Pihole")
	// flagSet.StringVar(&cfg.Pihole.TLS.Certificate, "pihole.tls.certificate", "", "Pihole mTLS certificate")
	// flagSet.StringVar(&cfg.Pihole.TLS.Key, "pihole.tls.key", "", "Pihole mTLS private key")
	flagSet.StringVar(&cfg.Pihole.BasicAuth.Username, "pihole.basic-auth.username", "", "Pihole basic auth username")
	flagSet.StringVar(&cfg.Pihole.BasicAuth.Password, "pihole.basic-auth.password", "", "Pihole basic auth password")
	flagSet.StringVar(&cfg.Pihole.APIToken, "pihole.api-token", "", "Pihole API token for authentication")
	flagSet.Int64Var(&cfg.Pihole.NumResults, "pihole.num-results", 30, "Number of results returned for each query\n(top domains, top queries...)")
	// Log
	flagSet.StringVar(&cfg.Log.Format, "log.format", "", "Logging format: [ text | json ]\ndefault: text")
	flagSet.StringVar(&cfg.Log.Output, "log.output", "", "Logging output: [ stdout | /path/to/file.log ]\ndefault: stdout")
	flagSet.StringVar(&cfg.Log.Level, "log.level", "info", "Logging level: [ info | warn | error | debug ]\ndefault: info")
	flagSet.BoolVar(&cfg.Log.Bare, "log.bare", false, "Hide log level and timestamps when logging")
	flagSet.Usage = func() {
		fmt.Printf(banner, Version)
		flagSet.PrintDefaults()
	}

	flagSet.Parse(args)
}

// validate checks config for version and help flags and exists if
// found. Also checks for presence of exporter listen address.
func (cfg *Config) validate() {
	if cfg == nil {
		return
	}
	if cfg.Version {
		fmt.Printf("%s", Version)
		os.Exit(0)
	}

	if cfg.Web.ListenAddress == "" {
		cfg.printHelpAndExit("Web listen address must be defined", 1)
	}

	if cfg.Pihole.ListenAddress == "" {
		cfg.printHelpAndExit("Pihole server address must be defined", 1)
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

// overwriteConfig overwrites baseConfig with non-default values
// from overridingConfig and returns the result.
func overwriteConfig(baseConfig, overridingConfig *Config) *Config {
	// Web
	if baseConfig.Web == nil {
		baseConfig.Web = &Web{}
	}
	if baseConfig.Web.BasicAuth == nil {
		baseConfig.Web.BasicAuth = &BasicAuth{}
	}
	if overridingConfig.Web.ListenAddress != "" {
		baseConfig.Web.ListenAddress = overridingConfig.Web.ListenAddress
	}
	if overridingConfig.Web.BasicAuth.IsBasicAuth() {
		baseConfig.Web.BasicAuth = overridingConfig.Web.BasicAuth
	}
	if overridingConfig.Web.MetricsPath != "" {
		baseConfig.Web.MetricsPath = overridingConfig.Web.MetricsPath
	}
	if baseConfig.Web.TLS == nil {
		baseConfig.Web.TLS = &TLS{}
	}
	if overridingConfig.Web.TLS.CACertificate != "" {
		baseConfig.Web.TLS.CACertificate = overridingConfig.Web.TLS.CACertificate
	}
	if overridingConfig.Web.TLS.Certificate != "" {
		baseConfig.Web.TLS.Certificate = overridingConfig.Web.TLS.Certificate
	}
	if overridingConfig.Web.TLS.Key != "" {
		baseConfig.Web.TLS.Key = overridingConfig.Web.TLS.Key
	}
	// Pihole
	if baseConfig.Pihole == nil {
		baseConfig.Pihole = &Pihole{}
	}
	if overridingConfig.Pihole.ListenAddress != "" {
		baseConfig.Pihole.ListenAddress = overridingConfig.Pihole.ListenAddress
	}
	if overridingConfig.Pihole.BasicAuth.IsBasicAuth() {
		baseConfig.Pihole.BasicAuth = overridingConfig.Pihole.BasicAuth
	}
	if baseConfig.Pihole.TLS == nil {
		baseConfig.Pihole.TLS = &TLS{}
	}
	if overridingConfig.Pihole.TLS.CACertificate != "" {
		baseConfig.Pihole.TLS.CACertificate = overridingConfig.Pihole.TLS.CACertificate
	}
	if overridingConfig.Pihole.TLS.Certificate != "" {
		baseConfig.Pihole.TLS.Certificate = overridingConfig.Pihole.TLS.Certificate
	}
	if overridingConfig.Pihole.TLS.Key != "" {
		baseConfig.Pihole.TLS.Key = overridingConfig.Pihole.TLS.Key
	}
	if baseConfig.Pihole.NumResults < 1 ||
		(overridingConfig.Pihole.NumResults > 0 && overridingConfig.Pihole.NumResults != 30) {
		baseConfig.Pihole.TLS.Key = overridingConfig.Pihole.TLS.Key
	}
	// Log
	if baseConfig.Log == nil {
		baseConfig.Log = &Log{}
	}
	if overridingConfig.Log.Bare {
		baseConfig.Log.Bare = overridingConfig.Log.Bare
	}
	if overridingConfig.Log.Output != "" {
		baseConfig.Log.Output = overridingConfig.Log.Output
	}
	if overridingConfig.Log.Level != "" {
		baseConfig.Log.Level = overridingConfig.Log.Level
	}
	return baseConfig
}

// yamlToConfig unmarshals a JSON byte array into a new Config.
func yamlToConfig(configBytes []byte) (c *Config, err error) {
	c = &Config{}
	err = yaml.Unmarshal(configBytes, c)
	return
}

// readFile reads a file from the provided file path.
func readFile(path string) (content []byte, err error) {
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

// readFile reads a file from the provided file path.
func readConfigFile(path string) (config *Config, err error) {
	f, err := readFile(path)
	if err != nil {
		return nil, err
	}
	config, err = yamlToConfig(f)
	return
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
	ListenAddress string `yaml:"listen_address,omitempty"`
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
		(!strings.HasPrefix(cfg.ListenAddress, "unix://") &&
			!strings.HasPrefix(cfg.ListenAddress, "http://"))
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
	CACertificate string `yaml:"ca_certificate,omitempty"`
	// Path to a server or client certificate
	Certificate string `yaml:"certificate,omitempty"`
	// Path to a certificate private key
	Key string `yaml:"key,omitempty"`
	// Do not verify TLS certificate validity
	Insecure bool `yaml:"insecure,omitempty"`
}

// IsSet returns if TLS config is set for TLS
func (cfg *TLS) IsTLS() bool {
	if cfg == nil {
		return false
	}
	if cfg.Insecure {
		return true
	}
	if cfg.CACertificate != "" {
		return true
	}
	if cfg.Certificate != "" && cfg.Key != "" {
		return true
	}
	return false
}

// BasicAuth holds basic auth configuration.
// Basic auth applies if both fields are not empty.
type BasicAuth struct {
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
}

// IsBasicAuth returns true if both username and password is set,
func (a *BasicAuth) IsBasicAuth() bool {
	return a != nil &&
		a.Username != "" &&
		a.Password != ""
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
