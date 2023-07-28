package path

import (
	"context"
	"encoding/hex"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/qszxnp1/aelfVaultPlugin/internal/model"
)

type signPathConfig struct {
	basePathConfig
}

func (a *signPathConfig) getPattern() string {
	appName := framework.GenericNameRegex("appName")
	appId := framework.GenericNameRegex("appId")
	return "app/" + appName + "/" + appId + "/sign"
}

func (a *signPathConfig) getHelpSynopsis() string {
	return "sign for app"
}

func (a *signPathConfig) getFields() map[string]*framework.FieldSchema {
	return map[string]*framework.FieldSchema{
		"appName": {
			Type:        framework.TypeString,
			Description: "The name of app",
		},
		"appId": {
			Type:        framework.TypeString,
			Description: "The id of app",
		},
		"data": {
			Type:        framework.TypeString,
			Description: "The data of app",
		},
	}
}

/*func (a *signPathConfig) getCallbacks() map[logical.Operation]framework.OperationFunc {
	return map[logical.Operation]framework.OperationFunc{
		logical.CreateOperation: a.sign,
		logical.UpdateOperation: a.sign,
	}
}*/

func (a *signPathConfig) getOperations() map[logical.Operation]framework.OperationHandler {
	return map[logical.Operation]framework.OperationHandler{
		logical.CreateOperation: &framework.PathOperation{
			Callback: a.sign,
		},
		logical.UpdateOperation: &framework.PathOperation{
			Callback: a.sign,
		},
	}
}

func (a *signPathConfig) sign(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	dataWrapper := model.NewFieldDataWrapper(data)
	appName, err := dataWrapper.MustGetString("appName")
	if err != nil {
		return nil, err
	}
	appId, err := dataWrapper.MustGetString("appId")
	if err != nil {
		return nil, err
	}

	extData, err := dataWrapper.MustGetString("data")
	if err != nil {
		return nil, err
	}

	account, err := a.getAppAccount(ctx, req, appName, appId)
	if err != nil {
		return nil, err
	}

	rawTransactionBytes, err := hex.DecodeString(extData)
	if err != nil {
		return nil, err
	}

	signData, err := account.Sign(rawTransactionBytes)
	if err != nil {
		return nil, err
	}

	return &logical.Response{
		Data: map[string]interface{}{
			"signData": signData,
		},
	}, nil
}
