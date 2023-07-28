package path

import (
	"context"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/qszxnp1/aelfVaultPlugin/internal/model"
)

type appPathConfig struct {
	basePathConfig
}

func (a *appPathConfig) getPattern() string {
	appName := framework.GenericNameRegex("appName")
	appId := framework.GenericNameRegex("appId")
	return "app/" + appName + "/" + appId + "/account"
}

func (a *appPathConfig) getHelpSynopsis() string {
	return "create an account for app, include public key and private key"
}

func (a *appPathConfig) getFields() map[string]*framework.FieldSchema {
	return map[string]*framework.FieldSchema{
		"appName": {
			Type:        framework.TypeString,
			Description: "The name of app",
		},
		"appId": {
			Type:        framework.TypeString,
			Description: "The id of app",
		},
		"privateKey": {
			Type:        framework.TypeString,
			Description: "The private key of app",
		},
		"publicKey": {
			Type:        framework.TypeString,
			Description: "The public key of app",
		},
		"data": {
			Type:        framework.TypeString,
			Description: "The data of app",
		},
	}
}

func (a *appPathConfig) getCallbacks() map[logical.Operation]framework.OperationFunc {
	return map[logical.Operation]framework.OperationFunc{
		logical.CreateOperation: a.createAppAccount,
		logical.UpdateOperation: a.updateAppAccount,
		logical.ReadOperation:   a.readAppAccount,
	}
}

func (a *appPathConfig) getOperations() map[logical.Operation]framework.OperationHandler {
	return map[logical.Operation]framework.OperationHandler{
		logical.CreateOperation: &framework.PathOperation{
			Callback: a.createAppAccount,
		},
		logical.UpdateOperation: &framework.PathOperation{
			Callback: a.updateAppAccount,
		},
		logical.ReadOperation: &framework.PathOperation{
			Callback: a.readAppAccount,
		},
	}
}

func (a *appPathConfig) readAppAccount(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	dataWrapper := model.NewFieldDataWrapper(data)
	appName, err := dataWrapper.MustGetString("appName")
	if err != nil {
		return nil, err
	}
	appId, err := dataWrapper.MustGetString("appId")
	if err != nil {
		return nil, err
	}
	account, err := a.getAppAccount(ctx, req, appName, appId)
	if err != nil {
		return nil, err
	}
	return &logical.Response{
		Data: map[string]interface{}{
			"appName":    account.AppName,
			"appId":      account.AppId,
			"privateKey": account.PrivateKey,
			"publicKey":  account.PublicKey,
			"address":    account.Address,
			"data":       account.Data,
		},
	}, nil
}

func (a *appPathConfig) updateAppAccount(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	dataWrapper := model.NewFieldDataWrapper(data)
	appName, err := dataWrapper.MustGetString("appName")
	if err != nil {
		return nil, err
	}
	appId, err := dataWrapper.MustGetString("appId")
	if err != nil {
		return nil, err
	}
	account, err := a.getAppAccount(ctx, req, appName, appId)
	if err != nil {
		return nil, err
	}
	privateKey, err := dataWrapper.MustGetString("privateKey")
	if err != nil {
		account.UpdatePrivateKey(privateKey)
	}
	publicKey, err := dataWrapper.MustGetString("publicKey")
	if err != nil {
		account.UpdatePublicKey(publicKey)
	}
	address, _ := dataWrapper.MustGetString("address")
	if err != nil {
		account.UpdateAddress(address)
	}
	extData, _ := dataWrapper.MustGetString("data")
	if err != nil {
		account.UpdateData(extData)
	}

	return &logical.Response{
		Data: map[string]interface{}{
			"appName": account.AppName,
			"appId":   account.AppId,
		},
	}, nil
}

func (a *appPathConfig) createAppAccount(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	dataWrapper := model.NewFieldDataWrapper(data)
	appName, err := dataWrapper.MustGetString("appName")
	if err != nil {
		return nil, err
	}
	appId, err := dataWrapper.MustGetString("appId")
	if err != nil {
		return nil, err
	}
	privateKey, err := dataWrapper.MustGetString("privateKey")
	if err != nil {
		return nil, err
	}
	publicKey, _ := dataWrapper.MustGetString("publicKey")
	address, _ := dataWrapper.MustGetString("address")
	extData, _ := dataWrapper.MustGetString("data")
	account := model.NewAppAccount(address, appName, appId, privateKey, publicKey, extData)

	entry, err := logical.StorageEntryJSON(req.Path, account)
	if err != nil {
		return nil, err
	}
	err = req.Storage.Put(ctx, entry)
	if err != nil {
		return nil, err
	}
	return &logical.Response{
		Data: map[string]interface{}{
			"appName": account.AppName,
			"appId":   account.AppId,
		},
	}, nil
}
