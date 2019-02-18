package main

const N = 10

type Point struct{
	Weight int
	Final bool
}


var matrix [N][N]int

func main() {
	initializeMatrix()
}

func initializeMatrix() {
	for i := 0; i < N; i++{
		for j := i; j < N; j++{
			if j == i {
				matrix[i][j] = 0
				continue
			}
			randNumber := 0
			matrix[j][i] = randNumber
			matrix[i][j] = randNumber
		}
	}
}