package main

import (
	"bytes"
	"fmt"

	gittools "github.com/sebastianRau/deployer/pkg/gitTools"
)

func main() {

	trash := bytes.NewBufferString("")

	_, _, err := gittools.UpdateNoKeyurl(
		"https://github.com/kdo-wildsau/KW_MissionBuilder.git",
		"./temp/KW_MisisonBuilder",
		trash,
	)

	fmt.Println(trash.String())

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
