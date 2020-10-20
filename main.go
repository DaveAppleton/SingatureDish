package main

import (
	"fmt"
	"log"

	"github.com/DaveAppleton/ether_go/ethKeys"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

func main() {
	mykey := ethKeys.NewKey("keys/myNiceShinyKey")
	err := mykey.RestoreOrCreate()
	if err != nil {
		log.Fatal(err)
	}

	NiceLongMessageToSign := "A nice long message to sign"

	hash := signHash([]byte(NiceLongMessageToSign))
	khash := crypto.Keccak256([]byte(NiceLongMessageToSign))
	fmt.Println("My address = ", mykey.PublicKeyAsHexString())
	fmt.Println("hash = ", "0x"+common.Bytes2Hex(khash))
	sig, err := crypto.Sign(hash, mykey.GetKey())
	v := sig[64]
	if v < 27 {
		fmt.Printf("original signature : 0x%02x\n", sig)
		fmt.Println("updating to make last byte 27 or 28 (not usually needed)")
		v += 27
		sig[64] = v
		fmt.Printf("modified signature : 0x%02x\n", sig)
	} else {
		fmt.Printf("signature : %s\n", "0x"+common.Bytes2Hex(sig))
	}
	r := sig[:0x20]
	fmt.Println("r = 0x" + common.Bytes2Hex(r))
	s := sig[0x20:0x40]
	fmt.Println("s = 0x" + common.Bytes2Hex(s))
	fmt.Printf("v = 0x%02x\n", v)
	fmt.Println("Verify the results at https://rinkeby.etherscan.io/address/0x84a8f06143749e5ba367052c23dc6d5166a83dc5#readContract")
}

/*
// 0x60b05BFb1d189F7BBDED27d3A491d7488da66793
// 0x84A8f06143749e5BA367052C23dC6d5166A83DC5

{
	"address": "0xb9e30dbaa784f84374c80ae1d4ea5a3d967b845f",
	"msg": "A nice long message to sign",
	"sig": "0xfbbac5c6fa81b9ec4226da3969156f658aa318a88fa5383673f504208a68351a1671620aba1c8459cffed5ac6e2220ebdc0ca82d8e1fedf63934a87676cc85a31c",
	"version": "2"
}

*/
