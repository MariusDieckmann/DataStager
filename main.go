package main

import (
	"log"

	"github.com/MariusDieckmann/DataStager/cmd"
	"github.com/spf13/viper"
)

//Version string
var Version string

func main() {
	viper.Set("Version", Version)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cmd.Execute()
}
