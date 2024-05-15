package apis

import (
	"fmt"
	"net/http"
	"time"

	"github.com/openjogd/sodagar/core"
)

type ServeConfig struct {
	// ShowStartBanner indicates whether to show or hide the server start console message.
	ShowStartBanner bool

	// HttpAddr is the TCP address to listen for the HTTP server (eg. `127.0.0.1:80`).
	HttpAddr string

	// HttpsAddr is the TCP address to listen for the HTTPS server (eg. `127.0.0.1:443`).
	HttpsAddr string

	// Optional domains list to use when issuing the TLS certificate.
	//
	// If not set, the host from the bound server address will be used.
	//
	// For convenience, for each "non-www" domain a "www" entry and
	// redirect will be automatically added.
	// CertificateDomains []string

	// AllowedOrigins is an optional list of CORS origins (default to "*").
	AllowedOrigins []string
}

func Serve(app core.App, config ServeConfig) error {
	if len(config.AllowedOrigins) == 0 {
		config.AllowedOrigins = []string{"*"}
	}

	router, err := InitApi(app)
	if err != nil {
		return err
	}

	mainAddr := config.HttpAddr
	if config.HttpsAddr != "" {
		mainAddr = config.HttpsAddr
	}

	server := &http.Server{

		ReadTimeout:       10 * time.Minute,
		ReadHeaderTimeout: 30 * time.Second,
		// WriteTimeout: 60 * time.Second, // breaks sse!
		Handler: router,
		Addr:    mainAddr,
	}

	// start HTTPS server
	if config.HttpsAddr != "" {
		println("server runnin on : ", config.HttpsAddr)
		return server.ListenAndServe()
	}

	// start HTTP server
	if config.HttpAddr != "" {
		println("server runnin on : ", config.HttpAddr)
		return server.ListenAndServe()
	}

	// TODO: Handle When http and https addresses are not set
	return fmt.Errorf("broken")
}
