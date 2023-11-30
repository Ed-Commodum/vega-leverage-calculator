package main

import (
	"flag"
	"log"
)

type Config struct {
	Tau    float64
	Lambda float64
	Sigma  float64
	Mu     float64
	R      float64
}

const (
	defaultTau    = 0.0
	defaultLambda = 0.0
	defaultSigma  = 0.0
	defaultMu     = 0.0
	defaultR      = 0.0
)

var (
	tau    float64
	lambda float64
	sigma  float64
	mu     float64
	r      float64
)

func defineFlags() {
	flag.Float64Var(&tau, "tau", defaultTau, "A value representing the time horizon expressed as a fraction of a year.")
	flag.Float64Var(&lambda, "lambda", defaultLambda, "The risk aversion parameter used in the Expected Shortfall calculation.")
	flag.Float64Var(&sigma, "sigma", defaultSigma, "Annualized volatility.")
	flag.Float64Var(&mu, "mu", defaultMu, "Annualized growth rate of the underlying asset.")
	flag.Float64Var(&r, "r", defaultR, "Risk free rate of return.")
}

func parseFlags() *Config {
	flag.Parse()

	if tau == 0.0 {
		log.Fatal("Error: Please provide a valid value for tau, valid range: 0 < tau <= 1")
	}
	if lambda == 0.0 {
		log.Fatal("Error: Please provide a valid value for lambda (risk aversion parameter).")
	}
	if sigma == 0.0 {
		log.Fatal("Error: Please provide a valid value for sigma (annualized volatilty), valid range: 0 < sigma")
	}

	return &Config{
		Tau:    tau,
		Lambda: lambda,
		Sigma:  sigma,
		Mu:     mu,
		R:      r,
	}
}

func getFlag()
