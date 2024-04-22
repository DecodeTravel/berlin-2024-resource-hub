package main

import (
	"fmt"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/hypersdk/codec"
	"github.com/chain4travel/caminotravelvm/auth"
	"github.com/chain4travel/caminotravelvm/consts"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Missing <private_key> argument!")
		return
	}
	pk := new(secp256k1.PrivateKey)
	privateKey := os.Args[1]
	if err := pk.UnmarshalText([]byte("\"" + privateKey + "\"")); err != nil {
		fmt.Errorf("failed to parse private key: %w", err)
	}
	addr, err := codec.AddressBech32(consts.HRP, auth.NewSECP256K1Address(*pk.PublicKey()))
	if err != nil {
		fmt.Errorf("failed to create address: %w", err)
	}
	fmt.Printf("Adderss: %s\n", addr)
}
