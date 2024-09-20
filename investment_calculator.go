package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10
	var doPow = math.Pow(1+float64(expectedReturnRate)/100, float64(years))
	var futureValue = float64(investmentAmount) * doPow
	fmt.Println(futureValue)
}
