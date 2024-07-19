package main

import (
	"fmt"
	"math"
)

func main() {
	// Call the function

	var investedAmount, years float64 = 1000, 10 //multiple explicit type
	expectedReturnRate := 5.5                    //infered type
	var label, offset = "Investment", 10         //multiple inferred type
	var result = CalculateInvestment(investedAmount, expectedReturnRate, years)
	fmt.Println(offset, " The future value of the ", label, " is: ", result)

}
func CalculateInvestment(investedAmount float64, expectedReturnRate float64, years float64) float64 {
	// Calculate the investment
	return investedAmount *
		math.Pow(1+expectedReturnRate/100, years)

}
