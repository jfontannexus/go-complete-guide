package main

import (
	"fmt"
	"math"
)

func main() {
	// Call the function
	const inflationRate, label = 2.0, "Investment Calculator"
	var investedAmount, years float64
	expectedReturnRate := 5.5 //infered type
	scanAmount(investedAmount, "Enter the amount to invest: ")
	scanAmount(years, "Enter the number of years to invest: ")
	var result = CalculateInvestment(investedAmount, expectedReturnRate, years)
	var inflationAdjustedResult = CalculateInvestment(investedAmount, expectedReturnRate-inflationRate, years)
	printResults(label, result, inflationAdjustedResult)
}

func scanAmount(amount float64, prompt string) {
	// Read the amount from user input
	print(prompt)
	fmt.Scanln(&amount)
}
func CalculateInvestment(investedAmount float64, expectedReturnRate float64, years float64) float64 {
	// Calculate the investment
	return investedAmount * math.Pow(1+expectedReturnRate/100, years)

}
func printInvestmentResult(label string, result float64) {
	// Print the result
	println(label, ":", result)
}
func printResults(label string, result float64, inflationAdjustedResult float64) {
	// Print the results
	printInvestmentResult(label, result)
	printInvestmentResult(label+" (Inflation Adjusted)", inflationAdjustedResult)
}
