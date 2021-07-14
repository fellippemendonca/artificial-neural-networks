package main

import (
	"fmt"

	perceptron "github.com/fellippemendonca/artificial-neural-networks/perceptron"
)

var plan = [][]float64{
	{0.4, 0.3},
	{0.3, 0.2},
	{0.2, 0.1},
	{0.1, 0.0},
	{0.0, 0.0},
	{0.0, 0.1},
	{0.1, 0.2},
	{0.2, 0.3},
	{0.3, 0.4},
	{0.4, 0.5},
}

var truths = []int{1, 1, 1, 1, 1, 0, 0, 0, 0, 0}

func checkResult(value int) {
	if value == 1 {
		fmt.Println("X >= Y")
	} else {
		fmt.Println("X < Y")
	}
}

func main() {

	x := 0.0
	y := 0.0

	p := perceptron.New(len(plan[0]), 0.05)
	for i := 0; i < 5; i++ {
		perceptron.Teaching(&p, plan, truths)
	}
	result2 := p.Predict([]float64{x, y})
	checkResult(result2)
}
