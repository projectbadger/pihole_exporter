
# server

```go
import pihole_exporter/server
```

## Server

Server configures and runs a *http.Server with the provided config.Config configuration.
It servers metrics on the provided metrics path.

## Index

- [Run() error](#func-run-error)


## func [Run() error](<server.go#L14>)

Run configures a http.Server and runs it with the
provided configuration.


```go
func Run(cfg *config.Config) error
```

