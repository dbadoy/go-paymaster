package signer

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/dbadoy/go-paymaster/pkg/userop"
)

type Signer struct {
	entryPoint common.Address
	chainID    *big.Int

	address    common.Address
	privateKey *ecdsa.PrivateKey
}

func New(pk string, entryPoint common.Address, chainID *big.Int) (*Signer, error) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, err
	}

	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKey)

	return &Signer{
		entryPoint,
		chainID,
		address,
		privateKey,
	}, nil
}

func (s *Signer) Address() common.Address {
	return s.address
}

func (s *Signer) Sign2UserOperation(op *userop.UserOperation) ([]byte, error) {
	preUserOpHash := op.UserOpHash(s.entryPoint, s.chainID)
	return s.Sign2Message(preUserOpHash.Bytes())
}

func (s *Signer) Sign2Message(msg []byte) ([]byte, error) {
	fullMsg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)
	hash := crypto.Keccak256([]byte(fullMsg))
	signature, err := crypto.Sign(hash, s.privateKey)
	if err != nil {
		return nil, err
	}
	signature[64] += 27
	return signature, nil
}
