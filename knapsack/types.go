package knapsack

import (
	"fmt"
	"math/big"
	"strings"
)

type PrivateKey struct {
	Key  []int64
	Mod  big.Int
	Mult big.Int
}

type PublicKey struct {
	Key []int64
}

type Cipher struct {
	PrivateKey PrivateKey
	PublicKey  PublicKey
}

type Binary []bool

func (b *Binary) BtoStr() string {
	var sb strings.Builder

	for _, v := range *b {
		if v {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
	}
	return sb.String()
}

func StrToB(str *string, length int) (*Binary, error) {
	result := make([]bool, length)

	skip := length - len(*str)

	for i, v := range *str {
		if v == '1' {
			result[i+skip] = true
		} else if v == '0' {
			result[i+skip] = false
		} else {
			return nil, fmt.Errorf("malformed input: %c", v)
		}
	}
	x := Binary(result)
	return &x, nil
}
