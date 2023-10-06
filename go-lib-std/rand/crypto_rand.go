package main

import (
	crypto_rand "crypto/rand"
	"fmt"
	"math/big"
)

func Handle_error(err error) {
	if err != nil {
		panic(err)
	}
}

func crypt_num() {
	// Cryptographic random numbers
	var max *big.Int = big.NewInt(6)
	// big is a package for high-precision arithmetic
	for i := 0; i < 10; i++ {
		crypt_rand_num, err := crypto_rand.Int(crypto_rand.Reader, max)
		Handle_error(err)
		// Move the Offset of 0 by adding 1
		crypt_num := crypt_rand_num.Add(crypt_rand_num, big.NewInt(1))
		fmt.Println(crypt_num)
	}

}
