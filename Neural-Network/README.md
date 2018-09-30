# Neural networks in golang

## Introduction
This module implements simple neural network models from scratch in golang. Currently simple MLP model has been done.

## Dependencies
* Gonum : Matrix manipulation library

## Usage
For simple MLP model, run:
```
package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	layers := []int{10, 5, 5, 5}
	a := NewNN(layers)
	//a.Show()
	answer := a.Predict(Randmat(10, 1))
	fmt.Println(mat.Formatted(answer))

}
```
