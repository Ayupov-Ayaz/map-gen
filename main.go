package main

import (
	"github.com/releaseband/map-gen/cli/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}