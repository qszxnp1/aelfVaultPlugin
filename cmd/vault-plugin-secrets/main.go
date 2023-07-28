package main

import (
	"github.com/hashicorp/go-hclog"
	"os"

	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/sdk/plugin"

	"github.com/qszxnp1/aelfVaultPlugin/internal/backend"
)

func main() {
	apiClientMeta := &api.PluginAPIClientMeta{}
	flags := apiClientMeta.FlagSet()
	logger := hclog.New(&hclog.LoggerOptions{})
	err := flags.Parse(os.Args[1:]) // Ignore command, strictly parse flags
	if err != nil {
		logger.Error("parse args err,", err)
		os.Exit(1)
	}

	tlsConfig := apiClientMeta.GetTLSConfig()
	tlsProviderFunc := api.VaultPluginTLSProvider(tlsConfig)

	err = plugin.Serve(&plugin.ServeOpts{
		BackendFactoryFunc: backend.Factory,
		TLSProviderFunc:    tlsProviderFunc,
	})
	if err != nil {
		logger.Error("plugin shutting down", "error", err)
		os.Exit(1)
	}
}
