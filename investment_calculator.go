package main

import (
	"fmt"
	"math"
)

func main() {

	var welcome string = stringGetter()
	fmt.Println(welcome)
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64
	//print in termial that you need to input the investment amount, years, and expected return rate
	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)
	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// doPow := math.Pow(1+expectedReturnRate/100, years)
	// futureValue := investmentAmount * doPow

	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	// fmt.Printf("The Future Value: %.2f\nThe Future Actual Value %.3f\n", futureValue, futureRealValue)

	fv, frv := calculateFutureValue(investmentAmount, years, expectedReturnRate, inflationRate)

	fmt.Printf("The Future Value: %.2f\nThe Future Actual Value %.3f\n", fv, frv)
}

func stringGetter() string {
	return "Welcome to the Investment Calculator"
}

func calculateFutureValue(investmentAmount, years, expectedReturnRate, inflationRate float64) (float64, float64) {
	doPow := math.Pow(1+expectedReturnRate/100, years)
	futureValue := investmentAmount * doPow
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	return futureValue, futureRealValue
}
