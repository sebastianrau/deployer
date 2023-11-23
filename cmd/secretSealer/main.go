package main

import (
	"flag"
	"fmt"
	"os"

	enc "github.com/sebastianrau/go-easyConfig/pkg/encryption"
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

	if *keyFile == "" {
		fmt.Println("no kexfile given")
		os.Exit(1)
	}

	if *generateKey {
		err := enc.CreateKeyFile(*keyFile)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	} else {
		if *text == "" {
			fmt.Println("no text to encrypt given")
			os.Exit(1)
		}

		key, err := os.ReadFile(*keyFile)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		val, err := enc.EncryptString(*text, key)

		fmt.Printf("%s: \n", *text)
		fmt.Println(val)
	}
}
