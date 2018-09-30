package main

import (
	"fmt"
	"math"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

// NeuralNetwork class
type NeuralNetwork struct {
	sizes      []int
	biases     []*mat.Dense
	weights    []*mat.Dense
	activation string
}

// NewNN : Constructor
func NewNN(sizes []int) NeuralNetwork {
	// Make biases slice
	b := make([]*mat.Dense, len(sizes[1:]))
	for i, y := range sizes[1:] {
		b[i] = Randmat(y, 1)
	}

	// Make weights
	w := make([]*mat.Dense, len(sizes[1:]))
	X, Y := sizes[:len(sizes)-1], sizes[1:]
	for i := range X {
		x, y := X[i], Y[i]
		w[i] = Randmat(y, x)
	}

	return NeuralNetwork{sizes, b, w, "sigmoid"}
}

// Show : displays contents
func (nn *NeuralNetwork) Show() {
	fmt.Println("Sizes: ", nn.sizes)
	fmt.Println("Biases :", "\n")
	for _, v := range nn.biases {
		fmt.Println(mat.Formatted(v), "\n")
	}
	fmt.Println("Weights :", "\n")
	for _, v := range nn.weights {
		fmt.Println(mat.Formatted(v), "\n")
	}
	//fmt.Println(mat.Formatted(Randmat(2, 2)))

}

// Predict :Predicts result based on input
func (nn *NeuralNetwork) Predict(inpLocal *mat.Dense) *mat.Dense {
	var c *mat.Dense
	for i := range nn.biases {
		// fmt.Println("Loop run", i+1)
		// fmt.Println(nn.weights[i].Dims())
		// fmt.Println(inpLocal.Dims())
		r1, _ := nn.weights[i].Dims()
		_, c2 := inpLocal.Dims()
		c = Randmat(r1, c2) // temp variable resize
		c.Mul(nn.weights[i], inpLocal)
		//fmt.Println(c.Dims())

		// Redefine inpLocal for dimension change
		row, col := c.Dims()
		inpLocal = Randmat(row, col)
		inpLocal.Add(c, nn.biases[i])
		//fmt.Println(inpLocal.Dims())
		activation(nn.activation, inpLocal)
		//fmt.Println(inpLocal.Dims())
	}

	return inpLocal
}

// Randmat matrix generator
func Randmat(r, c int) *mat.Dense {
	temp := make([]float64, r*c)
	for i := range temp {
		temp[i] = rand.NormFloat64() // between 0 and 1
	}
	return mat.NewDense(r, c, temp)
}

// Apply activation function
func activation(ac string, inp *mat.Dense) {
	if ac == "sigmoid" {
		r, c := inp.Dims()
		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				val0 := inp.At(i, j)
				value := 1 / (1 + math.Exp(val0))
				inp.Set(i, j, value)
			}
		}
	}

}
