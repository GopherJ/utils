package util

/**
 * prime.go
 * 2017-12-27 14:39:20
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
    "crypto/rand"
    "math/big"
)

func GenPrime(bits int)(*big.Int) {
   n, err := rand.Prime(rand.Reader, bits)
   if err != nil {
        panic(err)
   }
   return n
}

func NewBigInt(s string, b int)(*big.Int) {
    n := new(big.Int)
    n.SetString(s, b)
    return n
}

func IsProbablePrime(n *big.Int, p int)(bool) {
    return n.ProbablyPrime(p)
}
