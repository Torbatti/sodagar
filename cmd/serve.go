package cmd

import (
	"github.com/openjogd/sodagar/apis"
	"github.com/openjogd/sodagar/core"
	"github.com/spf13/cobra"
)

func NewServeCommand(app core.App) *cobra.Command {

	var allowedOrigins []string
	var httpAddr string
	var httpsAddr string

	command := &cobra.Command{
		Use:   "serve [domain(s)]",
		Args:  cobra.ArbitraryArgs,
		Short: "Starts the web server (default to 127.0.0.1:8090 if no domain is specified)",
		RunE: func(command *cobra.Command, args []string) error {
			// set default listener addresses if at least one domain is specified
			if len(args) > 0 {
				if httpAddr == "" {
					httpAddr = "0.0.0.0:80"
				}
				if httpsAddr == "" {
					httpsAddr = "0.0.0.0:443"
				}
			} else {
				if httpAddr == "" {
					httpAddr = "127.0.0.1:8090"
				}
			}

			for i := 0; i < len(args); i++ {
				println(args[i])
			}

			err := apis.Serve(app, apis.ServeConfig{
				HttpAddr:  httpAddr,
				HttpsAddr: httpsAddr,
				// ShowStartBanner: false,
				AllowedOrigins: allowedOrigins,
				// CertificateDomains: args,
			})

			return err
		},
	}

	return command
}
