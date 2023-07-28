package path

import (
	"encoding/hex"
	"github.com/qszxnp1/aelfVaultPlugin/internal/model"
	"testing"
)

func TestSign(t *testing.T) {
	// 创建一个 AppAccount 实例，用于测试
	appAccount := &model.AppAccount{
		PrivateKey: "b30126ec2d503fa0b4e78731584b3cca8b8d06097540722131355120f0b6a64d",
	}

	// 假设待签名的数据
	dataStr := "73720c9d3a3474d16a4f75149171ca3188fbb674af3abf8dd717156de6fef6df"
	rawTransactionBytes, err := hex.DecodeString(dataStr)
	if err != nil {
		t.Fatalf("Sign() returned an error: %v", err)
	}

	// 调用方法进行签名
	signature, err := appAccount.Sign(rawTransactionBytes)

	// 检查是否有错误发生
	if err != nil {
		t.Fatalf("Sign() returned an error: %v", err)
	}

	// 进行其他的断言测试
	// 例如，可以验证签名结果的长度是否正确等等

	// 期望的签名结果（这是一个假设的例子，请根据你的实际数据提供正确的签名结果）
	expectedSignature := "3a68c423965b48cad5af6aa805846670b6ebeea1d86f7b184267409709e5595818ba65d8da65ef5f9f3facf5c5327af49d4cedd03fbe32c878842dee97652c7201"

	// 检查实际签名结果是否与期望的一致
	if signature != expectedSignature {
		t.Errorf("Sign() returned incorrect signature. Expected: %s, Got: %s", expectedSignature, signature)
	}
}
