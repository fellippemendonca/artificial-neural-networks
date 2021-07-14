package perceptron

import (
	"fmt"
)

type Perceptron struct {
	weights      []float64
	learningRate float64
}

func (p *Perceptron) Predict(input []float64) int {
	activation := p.weights[0]

	for i := 0; i < len(input); i++ {
		activation += p.weights[i] * input[i]
	}
	if activation > 0 {
		return 1
	}
	return 0
}

func (p *Perceptron) Training(input []float64, target int) {
	prediction := p.Predict(input)
	error := target - prediction
	for i := 0; i < len(input); i++ {
		p.weights[i] += p.learningRate * float64(error) * input[i]
	}
}

func Teaching(p *Perceptron, inputs [][]float64, target []int) {
	fmt.Println("Weights", p.weights)
	for i := 0; i < len(inputs); i++ {
		p.Training(inputs[i], target[i])
	}
}

func New(size int, learningRate float64) Perceptron {
	if learningRate == 0 {
		learningRate = 0.01
	}
	return Perceptron{
		weights:      make([]float64, size),
		learningRate: learningRate,
	}
}
