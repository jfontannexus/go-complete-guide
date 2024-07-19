package main

import (
	"math"
)

func main() {
	// Call the function
	const inflationRate = 2.0
	var investedAmount, years float64 = 1000, 10 //multiple explicit type
	expectedReturnRate := 5.5                    //infered type
	var label, offset = "Investment", 10         //multiple inferred type
	var result = CalculateInvestment(investedAmount, expectedReturnRate, years)
	var inflationAdjustedResult = CalculateInvestment(investedAmount, expectedReturnRate-inflationRate, years)
	printResults(offset, label, result, inflationAdjustedResult)
}
func CalculateInvestment(investedAmount float64, expectedReturnRate float64, years float64) float64 {
	// Calculate the investment
	return investedAmount * math.Pow(1+expectedReturnRate/100, years)

}
func printInvestmentResult(offset int, label string, result float64) {
	// Print the result
	println(label, ":", result)
}
func printResults(offset int, label string, result float64, inflationAdjustedResult float64) {
	// Print the results
	printInvestmentResult(offset, label, result)
	printInvestmentResult(offset, label+" (Inflation Adjusted)", inflationAdjustedResult)
}
