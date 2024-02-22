
# config

```go
import pihole_exporter/config
```

## Config

Config package provides the Config type that holds all the configuration data for PiHole Exporter. It parses CLI flags and is able to define itself from a file, provided by a CLI flag.

## Index

- [type BasicAuth](#type-basicauth)
  - [IsBasicAuth() bool](#func-basicauth-isbasicauth-bool)
- [type Config](#type-config)
  - [NewConfig() (Config, error)](#func-newconfig-config-error)
- [type Log](#type-log)
  - [SLogLevel() slog](#func-log-sloglevel-slog)
- [type Pihole](#type-pihole)
  - [GetAPIPath() string](#func-pihole-getapipath-string)
  - [IsBasicAuth() bool](#func-pihole-isbasicauth-bool)
  - [IsTLS() bool](#func-pihole-istls-bool)
- [type TLS](#type-tls)
- [type Web](#type-web)
  - [IsBasicAuth() bool](#func-web-isbasicauth-bool)
  - [IsTLS() bool](#func-web-istls-bool)
- [Variables](#variables)

## Variables
```go
var (
	// Application version
	Version	= "dev"
	// Application name
	PackageName	= "pihole_exporter"
)

```


## type [BasicAuth](<config.go#L303>)

BasicAuth holds basic auth configuration.
Basic auth applies if both fields are not empty.
```go
type BasicAuth struct {
	Username	string	`yaml:"username,omitempty"`
	Password	string	`yaml:"password,omitempty"`
}
```

## func (*BasicAuth) [IsBasicAuth() bool](<config.go#L338>)

IsBasicAuth returns true if both username and password is set,


```go
func (a *BasicAuth) IsBasicAuth() bool
```

## type [Config](<config.go#L29>)

Config holds all of the configuration options.
```go
type Config struct {
	// Metrics server configuration
	Web	*Web	`yaml:"web,omitempty"`
	// PiHole server configuration
	Pihole	*Pihole	`yaml:"pihole,omitempty"`
	// Logging configuration
	Log	*Log	`yaml:"log,omitempty"`
	// Verbose logging
	Debug	bool	`yaml:"-"`
	// Print version and exit
	Version	bool	`yaml:"-"`
	// Use a configuration file
	Config	string	`yaml:"-"`
}
```

## func [NewConfig() (Config, error)](<config.go#L50>)

NewConfig returns a new *Config.
It parses application flags and applies them to the struct.
If a config file is used, it unmarshals the config file and
applies CLI flags on top.

CLI flags override config file values.


```go
func NewConfig() (*Config, error)
```

## type [Log](<config.go#L309>)

Log holds logging configuration.
```go
type Log struct {
	Format	string	`yaml:"format,omitempty"`
	Bare	bool	`yaml:"bare,omitempty"`
	Level	string	`yaml:"level,omitempty"`
	Output	string	`yaml:"output,omitempty"`
}
```

## func (*Log) [SLogLevel() slog](<config.go#L318>)

SLogLevel returns logging level in slog format.
Default is level info.


```go
func (l *Log) SLogLevel() slog.Level
```

## type [Pihole](<config.go#L240>)

Pihole holds PiHole server configuration data.
```go
type Pihole struct {
	// PiHole server address
	Listen	string	`yaml:"listen,omitempty"`
	// TLS config for PiHole server connection
	TLS	*TLS	`yaml:"tls,omitempty"`
	// Basic auth config for PiHole server connection
	BasicAuth	*BasicAuth	`yaml:"basic_auth,omitempty"`
	// PiHole admin API key
	APIToken	string	`yaml:"api_token,omtiempty"`
	// PiHole admin API URL path
	APIPath	string	`yaml:"api_path,omtiempty"`
	// Number of returned results for each query.
	// Default 30.
	NumResults	int64	`yaml:"num_results,omitempty"`
}
```

## func (*Pihole) [GetAPIPath() string](<config.go#L274>)

GetAPIPath returns PiHole API path. Defaults to `/admin/api.php`.


```go
func (cfg *Pihole) GetAPIPath() string
```
## func (*Pihole) [IsBasicAuth() bool](<config.go#L282>)

IsBasicAuth return whether basic auth is enabled.


```go
func (cfg *Pihole) IsBasicAuth() bool
```
## func (*Pihole) [IsTLS() bool](<config.go#L257>)

IsTLS return whether TLS is enabled.


```go
func (cfg *Pihole) IsTLS() bool
```

## type [TLS](<config.go#L290>)

TLS holds TLS configuration.
```go
type TLS struct {
	// Path to a root CA certificate
	CACertificate	string	`yaml:"ca_certificate.omitempty"`
	// Path to a server or client certificate
	Certificate	string	`yaml:"certificate.omitempty"`
	// Path to a certificate private key
	Key	string	`yaml:"key.omitempty"`
	// Do not verify TLS certificate validity
	Insecure	bool	`yaml:"insecure.omitempty"`
}
```

## type [Web](<config.go#L207>)

Web holds PiHole exporter metrics server configuration data.
```go
type Web struct {
	// PiHole exporter metrics server listen address
	ListenAddress	string	`yaml:"listen_address,omitempty"`
	// PiHole exporter URL path for serving metrics
	MetricsPath	string	`yaml:"metrics_path,omitempty"`
	// TLS config for PiHole exporter metrics server connection
	TLS	*TLS	`yaml:"tls,omitempty"`
	// Basic auth config for PiHole exporter
	// metrics server connection
	BasicAuth	*BasicAuth	`yaml:"basic_auth,omitempty"`
}
```

## func (*Web) [IsBasicAuth() bool](<config.go#L232>)

IsBasicAuth return whether basic auth is enabled.


```go
func (cfg *Web) IsBasicAuth() bool
```
## func (*Web) [IsTLS() bool](<config.go#L220>)

IsTLS return whether TLS is enabled.


```go
func (cfg *Web) IsTLS() bool
```

