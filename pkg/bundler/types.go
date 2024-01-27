package bundler

import "math/big"

type GasEstimates struct {
	PreVerificationGas   *big.Int `json:"preVerificationGas"`
	VerificationGasLimit *big.Int `json:"verificationGasLimit"`
	CallGasLimit         *big.Int `json:"callGasLimit"`
}
