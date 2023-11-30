package main

import (
	risk "code.vegaprotocol.io/quant/riskmodelbs"
	"log"
)

func init() {
	defineFlags()
}

func main() {
	config := parseFlags()

	modelParams := risk.ModelParamsBS{
		Mu:    config.Mu,
		R:     config.R,
		Sigma: config.Sigma,
	}

	riskFactors := risk.RiskFactorsForward(config.Lambda, config.Tau, modelParams)

	log.Printf("Estimated max long leverage: %v\n", 1/riskFactors.Long)
	log.Printf("Estimated max short leverage: %v\n", 1/riskFactors.Short)
}
