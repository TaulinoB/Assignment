package main

import (
	"fmt"
	"sync"
	"time"
)

// Function to calculate the determinant of a matrix
func determinant(matrix [][]float64, n int) float64 {
	if n == 1 {
		return matrix[0][0]
	}
	if n == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}

	det := 0.0
	for c := 0; c < n; c++ {
		subMatrix := make([][]float64, n-1)
		for i := range subMatrix {
			subMatrix[i] = make([]float64, n-1)
		}
		for i := 1; i < n; i++ {
			for j := 0; j < n; j++ {
				if j < c {
					subMatrix[i-1][j] = matrix[i][j]
				} else if j > c {
					subMatrix[i-1][j-1] = matrix[i][j]
				}
			}
		}
		sign := 1.0
		if c%2 == 1 {
			sign = -1.0
		}
		det += sign * matrix[0][c] * determinant(subMatrix, n-1)
	}
	return det
}

// Function to replace a column in the matrix
func replaceColumn(matrix [][]float64, column int, vector []float64) [][]float64 {
	newMatrix := make([][]float64, len(matrix))
	for i := range newMatrix {
		newMatrix[i] = make([]float64, len(matrix[i]))
		copy(newMatrix[i], matrix[i])
		newMatrix[i][column] = vector[i]
	}
	return newMatrix
}

// Function to solve the system using Cramer's Rule
func cramer(matrix [][]float64, b []float64, n int) []float64 {
	var wg sync.WaitGroup
	detA := determinant(matrix, n)
	solutions := make([]float64, n)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			detAi := determinant(replaceColumn(matrix, i, b), n)
			solutions[i] = detAi / detA
		}(i)
	}
	wg.Wait()
	return solutions
}

func main() {
	var n int

	// User input for matrix dimension
	fmt.Print("Enter the number of variables (dimension of the matrix): ")
	fmt.Scan(&n)

	// Create a matrix and vector for the equations
	matrix := make([][]float64, n)
	b := make([]float64, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]float64, n)
		fmt.Printf("Enter coefficients for equation %d (space-separated): ", i+1)
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
		fmt.Printf("Enter the constant for equation %d: ", i+1)
		fmt.Scan(&b[i])
	}

	// Start timing the Cramer's Rule process
	start := time.Now()

	// Solve the system using Cramer's Rule
	solutions := cramer(matrix, b, n)

	// Calculate elapsed time
	elapsed := time.Since(start)

	// Output the solution
	fmt.Println("Solution:")
	for i, val := range solutions {
		fmt.Printf("x[%d] = %f\n", i, val)
	}

	// Output the time taken
	fmt.Printf("Time taken to solve the equations: %s\n", elapsed)
}
