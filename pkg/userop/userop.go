package userop

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type UserOperation struct {
	Sender               common.Address `json:"sender"`
	Nonce                *big.Int       `json:"nonce"`
	InitCode             []byte         `json:"initCode"`
	CallData             []byte         `json:"callData"`
	CallGasLimit         *big.Int       `json:"callGasLimit"`
	VerificationGasLimit *big.Int       `json:"verificationGasLimit"`
	PreVerificationGas   *big.Int       `json:"preVerificationGas"`
	MaxFeePerGas         *big.Int       `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *big.Int       `json:"maxPriorityFeePerGas"`
	PaymasterAndData     []byte         `json:"paymasterAndData"`
	Signature            []byte         `json:"signature"`
}

func (op *UserOperation) Paymaster() common.Address {
	if len(op.PaymasterAndData) < common.AddressLength {
		return common.Address{}
	}

	return common.BytesToAddress(op.PaymasterAndData[:common.AddressLength])
}

func (op *UserOperation) Factory() common.Address {
	if len(op.PaymasterAndData) < common.AddressLength {
		return common.Address{}
	}

	return common.BytesToAddress(op.InitCode[:common.AddressLength])
}

func (op *UserOperation) Pack() []byte {
	args := abi.Arguments{
		{Name: "UserOp", Type: UserOpType},
	}
	packed, _ := args.Pack(&UserOperation{
		op.Sender,
		op.Nonce,
		op.InitCode,
		op.CallData,
		op.CallGasLimit,
		op.VerificationGasLimit,
		op.PreVerificationGas,
		op.MaxFeePerGas,
		op.MaxPriorityFeePerGas,
		op.PaymasterAndData,
		op.Signature,
	})

	enc := hexutil.Encode(packed)
	enc = "0x" + enc[66:]
	return hexutil.MustDecode(enc)
}

func (op *UserOperation) PackForSign() []byte {
	args := abi.Arguments{
		{Name: "sender", Type: address},
		{Name: "nonce", Type: uint256},
		{Name: "hashInitCode", Type: bytes32},
		{Name: "hashCallData", Type: bytes32},
		{Name: "callGasLimit", Type: uint256},
		{Name: "verificationGasLimit", Type: uint256},
		{Name: "preVerificationGas", Type: uint256},
		{Name: "maxFeePerGas", Type: uint256},
		{Name: "maxPriorityFeePerGas", Type: uint256},
		{Name: "hashPaymasterAndData", Type: bytes32},
	}
	packed, _ := args.Pack(
		op.Sender,
		op.Nonce,
		crypto.Keccak256Hash(op.InitCode),
		crypto.Keccak256Hash(op.CallData),
		op.CallGasLimit,
		op.VerificationGasLimit,
		op.PreVerificationGas,
		op.MaxFeePerGas,
		op.MaxPriorityFeePerGas,
		crypto.Keccak256Hash(op.PaymasterAndData),
	)

	return packed
}

func (op *UserOperation) UserOpHash(entryPoint common.Address, chainID *big.Int) common.Hash {
	return crypto.Keccak256Hash(
		crypto.Keccak256(op.PackForSign()),
		common.LeftPadBytes(entryPoint.Bytes(), 32),
		common.LeftPadBytes(chainID.Bytes(), 32),
	)
}

func (op *UserOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Sender               string `json:"sender"`
		Nonce                string `json:"nonce"`
		InitCode             string `json:"initCode"`
		CallData             string `json:"callData"`
		CallGasLimit         string `json:"callGasLimit"`
		VerificationGasLimit string `json:"verificationGasLimit"`
		PreVerificationGas   string `json:"preVerificationGas"`
		MaxFeePerGas         string `json:"maxFeePerGas"`
		MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
		PaymasterAndData     string `json:"paymasterAndData"`
		Signature            string `json:"signature"`
	}{
		Sender:               op.Sender.String(),
		Nonce:                hexutil.EncodeBig(op.Nonce),
		InitCode:             hexutil.Encode(op.InitCode),
		CallData:             hexutil.Encode(op.CallData),
		CallGasLimit:         hexutil.EncodeBig(op.CallGasLimit),
		VerificationGasLimit: hexutil.EncodeBig(op.VerificationGasLimit),
		PreVerificationGas:   hexutil.EncodeBig(op.PreVerificationGas),
		MaxFeePerGas:         hexutil.EncodeBig(op.MaxFeePerGas),
		MaxPriorityFeePerGas: hexutil.EncodeBig(op.MaxPriorityFeePerGas),
		PaymasterAndData:     hexutil.Encode(op.PaymasterAndData),
		Signature:            hexutil.Encode(op.Signature),
	})
}
