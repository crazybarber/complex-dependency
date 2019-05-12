package main

import (
	"docugraphy/config"
	"docugraphy/repository"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/julienschmidt/httprouter"
)

var applicationName = os.Args[0]

const (
	installCommand string = "install"
	runCommand            = "run"
)

func main() {
	var err error = nil

	/* PASS ARGV */
	if len(os.Args) < 2 {
		usage()
		log.Fatalf("Command missing\n")
	}

	switch command := os.Args[1]; command {
	case installCommand:
		handleConfigFile()
		log.Println("Installing db schema")
		err = install()
	case runCommand:
		handleConfigFile()
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
	err := repository.Create()
	if nil != err {
		return err
	}
	return nil
}

func usage() {
	usage := "Usage: " + applicationName + " <command>; " +
		fmt.Sprintf("Possible commands: %s, %s", installCommand, runCommand)
	log.Println(usage)
}

func handleConfigFile() {
	flagSet := flag.NewFlagSet("installFlags", flag.ExitOnError)
	configFileLocation :=  *(flagSet.String("configFile", "config.json", "Path to json config file"))
	err := flagSet.Parse(os.Args[2:])
	if nil != err {
		log.Fatalln("Command line flags processing error: " + err.Error())
	}
	err = config.Load(configFileLocation)
	if nil != err {
		log.Fatalln("Config problem: " + err.Error())
	}

}

func startService() {

	router := httprouter.New()
	router.POST("/source_system", nil)
	router.GET("/source_systems", nil)
	log.Fatal(http.ListenAndServe(":8080", router))

}
