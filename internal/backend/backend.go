package backend

import (
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/qszxnp1/aelfVaultPlugin/internal/path"
)

type aelfBackend struct {
	*framework.Backend
}

// returns ethereumBackend
func newBackend(conf *logical.BackendConfig) (*aelfBackend, error) {
	var b aelfBackend
	b.Backend = &framework.Backend{
		Help: "",
		Paths: framework.PathAppend(
			path.GetPaths(),
		),
		PathsSpecial: &logical.Paths{
			SealWrapStorage: []string{
				"app/",
			},
		},
		Secrets:     []*framework.Secret{},
		BackendType: logical.TypeLogical,
	}
	return &b, nil
}
