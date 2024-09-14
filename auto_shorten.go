package main

import (
	"log"
	"crypto/rand"
	"math/big"
	"bytes"
)

func getRandomIndexFromString(s string) byte {
	length := len(s)
	n, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
	if err != nil {
		log.Fatal(err)
	}
	return s[n.Int64()]
}

func generateAutomaticShortLink() string {
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	numbers := "1234567890"
	var b bytes.Buffer
	for i := 0; i < 4; i++ {
		if i == 1 {
			b.WriteByte(getRandomIndexFromString(upperCase))
		} else if i == 2 {
			b.WriteByte(getRandomIndexFromString(numbers))
		} else {
			b.WriteByte(getRandomIndexFromString(lowercase))	
		}
	}
	return b.String()
}