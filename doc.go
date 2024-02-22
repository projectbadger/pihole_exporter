// PiHole Exporter is a Prometheus Exporter for PiHole.
/*
It checks your [PiHole](https://pi-hole.net/) statistics through the available API, supporting basic TLS settings and basic authentication.

### Available metrics

* Ads blocked - counter
* Ads percentage - gauge
* Domains blocked - counter
* DNS Queries - counter
* Unique domains - counter
* Queries forwarded - counter
* Queries cached - counter
* Clients ever seen - counter
* Unique clients - counter
* DNS queries all types - counter
* DNS queries all replies - counter
* Gravity last updated - unix timestamp
* Privacy level
* Top Ads - counter with label `domain``
* Top Queries - counter with label `domain`
* Top clients - counter with label `domain`
* Query types - counter with label `type`
* Forward destinations - counter with label `destination`
* Replies - counter with label `domain`

### Example of available metrics

```
# HELP pihole_ads_blocked Ads blocked today
# TYPE pihole_ads_blocked gauge
pihole_ads_blocked 2054
# HELP pihole_ads_percentage Ads percentage today
# TYPE pihole_ads_percentage gauge
pihole_ads_percentage 2.897692
# HELP pihole_clients_ever_seen Number of different clients ever
# TYPE pihole_clients_ever_seen gauge
pihole_clients_ever_seen 16
# HELP pihole_dns_queries_all_replies DNS queries all replies
# TYPE pihole_dns_queries_all_replies gauge
pihole_dns_queries_all_replies 70884
# HELP pihole_dns_queries_all_types DNS queries all types
# TYPE pihole_dns_queries_all_types gauge
pihole_dns_queries_all_types 70884
# HELP pihole_dns_queries_today DNS queries today
# TYPE pihole_dns_queries_today gauge
pihole_dns_queries_today 70884
# HELP pihole_domains_being_blocked Domains being blocked today
# TYPE pihole_domains_being_blocked gauge
pihole_domains_being_blocked 154587
# HELP pihole_privacy_level Privacy level
# TYPE pihole_privacy_level gauge
pihole_privacy_level 0
# HELP pihole_queries_cached Cached queries
# TYPE pihole_queries_cached gauge
pihole_queries_cached 55653
# HELP pihole_queries_forwarded Forwarded queries
# TYPE pihole_queries_forwarded gauge
pihole_queries_forwarded 12901
# HELP pihole_query_types DNS query types
# TYPE pihole_query_types gauge
pihole_query_types{type="A (IPv4)"} 59.98
pihole_query_types{type="AAAA (IPv6)"} 36.71
pihole_query_types{type="ANY"} 0
pihole_query_types{type="DNSKEY"} 0
...
```

### Available configuration options

  pihole_exporter -
  -config string
        Path to a config file
  -debug
        print verbose debugging information
  -log.bare
        Hide log level and timestamps when logging
  -log.format string
        Logging format: [ text | json ]
        default: text
  -log.level string
        Logging level: [ info | warn | error | debug ]
        default: info (default "info")
  -log.output string
        Logging output: [ stdout | /path/to/file.log ]
        default: stdout
  -pihole.api-token string
        Pihole API token for authentication
  -pihole.basic-auth.password string
        Pihole basic auth password
  -pihole.basic-auth.username string
        Pihole basic auth username
  -pihole.listen-address string
        Pihole endpoint URL
  -pihole.num-results int
        Number of results returned for each query
  -pihole.tls.ca-certificate string
        CA certificate to trust when connecting to Pihole
  -version
        print version and exit
  -web.basic-auth.password string
        Web server basic auth password
  -web.basic-auth.username string
        Web server basic auth username
  -web.listen-address string
        Metrics server listener address (default ":9311")
  -web.metrics-path string
        URL path on which metrics are exposed (default "/metrics")
  -web.tls.certificate string
        Metrics server TLS certificate
  -web.tls.key string
        Metrics server TLS private key

### Example config file

  ---
  web:
    listen_address: :9311
    metrics_path: /metrics
    basic_auth:
      username: "pihole_exporter_user"
      password: "pihole_exporter_pass"
    tls:
      certificate: ./cert.pem
      key: ./key.pem
  pihole:
    listen_address: https://pihole.localhost.localdomain
    api_token: "90a8una09n8ua0wpe8nuan09a8nefnaawe2a3bta34t"
    basic_auth:
      username: "pihole_user"
      password: "pihole_pass"
    tls:
      ca: ./ca.cert.pem
  log:
    format: text
    output: stdout
    level: error

## Installation

You can download the binaries from the [release page](https://github.com/projectbadger/pihole_exporter/-/releases) or build the exporter yourself.

## Usage

Run PiHole exporter:
```sh
# Show help
pihole_exporter -help
# Run with default settings and connect to a local http-only PiHole
pihole_exporter -pihole.listen-address http://pi.hole
# Run with debug log level and connect to a PiHole server
# with a self-signed certificate
pihole_exporter -log.level debug -pihole.listen-address https://pi.hole -pihole.tls.ca-certificate ./custom.ca.crt
```

## Development

* Build tool :
```sh
make build-all
# build for specific platform
make build-linux-amd64
```

* Launch unit tests :

```sh
make test
```

* Launch with compiler flags :

```sh
# run as `pihole_exporter -h`
make run RUN_ARGS="-"
```

## Local Deployment

* Run Prometheus using a configuration file with pihole_exporter as target
```sh
prometheus -config.file prometheus.yml
```

* Launch exporter:
```sh
pihole_exporter -config pihole_exporter.yml
```

* Check that Prometheus finds the PiHole exporter on `http://localhost:9090/targets`

## Docker Deployment

* Build Image:
```sh
docker build -t github.com/projectbadger/pihole_exporter .
```

* Start Container
```sh
docker run -d -p 9311:9311 github.com/projectbadger/pihole_exporter -pihole.listen-address https://pi.hole
# Start container with config and a trusted root CA
docker run -d -p 9311:9311 -v config.yml:/etc/pihole_exporter/config.yml:ro -v ca.crt:/etc/pihole_exporter/ca.crt:ro github.com/projectbadger/pihole_exporter -config /etc/pihole_exporter/config.yml -pihole.tls.ca-certificate /etc/pihole_exporter/ca.crt
```
*/
package main
