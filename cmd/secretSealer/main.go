package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"

	"github.com/sebastianRau/deployer/pkg/templating"
)

var (
	version = "1.0.0 abcdef"
)

func main() {
	var (
		text        = flag.String("t", "", "Text to Encrypt") //toDo remove default
		keyFile     = flag.String("k", "", "Key file")
		generateKey = flag.Bool("g", false, "generate Key file")
	)
	flag.Parse()

	var (
		key []byte
	)

	if *generateKey {
		f, err := os.OpenFile(*keyFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		key = GenerateKey()
		f.Write(key)

	} else {
		f, err := os.Open(*keyFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		key, err = io.ReadAll(f)
		if err != nil {
			panic(err)
		}
	}

	enc := templating.Encrypt(*text, key)
	dec := templating.Decrypt(enc, key)

	fmt.Println(dec + ":")
	fmt.Println(enc)

}

func GenerateKey() []byte {
	n := 32
	digits := "0123456789"
	capLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letters := "abcdefghijklmnopqrstuvwxyz"

	letterRunes := digits + capLetters + letters

	b := make([]byte, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return b
}
