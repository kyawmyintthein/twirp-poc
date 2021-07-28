package main

import (
	"flag"

	"github.com/kyawmyintthein/twirp-poc/app"
)

func main() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", ".env", "absolute path to the configuration file")
	flag.Parse()

	application, err := app.New(configFilePath)
	if err != nil {
		panic(err)
	}

	// initiate dependency modules before start
	err = application.Init()
	if err != nil {
		panic(err)
	}

	// start servers
	application.Start()
	select {}
}
