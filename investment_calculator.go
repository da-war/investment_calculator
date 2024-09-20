package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	doPow := math.Pow(1+expectedReturnRate/100, years)
	futureValue := investmentAmount * doPow
	fmt.Println(futureValue)
}
