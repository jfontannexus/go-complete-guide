package main

import (
	"fmt"
	"math"
)

func main() {
	// Call the function

	var investedAmount float64 = 1000
	var expectedReturnRate float64 = 5.5
	var years float64 = 10
	var result = CalculateInvestment(investedAmount, expectedReturnRate, years)
	fmt.Println("The future value of the investment is: ", result)

}
func CalculateInvestment(investedAmount float64, expectedReturnRate float64, years float64) float64 {
	// Calculate the investment
	return investedAmount *
		math.Pow(1+expectedReturnRate/100, years)

}
