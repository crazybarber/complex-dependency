package main

import (
	"complex-dependency/config"
	"complex-dependency/repository/postgres"
	"flag"
	"fmt"
	"log"
	"os"
)

var applicationName = os.Args[0]

const (
	installCommand string = "install"
	runCommand            = "run"
)

func main() {

	/* PASS ARGV */
	if len(os.Args) < 2 {
		usage()
		log.Fatalf("Command missing\n %s\n")
	}

	flagSet := flag.NewFlagSet("installFlags", flag.ExitOnError)
	configFileLocationPtr := flagSet.String("configFile", "config.json", "Path to json config file")
	err := flagSet.Parse(os.Args[1:])
	if nil != err {
		log.Fatalln("Command line flags processing error: " + err.Error())
	}

	err = config.Load(*configFileLocationPtr)
	if nil != err {
		log.Fatalln("Config problem: " + err.Error())
	}

	switch command := os.Args[1]; command {
	case installCommand:
		log.Println("Installing db schema")
		err = install()
	case runCommand:

	default:
		usage()
		log.Fatalf("Unknown command %s\n", command)
	}

	if nil != err {
		log.Fatalf("Problems while processing: %s\n", err.Error())
	}

	os.Exit(0)
}

func install() error {
	var err error = nil

	switch dbConfig := config.GetConfig().Db; dbConfig {
	case postgres.DbModuleConfigName:
		err = postgres.SetupSchema()
	default:
		err = fmt.Errorf("Unknown database id %s\n", dbConfig)
	}

	if nil != err {
		return err
	}
	return nil
}

func usage() {
	usage := "Usage: " + applicationName + " <command>\n" +
		fmt.Sprintf("Possible commands: %s, %s", installCommand, runCommand)
	fmt.Println(usage)
}
