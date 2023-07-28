package model

import (
	"crypto/sha256"
	"encoding/hex"

	secp256 "github.com/haltingstate/secp256k1-go"
)

type AppAccount struct {
	Address    string `json:"address"`
	AppName    string `json:"app_name"`
	AppId      string `json:"app_id"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	Data       string `json:"data"`
}

func NewAppAccount(address string, appName string, appId string, privateKey string, publicKey string, data string) *AppAccount {
	return &AppAccount{
		Address:    address,
		AppName:    appName,
		AppId:      appId,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Data:       data,
	}
}

func (a *AppAccount) UpdatePrivateKey(privateKey string) {
	a.PrivateKey = privateKey
}

func (a *AppAccount) UpdatePublicKey(publicKey string) {
	a.PublicKey = publicKey
}

func (a *AppAccount) UpdateData(data string) {
	a.Data = data
}

func (a *AppAccount) UpdateAddress(address string) {
	a.Address = address
}

func (a *AppAccount) Sign(data []byte) (string, error) {
	privateKeyBytes, _ := hex.DecodeString(a.PrivateKey)
	txDataBytes := sha256.Sum256(data)
	signatureBytes := secp256.Sign(txDataBytes[:], privateKeyBytes)
	return hex.EncodeToString(signatureBytes), nil
}
