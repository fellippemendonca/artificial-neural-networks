package perceptron

import (
	"fmt"
)

type Perceptron struct {
	weights      []float64
	learningRate float64
	iterations   int
}

func (p *Perceptron) Predict(inputs []float64) int {
	activation := p.weights[0]

	for i := 0; i < len(inputs); i++ {
		activation += p.weights[i] * inputs[i]
	}
	if activation > 0 {
		return 1
	}
	return 0
}

func (p *Perceptron) Fit(input [][]float64, target []int) {
	var prediction int
	var error int

	for it := 0; it < p.iterations; it++ {
		fmt.Println("Weights", p.weights)
		for i := 0; i < len(input); i++ {
			prediction = p.Predict(input[i])
			error = target[i] - prediction
			p.weights[0] += float64(error) * p.learningRate

			for j := 0; j < len(input[i]); j++ {
				p.weights[j] += p.learningRate * float64(error) * input[i][j]

			}
		}
	}
}

func New(size int, learningRate float64, iterations int) Perceptron {
	if learningRate == 0 {
		learningRate = 0.01
	}
	if iterations == 0 {
		iterations = 10
	}
	return Perceptron{
		weights:      make([]float64, size),
		learningRate: learningRate,
		iterations:   iterations,
	}
}
