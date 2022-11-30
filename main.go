package main

import (
	"fmt"
	"math/rand"
	"time"
)

const rowsm1 = 3
const rowsm2 = 5
const colm2 = 4

var matrix35 [rowsm1][rowsm2]int
var matrix54 [rowsm2][colm2]int

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(matrix35); i++ {
		for j := 0; j < len(matrix35[i]); j++ {
			matrix35[i][j] = rand.Intn(5)
		}
	}

	for i := 0; i < len(matrix54); i++ {
		for j := 0; j < len(matrix54[i]); j++ {
			matrix54[i][j] = rand.Intn(3)
		}
	}
	fmt.Println("Входная матрица", rowsm1, "*", rowsm2)
	for i := 0; i < len(matrix35); i++ {
		fmt.Println(matrix35[i])
	}
	fmt.Println("Входная матрица ", rowsm2, "*", colm2)
	for i := 0; i < len(matrix54); i++ {
		fmt.Println(matrix54[i])
	}
	fmt.Println("Результат сложения")
	final := miltiplymatrix(matrix35, matrix54)
	for i := 0; i < len(final); i++ {
		fmt.Println(final[i])
	}

}

func miltiplymatrix(a [rowsm1][rowsm2]int, b [rowsm2][colm2]int) [rowsm1][colm2]int {
	var matrix34 [rowsm1][colm2]int
	for i := 0; i < rowsm1; i++ {
		for j := 0; j < colm2; j++ {
			for k := 0; k < rowsm2; k++ {
				matrix34[i][j] = matrix34[i][j] + a[i][k]*b[k][j]
			}
		}
	}
	return matrix34
}
