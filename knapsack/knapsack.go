package knapsack

import (
	"math/big"
)

func (c *Cipher) Decrypt(cipher *big.Int) *Binary {
	mod := new(big.Int)
	mult := new(big.Int)

	mod.Set(&c.PrivateKey.Mod)
	mult.Set(&c.PrivateKey.Mult)

	inverse := mult.ModInverse(mult, mod)

	cipher.Mul(cipher, inverse).Mod(cipher, mod)

	return c.binaryRep(cipher)
}

func (c *Cipher) Encrypt(b *Binary) *big.Int {
	mod := new(big.Int)
	mult := new(big.Int)

	mod.Set(&c.PrivateKey.Mod)
	mult.Set(&c.PrivateKey.Mult)

	var sum int64 = 0
	for i, v := range c.FindPublicKey().Key {
		if (*b)[i] == true {
			sum += v
		}
	}
	return big.NewInt(sum)
}

func (c *Cipher) binaryRep(x *big.Int) *Binary {
	value := x.Int64()
	key := c.PrivateKey.Key

	length := len(key)
	rep := make([]bool, length)

	for i := length - 1; i >= 0; i-- {

		kex := key[i]

		rep[i] = false

		if kex <= value {
			value -= kex
			rep[i] = true
		}
	}

	result := Binary(rep)
	return &result
}

func (c *Cipher) FindPublicKey() *PublicKey {
	if c.PublicKey.Key != nil {
		return &c.PublicKey
	}

	key := c.PrivateKey.Key
	length := len(key)

	mod := new(big.Int)
	mult := new(big.Int)

	mod.Set(&c.PrivateKey.Mod)
	mult.Set(&c.PrivateKey.Mult)

	public := make([]int64, length)

	for i, v := range key {
		value := big.NewInt(v)
		public[i] = value.Mul(value, mult).Mod(value, mod).Int64()
	}

	c.PublicKey.Key = public
	return &c.PublicKey
}
