package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// Function to perform Gaussian elimination
func gaussianElimination(matrix [][]float64, n int) {
	var wg sync.WaitGroup

	// Forward elimination
	for i := 0; i < n; i++ {
		// Pivoting
		maxRow := i
		for k := i + 1; k < n; k++ {
			if math.Abs(matrix[k][i]) > math.Abs(matrix[maxRow][i]) {
				maxRow = k
			}
		}
		matrix[i], matrix[maxRow] = matrix[maxRow], matrix[i]

		// Elimination
		for j := i + 1; j < n; j++ {
			wg.Add(1)
			go func(j, i int) {
				defer wg.Done()
				factor := matrix[j][i] / matrix[i][i]
				for k := i; k < n+1; k++ {
					matrix[j][k] -= factor * matrix[i][k]
				}
			}(j, i)
		}
		wg.Wait()
	}
}

// Function to back substitute to find the solution
func backSubstitution(matrix [][]float64, n int) []float64 {
	solution := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		solution[i] = matrix[i][n] / matrix[i][i]
		for j := i + 1; j < n; j++ {
			solution[i] -= matrix[i][j] * solution[j] / matrix[i][i]
		}
	}
	return solution
}

func main() {
	var n int

	// User input for matrix dimension
	fmt.Print("Enter the number of variables (dimension of the matrix): ")
	fmt.Scan(&n)

	// Create an augmented matrix
	matrix := make([][]float64, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]float64, n+1) // n variables + 1 for the constants
	}

	// User input for matrix elements
	fmt.Println("Enter the augmented matrix (each row should have n+1 elements):")
	for i := 0; i < n; i++ {
		for j := 0; j < n+1; j++ {
			fmt.Printf("Element [%d][%d]: ", i, j)
			fmt.Scan(&matrix[i][j])
		}
	}

	// Start timing the Gaussian elimination process
	start := time.Now()

	gaussianElimination(matrix, n)
	solution := backSubstitution(matrix, n)

	// Calculate elapsed time
	elapsed := time.Since(start)

	// Output the solution
	fmt.Println("Solution:")
	for i, val := range solution {
		fmt.Printf("x[%d] = %f\n", i, val)
	}

	// Output the time taken
	fmt.Printf("Time taken to solve the equations: %s\n", elapsed)
}
