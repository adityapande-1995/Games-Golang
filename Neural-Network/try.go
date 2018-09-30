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
