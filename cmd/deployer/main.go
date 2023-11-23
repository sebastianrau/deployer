package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/sebastianrau/deployer/pkg/steps"
)

var (
	version = "1.0.0 abcdef"
)

func main() {
	var (
		configTemplateFile = flag.String("configFile", "", "Config File") //toDo remove default
		configDataFile     = flag.String("configData", "", "Data file")
		encryptionFile     = flag.String("keyfile", "", "encryption Keyfile")
		configOutputFile   = flag.String("configOutput", "", "Output json file")
		verbose            = flag.Bool("v", false, "verbose output")
		print              = flag.Bool("print", false, "print yaml only")
	)
	flag.Parse()

	var (
		err error
	)

	if *configTemplateFile == "" {
		panic("No Config File given")
	}

	// read and unmarshal template
	configSteps, err := steps.UnmarshalConfigTemplate(*configTemplateFile, *configDataFile, *encryptionFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *configOutputFile != "" {
		err := configSteps.WriteConfigToFile(*configOutputFile)
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}

	if *print {
		err := configSteps.WriteConfigTo(os.Stdout)
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}

	fmt.Printf("Deployer v%s\n", version)

	start := time.Now()
	err = configSteps.Exceute(os.Stdout, *verbose)
	if err != nil {
		fmt.Println(err.Error())
	}
	stop := time.Now()
	elapsed := stop.Sub(start)

	fmt.Printf("Time: %s\n", elapsed.Round(time.Millisecond))

}
