package main

import (
	crypto_rand "crypto/rand"
	"fmt"
	"math/rand"
	"time"
)

func randomReadBytes() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand_byte := make([]byte, 5)
	fmt.Println(rand_byte)
	_, err := r.Read(rand_byte)
	Handle_error(err)
	fmt.Println(rand_byte)
	fmt.Printf("%c \n", rand_byte)

	crypt_rand_byte := make([]byte, 5)
	fmt.Println(crypt_rand_byte)
	_, err = crypto_rand.Read(crypt_rand_byte)
	Handle_error(err)
	fmt.Println(crypt_rand_byte)
	fmt.Printf("%c \n", crypt_rand_byte)
}
