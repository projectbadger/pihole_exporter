package server

import (
	"crypto/tls"
	"log/slog"
	"net/http"
	"pihole_exporter/config"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Run configures a http.Server and runs it with the
// provided configuration.
func Run(cfg *config.Config) error {
	router := http.NewServeMux()
	slog.Debug("Metrics path", "path", cfg.Web.MetricsPath)
	router.HandleFunc("/", rootHandlerFunc(cfg))
	router.Handle(cfg.Web.MetricsPath, basicAuthMiddleware(promhttp.Handler(), cfg))

	handler := slog.Default().Handler()
	serverLogger := slog.NewLogLogger(handler, cfg.Log.SLogLevel())
	server := &http.Server{
		Addr:     cfg.Web.ListenAddress,
		Handler:  router,
		ErrorLog: serverLogger,
	}

	slog.Info("Starting listener", "address", cfg.Web.ListenAddress)
	if cfg.Web.IsTLS() {
		if cfg.Web.TLS.Certificate != "" && cfg.Web.TLS.Key != "" {
			cert, err := tls.LoadX509KeyPair(cfg.Web.TLS.Certificate, cfg.Web.TLS.Key)
			if err != nil {
				return err
			}
			server.TLSConfig = &tls.Config{
				Certificates: []tls.Certificate{cert},
			}
			return server.ListenAndServeTLS("", "")
		}
	}
	return server.ListenAndServe()
}

// rootHandlerFunc handler the root `/` path on the server.
func rootHandlerFunc(cfg *config.Config) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<html>
		 <head><title>PiHole Exporter</title></head>
		 <body>
		 <h1>Pihole Exporter</h1>
		 <p><a href="` + cfg.Web.MetricsPath + `">Metrics</a></p>
		 </body>
		 </html>`))
	})
}

// basicAuthMiddleware is a middleware that requires and
// authenticates basic auth on request, if basic auth is enabled,
func basicAuthMiddleware(handler http.Handler, cfg *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if cfg.Web.BasicAuth.IsBasicAuth() {
			user, pass, ok := req.BasicAuth()
			if !ok ||
				(user != cfg.Web.BasicAuth.Username ||
					pass != cfg.Web.BasicAuth.Password) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
		handler.ServeHTTP(w, req)
	})
}
