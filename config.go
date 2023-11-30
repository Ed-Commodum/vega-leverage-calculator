package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
		val := floatPrompt("Input the value of Tau: \n")
		if val <= 0 || val > 1 {
			log.Fatal("Error: Please provide a valid value for tau, valid range: 0 < tau <= 1")
		}
		tau = val
	}
	if lambda == 0.0 {
		val := floatPrompt("Input the value of Lambda (risk aversion parameter): \n")
		if val <= 1e-8 || val > 0.1 {
			log.Fatal("Error: Please provide a valid value for lambda, valid range: 1e-8 <= lambda < 0.1")
		}
		lambda = val
	}
	if sigma == 0.0 {
		val := floatPrompt("Input the value of Sigma: \n")
		if val <= 0 {
			log.Fatal("Error: Please provide a valid value for sigma (annualized volatilty), valid range: 0 < sigma")
		}
		sigma = val
	}

	return &Config{
		Tau:    tau,
		Lambda: lambda,
		Sigma:  sigma,
		Mu:     mu,
		R:      r,
	}
}

func floatPrompt(label string) float64 {

	var str string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label)
		str, _ = reader.ReadString('\n')
		if str != "" {
			break
		}
	}
	log.Printf("str: %v", str)
	val, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		log.Fatal("Error parsing user input: %v", err)
	}
	return val
}
