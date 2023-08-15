package main

import (
	"encoding/hex"
	"fmt"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

func main() {
	b2octh1, err := hex.DecodeString("01795EDF0D54DB760F156D0DAC04C0322B3A204224")
	if err != nil {
		panic(err)
	}
	fmt.Println(b2octh1, len(b2octh1))

	x, err := hex.DecodeString("009A4D6792295A7F730FC3F2B49CBC0F62E862272F")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(x))

	kb := secp256k1.NonceRFC6979(x, b2octh1, nil, nil, 0).Bytes()
	khex := hex.EncodeToString(kb[:])
	fmt.Println("k", khex, len(khex))
}
