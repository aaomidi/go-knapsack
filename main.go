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

	//private := knapsack.PrivateKey{
	//	Key:  []int64{2, 7, 11, 21, 42, 89, 180, 354},
	//	Mod:  *big.NewInt(881),
	//	Mult: *big.NewInt(588),
	//}

	cipher := knapsack.Cipher{
		PrivateKey: private,
	}

	x := cipher.Decrypt(big.NewInt(20))
	fmt.Println(x.BtoStr())

	x = cipher.Decrypt(big.NewInt(29))
	fmt.Println(x.BtoStr())

	cipher.FindPublicKey()
	fmt.Println(cipher.PublicKey.Key)

	input := []bool{false, true, true, false, false, false, false, true}
	binaryInput := knapsack.Binary(input)

	y := cipher.Encrypt(&binaryInput)
	fmt.Println(y)
}
