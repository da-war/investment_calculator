package main

import (
	"fmt"
)

// func main() {
// 	const inflationRate = 2.5
// 	var investmentAmount, years, expectedReturnRate float64
// 	//print in termial that you need to input the investment amount, years, and expected return rate
// 	fmt.Print("Enter the investment amount: ")
// 	fmt.Scan(&investmentAmount)
// 	fmt.Print("Enter the number of years: ")
// 	fmt.Scan(&years)
// 	fmt.Print("Enter the expected return rate: ")
// 	fmt.Scan(&expectedReturnRate)

// 	doPow := math.Pow(1+expectedReturnRate/100, years)
// 	futureValue := investmentAmount * doPow

// 	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
// 	fmt.Println("Future Value:", futureValue)
// 	fmt.Println("Future Real Value:", futureRealValue)
// }

func main() {
	var a string = "42"
	var b int = 100
	var x float32 = 3.1415
	fmt.Printf("the value of a is %s and b is %d and x is %f\n", a, b, x)
}
