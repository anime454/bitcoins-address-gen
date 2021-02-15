package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"golang.org/x/crypto/ripemd160"

	"github.com/btcsuite/btcutil/base58"
)

func main() {
	passphrase := getPassphrase()
	prKey := genPrivateKey(passphrase)
	encodedKey := encodedPrivateKey(prKey[:])
	fmt.Printf("final encode private key : %s \n", encodedKey)

	newX, newY := secp256k1.S256().ScalarBaseMult(prKey[:])
	fmt.Printf("newX is : %x \nnewY is : %x \n", newX, newY)

	evenChecker := big.NewInt(0).Mod(newY, big.NewInt(2))
	newHeader := newHeaderBit(evenChecker)
	fmt.Printf("new header is : %x\n", newHeader)

	newPrivatekey := fmt.Sprintf("%s%s", newHeader, newX.Bytes())
	fmt.Printf("new Private key is : %x \n", newPrivatekey)
	a := sha256.Sum256([]byte(newPrivatekey))
	fmt.Printf("private ket encrypted is : %x \n", a)
	hasher := ripemd160.New()
	hasher.Write(a[:])
	hashBytes := hasher.Sum(nil)

	versionNumber := []byte("\x00")
	prKeyAndVersion := fmt.Sprintf("%s%s", versionNumber, hashBytes)
	checksum := getChecksum([]byte(prKeyAndVersion))
	fmt.Printf("Checksum : %x \n", checksum)
	prKeyAndVersionAndChecksum := fmt.Sprintf("%s%s%s", versionNumber, []byte(hashBytes), checksum)
	fmt.Printf("Private key, version and checksum : %x \n", prKeyAndVersionAndChecksum)
	encoded := base58.Encode([]byte(prKeyAndVersionAndChecksum))
	fmt.Println("Encoded Data:", encoded)
}

func encodedPrivateKey(prKey []byte) string {
	fmt.Printf("Private key : %x \n", prKey)
	versionNumber := []byte("\x80")
	compressionFlag := []byte("\x01")

	prKeyAndVersion := fmt.Sprintf("%s%s%s", versionNumber, prKey, compressionFlag)
	fmt.Printf("Private key and version : %x \n", prKeyAndVersion)

	checksum := getChecksum([]byte(prKeyAndVersion))
	fmt.Printf("Checksum : %x \n", checksum)
	prKeyAndVersionAndChecksum := fmt.Sprintf("%s%s%s%s", versionNumber, prKey, compressionFlag, checksum)
	fmt.Printf("Private key, version and checksum : %x \n", prKeyAndVersionAndChecksum)

	encoded := base58.Encode([]byte(prKeyAndVersionAndChecksum))
	fmt.Println("Encoded Data:", encoded)

	return encoded
}

func newHeaderBit(evenChecker *big.Int) []byte {
	newHeader := []byte("")
	if evenChecker.Cmp(big.NewInt(0)) == 0 { // IN CASE OF EVEN
		newHeader = []byte("\x02")
	} else {
		newHeader = []byte("\x03")
	}
	return newHeader
}

func getChecksum(input []byte) []byte {
	first := sha256.Sum256(input)
	second := sha256.Sum256(first[:])
	return second[:4]
}

func genPrivateKey(passphrase string) [32]byte {
	privateKey := sha256.Sum256([]byte(passphrase))
	return privateKey
}

func getPassphrase() string {
	passphrase := os.Args[1:]
	allPassphrase := strings.Join(passphrase, " ")
	return allPassphrase
}
