package main

import (
	"fmt"
	"math/big"

	"github.com/aaomidi/go-knapsack/knapsack"
)

func main() {
	private := knapsack.PrivateKey{
		Key:  []int64{3, 5, 10, 23},
		Mod:  *big.NewInt(47),
		Mult: *big.NewInt(8),
	}

	cipher := knapsack.Cipher{
		PrivateKey: private,
	}

	x := cipher.Decrypt(big.NewInt(20))
	fmt.Printf("Decrypted value of 20: %s\n", x.BtoStr())

	x = cipher.Decrypt(big.NewInt(29))
	fmt.Printf("Decrypted value of 29: %s\n", x.BtoStr())

	cipher.FindPublicKey()
	fmt.Printf("Public key: %v\n", cipher.PublicKey.Key)
}
