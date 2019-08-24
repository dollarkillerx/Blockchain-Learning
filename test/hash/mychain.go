package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func calculateHash(toBeHashed string) string {
	hashInBytes := sha256.Sum256([]byte(toBeHashed))
	s := hex.EncodeToString(hashInBytes[:])
	return s
}

func main() {
	hash := calculateHash("123456")
	log.Println(hash)
}