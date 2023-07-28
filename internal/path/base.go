package path

import (
	"context"
	"fmt"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/qszxnp1/aelfVaultPlugin/internal/model"
)

type basePathConfig struct {
	config
}

func (b basePathConfig) getExistenceFunc() framework.ExistenceFunc {
	return func(ctx context.Context, req *logical.Request, data *framework.FieldData) (bool, error) {
		entry, err := req.Storage.Get(ctx, req.Path)
		if err != nil {
			return false, fmt.Errorf("existence check failed, %v", err)
		}
		return entry != nil, nil
	}
}

func (b basePathConfig) getAppAccountPath(appName string, appId string) string {
	return fmt.Sprintf("app/%s/%s/account", appName, appId)
}

func (b basePathConfig) getAppAccount(ctx context.Context, req *logical.Request, appName string, appId string) (*model.AppAccount, error) {
	path := b.getAppAccountPath(appName, appId)
	entry, err := req.Storage.Get(ctx, path)
	if err != nil || entry == nil {
		return nil, fmt.Errorf("existence check failed, %v", err)
	}

	var account *model.AppAccount
	err = entry.DecodeJSON(&account)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize account at %s", path)
	}
	if account == nil {
		return nil, fmt.Errorf("account not existed at %s", path)
	}
	return account, nil
}
