package main

import (
	"log"
	"math"

	pd "code.vegaprotocol.io/quant/pricedistribution"
	risk "code.vegaprotocol.io/quant/riskmodelbs"
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
	log.Printf("Estimated max short leverage: %v\n\n", 1/riskFactors.Short)

	tau50, tau100, tau200, tau500 := config.Tau*0.5, config.Tau, config.Tau*2, config.Tau*5

	dist50 := modelParams.GetProbabilityDistribution(100.0, tau50)
	dist100 := modelParams.GetProbabilityDistribution(100.0, tau100)
	dist200 := modelParams.GetProbabilityDistribution(100.0, tau200)
	dist500 := modelParams.GetProbabilityDistribution(100.0, tau500)

	min50, max50 := pd.PriceRange(dist50, 0.999999)
	min100, max100 := pd.PriceRange(dist100, 0.999999)
	min200, max200 := pd.PriceRange(dist200, 0.999999)
	min500, max500 := pd.PriceRange(dist500, 0.999999)

	range50 := math.Floor(100 * (100.0 - min50))
	range100 := math.Floor(100 * (100.0 - min100))
	range200 := math.Floor(100 * (100.0 - min200))
	range500 := math.Floor(100 * (100.0 - min500))

	log.Printf("Tau Scaling: %v, Min price: %v, Max Price: %v, Price Rage (bp): %v\n", 0.5, min50, max50, range50)
	log.Printf("Tau Scaling: %v, Min price: %v, Max Price: %v, Price Rage (bp): %v\n", 1, min100, max100, range100)
	log.Printf("Tau Scaling: %v, Min price: %v, Max Price: %v, Price Rage (bp): %v\n", 2, min200, max200, range200)
	log.Printf("Tau Scaling: %v, Min price: %v, Max Price: %v, Price Rage (bp): %v\n\n", 5, min500, max500, range500)

	for _, basisPoints := range []float64{5, 10, 20, 30} {

		bidPrice := 100 - (basisPoints / 100)
		askPrice := 100 + (basisPoints / 100)

		bidP50 := pd.ProbabilityOfTrading(dist50, bidPrice, true, true, 0.0, 100.0)
		askP50 := pd.ProbabilityOfTrading(dist50, askPrice, false, true, 100.0, 600.0)
		bidP100 := pd.ProbabilityOfTrading(dist100, bidPrice, true, true, 0.0, 100.0)
		askP100 := pd.ProbabilityOfTrading(dist100, askPrice, false, true, 100.0, 600.0)
		bidP200 := pd.ProbabilityOfTrading(dist200, bidPrice, true, true, 0.0, 100.0)
		askP200 := pd.ProbabilityOfTrading(dist200, askPrice, false, true, 100.0, 600.0)
		bidP500 := pd.ProbabilityOfTrading(dist500, bidPrice, true, true, 0.0, 100.0)
		askP500 := pd.ProbabilityOfTrading(dist500, askPrice, false, true, 100.0, 600.0)

		log.Printf("-------------------- Basis points from best bid/ask: %v --------------------\n", basisPoints)
		log.Printf("Tau scaling: %v, bid probability: %v, ask probability: %v\n", 0.5, bidP50, askP50)
		log.Printf("Tau scaling: %v, bid probability: %v, ask probability: %v\n", 1, bidP100, askP100)
		log.Printf("Tau scaling: %v, bid probability: %v, ask probability: %v\n", 2, bidP200, askP200)
		log.Printf("Tau scaling: %v, bid probability: %v, ask probability: %v\n\n", 5, bidP500, askP500)
	}

}
