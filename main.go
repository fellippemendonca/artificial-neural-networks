package main

import (
	"fmt"

	perceptron "github.com/fellippemendonca/artificial-neural-networks/perceptron"
)

var x = [][]float64{
	{180.0, 80.0},
	{175.0, 67.0},
	{100.0, 30.0},
	{120.0, 32.0},
}

var y = []int{0, 0, 1, 1}

// 0 - adult
// 1 - child

func main() {
	p := perceptron.New(len(x[0]), 1.0, 10)
	p.Fit(x, y)

	result := p.Predict([]float64{150.0, 50.0})
	if result == 0 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Child")
	}

}
