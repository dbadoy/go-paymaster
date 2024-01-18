package userop

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

var (
	entryPoint = common.HexToAddress("0x16a91A2E386d6DCa02EbB7611055CCa017e086E4")
	chainId    = big.NewInt(80001)
)

func TestUserOperation(t *testing.T) {
	op := UserOperation{
		Sender:               common.HexToAddress("0x2e8b6580cb32c0EB273093b39B756d0A19832295"),
		Nonce:                big.NewInt(0),
		InitCode:             common.Hex2Bytes("e403cBBA30e310B9d5bA5E5430EDeCf0e7AB3E00088924ef0000000000000000000000008ac1FdB47F1D4e63D6B19E92B631B10E2F7140070000000000000000000000000000000000000000000000000000000000000000"),
		CallData:             common.Hex2Bytes(""),
		CallGasLimit:         big.NewInt(300000),
		VerificationGasLimit: big.NewInt(1500000),
		PreVerificationGas:   big.NewInt(185000),
		MaxFeePerGas:         big.NewInt(30000000000),
		MaxPriorityFeePerGas: big.NewInt(2000000000),
		PaymasterAndData:     common.Hex2Bytes("2e5803d0E06A46970c83c37b7406E94bD0395c4700000000000000000000000039dae9e9ca8b284eb467f7694733ba173ba37579000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000413c5476c1be6b85d9a5343bdfe4a9325b81cfc35658a8915052aaa1e21f2e2b2b596e4661539fce3d0f5ac7c3aa7fea2d72b67728d3458e20235b4603d4f1fab71b00000000000000000000000000000000000000000000000000000000000000"),
		Signature:            common.Hex2Bytes("322604fb3ee4e2aeb76ed9ea7bb19216d8442b794ffcc4782441531bfb560cf30233c7afd5e2f4f964107b775c77a3b2dc2d7579f7d5e494b39d9779f016feeb1c"),
	}

	expectedPaymaster := common.HexToAddress("0x2e5803d0E06A46970c83c37b7406E94bD0395c47")
	if op.Paymaster() != expectedPaymaster {
		t.Errorf("expected paymaster %s, got %s", expectedPaymaster, op.Paymaster())
	}

	expectedFactory := common.HexToAddress("0xe403cBBA30e310B9d5bA5E5430EDeCf0e7AB3E00")
	if op.Factory() != expectedFactory {
		t.Errorf("expected factory %s, got %s", expectedFactory, op.Factory())
	}
}

func TestUserOperationHash(t *testing.T) {
	op := UserOperation{
		Sender:               common.HexToAddress("0x2e8b6580cb32c0EB273093b39B756d0A19832295"),
		Nonce:                big.NewInt(0),
		InitCode:             common.Hex2Bytes("e403cBBA30e310B9d5bA5E5430EDeCf0e7AB3E00088924ef0000000000000000000000008ac1FdB47F1D4e63D6B19E92B631B10E2F7140070000000000000000000000000000000000000000000000000000000000000000"),
		CallData:             common.Hex2Bytes(""),
		CallGasLimit:         big.NewInt(300000),
		VerificationGasLimit: big.NewInt(1500000),
		PreVerificationGas:   big.NewInt(185000),
		MaxFeePerGas:         big.NewInt(30000000000),
		MaxPriorityFeePerGas: big.NewInt(2000000000),
		PaymasterAndData:     common.Hex2Bytes(""),
		Signature:            common.Hex2Bytes(""),
	}

	expectedPacked := common.HexToHash("0xe67d0ba4edf61b1f3bb4725abc275be6785931f7a384c7246b5eeb557214d917")
	opHash := op.UserOpHash(entryPoint, chainId)
	if !bytes.Equal(opHash.Bytes(), expectedPacked.Bytes()) {
		t.Errorf("expected packed hash %x, got %x", expectedPacked, opHash)
	}
}
