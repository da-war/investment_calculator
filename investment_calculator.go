package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	doPow := math.Pow(1+expectedReturnRate/100, years)
	futureValue := investmentAmount * doPow

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future Value:", futureValue)
	fmt.Println("Future Real Value:", futureRealValue)
	fmt.Println("Working")
}
