package main

// this code is my very first code in golang when it comes to cryptography
// I'm sorry for any inefficiency that is implemented here
// if you have any improvements don't hesitate to send a pull request

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// karudo: w3

func generate(length int) {
	// here we generate the key
	privatekey, err := rsa.GenerateKey(rand.Reader, length)
	if err != nil {
		fmt.Printf("Cannot generate RSA key, exiting...\n")
		os.Exit(1)
	}
	publickey := &privatekey.PublicKey

	// now dump the private key to a file
	var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)
	PRIVblock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	PEMPriv, err := os.Create("private.pem")
	if err != nil {
		fmt.Printf("error when creating the file private.pem: %s \n", err)
		os.Exit(1)
	}
	err = pem.Encode(PEMPriv, PRIVblock)
	if err != nil {
		fmt.Printf("error when encoding the file private.pem: %s \n", err)
		os.Exit(1)
	}
	// now we dump the public key to a file
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		fmt.Printf("error when dumping the public key: %s \n", err)
		os.Exit(1)
	}
	PUBblock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	PEMPub, err := os.Create("public.pem")
	if err != nil {
		fmt.Printf("error when creating the file public.pem: %s \n", err)
		os.Exit(1)
	}
	err = pem.Encode(PEMPub, PUBblock)
	if err != nil {
		fmt.Printf("error when encoding the file public.pem: %s \n", err)
		os.Exit(1)
	}
	fmt.Printf("Saved output to private.pem and public.pem \n")
}
func main() {

	var choice int

	fmt.Println("Here are the available key lengths : ")
	fmt.Println("1 -> 512 bits")
	fmt.Println("2 -> 1024 bits")
	fmt.Println("3 -> 2048 bits")
	fmt.Print("Enter the number corresponding to the desired key length : ")

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		generate(512)
	case 2:
		generate(1024)
	default:
		generate(2048)
	}
}
