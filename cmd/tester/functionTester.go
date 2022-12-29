package main

import (
	"bytes"
	"fmt"

	filetools "github.com/sebastianRau/deployer/pkg/fileTools"
)

func main() {

	trash := bytes.NewBufferString("")

	/*
			hash, newHash, err := gittools.Pull(
				"./temp/KW_BEM_BasicScripts",
				"./test/keys/id_deployment",
				"",
				trash,
			)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Printf("%v\n", hash)
			fmt.Printf("new hash: %v\n", newHash)

			err = filetools.ParseTemplate(
				"./test/kw_mod_builder_1.4.1.go.tpl",
				"./test/kw_mod_builder_1.4.1.data.json",
				"./temp/kw_mod_builder_1.4.1.json",
				trash,
			)

			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}

			fmt.Println("LOG:")
			fmt.Println(trash.String())


		err := ostools.Exec("ls", "./test/", trash, "-la")
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}
		fmt.Println("Dir:")
		fmt.Println(trash.String())


	*/

	err := filetools.ParseTemplate(
		"./test/kw_bem_make.json.tpl",
		"./test/kw_bem_make.data.json",
		"./test/kw_bem_make.json",
		trash,
	)

	fmt.Println(trash.String())

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
