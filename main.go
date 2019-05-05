package main

import (
	"complex-dependency/config"
	"complex-dependency/repository/postgres"
	"flag"
	"fmt"
	"os"
)

func main() {

	/* FLAGS */
	configFileLocationPtr := flag.String("configFile", "config.json", "Path to json config file")
	flag.Parse()

	err := config.Load(*configFileLocationPtr)
	if nil != err {
		fmt.Println("Config problem: " + err.Error())
		os.Exit(1)
	}

	err = postgres.SetupSchema()
	if nil != err {
		fmt.Println("DB problem: " + err.Error())
		os.Exit(1)
	}

}
