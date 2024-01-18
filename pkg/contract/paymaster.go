package contract

import (
	"github.com/dbadoy/go-paymaster/pkg/signer"
	"github.com/dbadoy/go-paymaster/pkg/userop"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type PaymasterCaller interface {
	GetHash(userOp *userop.UserOperation, paymasterID common.Address) ([32]byte, error)
}

var (
	abi_address, _ = abi.NewType("address", "", nil)
	abi_bytes, _   = abi.NewType("bytes", "", nil)
)

type Paymaster struct {
	address common.Address
	caller  PaymasterCaller
	funder  *signer.Signer
}

func (p *Paymaster) ContractAddress() common.Address {
	return p.address
}

func (p *Paymaster) Sign(userOp *userop.UserOperation) error {
	hashBytes, err := p.caller.GetHash(userOp, p.address)
	if err != nil {
		return err
	}

	var hash common.Hash = hashBytes

	sign, err := p.funder.Sign2Message(hash.Bytes())
	if err != nil {
		return err
	}

	args := abi.Arguments{
		{Name: "address", Type: abi_address},
		{Name: "bytes", Type: abi_bytes},
	}

	packed, _ := args.Pack(
		p.address,
		sign,
	)

	paymaseterAndData := make([]byte, len(p.address)+len(packed))
	copy(paymaseterAndData[:len(p.address)], p.address.Bytes())
	copy(paymaseterAndData[len(p.address):], packed)

	userOp.PaymasterAndData = paymaseterAndData
	return nil
}
