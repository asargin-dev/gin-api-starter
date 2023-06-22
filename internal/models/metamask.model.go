package models

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Metamask struct {
	SignedStr string `json:"signedStr" binding:"required"`
}

func (meta *Metamask) ExtractPublicAddress() common.Address {

	// Verify the signed string
	signature := meta.SignedStr[2:]

	// Parse the signature from the signed string
	sigBytes, err := hex.DecodeString(signature)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sigBytes)

	hash := crypto.Keccak256Hash(sigBytes)
	// Create a new ecdsa.PublicKey with the recovered signature values
	publicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), sigBytes)
	if err != nil {
		fmt.Println(err)
	}

	// Convert the ecdsa.PublicKey to crypto.Pubkey and extract the address
	publicKey := (*ecdsa.PublicKey)(publicKeyECDSA)
	addressFromPubKey := crypto.PubkeyToAddress(*publicKey)

	// Print the public address
	fmt.Println("Public Address:", addressFromPubKey)

	return addressFromPubKey
}
