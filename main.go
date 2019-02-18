package main

import (
	"fmt"
	"math/rand"
)

const N = 5

type Point struct{
	Weight int
	Final bool
}

var Points [N]Point
var matrix [N][N]int
var way []int

func main() {
	initializeMatrix()

	from := 0
	to := 5
	way = countWay(from, to)

	showResults()
}

func countWay(from, to int) []int {
	way := make([]int, 1)

	for i := 0; i < N; i++ {

	}

	return way
}

func initializeMatrix() {
	for i := 0; i < N; i++{
		for j := i; j < N; j++{
			if j == i {
				matrix[i][j] = 0
				continue
			}
			randNumber := rand.Int() % 10
			matrix[j][i] = randNumber
			matrix[i][j] = randNumber
		}
		Points[i].Final = false
		Points[i].Weight = 100000
	}
}

func showResults() {
	for i := 0; i < N; i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println(Points)
}