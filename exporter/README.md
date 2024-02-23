
# exporter

```go
import pihole_exporter/exporter
```

## Exporter

Exporter package provides the Exporter type that implements prometheus.Collector interface. It registers itself and collects metrics from a PiHole server with pihole.Client.

## Index

- [NewMetrics() metrics](#func-newmetrics-metrics)

- [type Exporter](#type-exporter)
  - [NewExporter() (Exporter, error)](#func-newexporter-exporter-error)
  - [Collect()](#func-exporter-collect)
  - [Describe()](#func-exporter-describe)
  - [Register()](#func-exporter-register)

## func [NewMetrics() metrics](<metrics.go#L78>)

NewMetrics returns a new *metrics struct holding all of
the prometheus.Desc metrics descriptions.


```go
func NewMetrics() *metrics
```


## type [Exporter](<exporter.go#L19>)

Exporter collects Pihole metrics from the given address and
exports them using the Prometheus metrics package.
```go
type Exporter struct {
	// PiHole client for calling the admin API
	Pihole *pihole.Client
	// contains filtered or unexported fields
}
```

## func [NewExporter() (Exporter, error)](<exporter.go#L26>)

NewExporter returns an initialized Exporter.


```go
func NewExporter(cfg *config.Config) (*Exporter, error)
```

## func (*Exporter) [Collect()](<exporter.go#L52>)

Collect collects the metrics from the channel and sends them
as Prometheus metrics.


```go
func (e *Exporter) Collect(ch chan<- prometheus.Metric)
```
## func (*Exporter) [Describe()](<exporter.go#L43>)

Describe publishes all of the collected PiHole metrics to the
provided channel by calling the underlying *metrics.Describe(ch).


```go
func (e *Exporter) Describe(ch chan<- *prometheus.Desc)
```
## func (*Exporter) [Register()](<exporter.go#L108>)

Register registers metrics in the prometheus package.


```go
func (e *Exporter) Register()
```

