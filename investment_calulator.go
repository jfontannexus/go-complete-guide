package main

import (
	"fmt"
	"math"
)

func main() {
	// Call the function

	var investedAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10
	var result = CalculateInvestment(investedAmount, expectedReturnRate, years)
	fmt.Println("The future value of the investment is: ", result)

}
func CalculateInvestment(investedAmount int, expectedReturnRate float64, years int) float64 {
	// Calculate the investment
	var futureValue = float64(investedAmount) * math.Pow(1+expectedReturnRate/100)
	return futureValue

}
