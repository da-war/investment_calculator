package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64
	//print in termial that you need to input the investment amount, years, and expected return rate
	fmt.Print("Enter the investment amount: ")
	fmt.Scanln(&investmentAmount)
	fmt.Print("Enter the number of years: ")
	fmt.Scanln(&years)
	fmt.Print("Enter the expected return rate: ")
	fmt.Scanln(&expectedReturnRate)

	doPow := math.Pow(1+expectedReturnRate/100, years)
	futureValue := investmentAmount * doPow

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future Value:", futureValue)
	fmt.Println("Future Real Value:", futureRealValue)
}
