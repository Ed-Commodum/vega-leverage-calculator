package main

import (
	risk "code.vegaprotocol.io/quant/riskmodelbs"
	"log"
)

func init() {
	log.Printf("Initializing...")
	defineFlags()
}

func main() {
	config := parseFlags()

}
