package main

import (
	"fmt"
)

// Fungsi untuk menjumlahkan dua matriks
func addMatrices(A, B [][]int) [][]int {
	rows := len(A)
	cols := len(A[0])
	result := make([][]int, rows)

	for i := 0; i < rows; i++ {
		result[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			result[i][j] = A[i][j] + B[i][j]
		}
	}
	return result
}

// Fungsi untuk mengalikan dua matriks
func multiplyMatrices(A, B [][]int) [][]int {
	rowsA := len(A)
	colsA := len(A[0])
	colsB := len(B[0])
	result := make([][]int, rowsA)

	for i := 0; i < rowsA; i++ {
		result[i] = make([]int, colsB)
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return result
}

// Fungsi untuk mencetak matriks
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

// Fungsi utama
func main() {
	// Definisikan dua matriks
	A := [][]int{
		{1, 2},
		{3, 4},
	}
	B := [][]int{
		{5, 6},
		{7, 8},
	}
	C := [][]int{
		{9, 10},
		{11, 12},
	}

	// 1. Sifat Komutatif Penjumlahan
	fmt.Println("Sifat Komutatif Penjumlahan:")
	sumAB := addMatrices(A, B)
	sumBA := addMatrices(B, A)
	fmt.Println("A + B:")
	printMatrix(sumAB)
	fmt.Println("B + A:")
	printMatrix(sumBA)
	fmt.Println("Hasil sama:", sumAB[0][0] == sumBA[0][0], "dan", sumAB[1][1] == sumBA[1][1])

	// 2. Sifat Asosiatif Penjumlahan
	fmt.Println("\nSifat Asosiatif Penjumlahan:")
	sumABC1 := addMatrices(addMatrices(A, B), C)
	sumABC2 := addMatrices(A, addMatrices(B, C))
	fmt.Println("(A + B) + C:")
	printMatrix(sumABC1)
	fmt.Println("A + (B + C):")
	printMatrix(sumABC2)
	fmt.Println("Hasil sama:", sumABC1[0][0] == sumABC2[0][0], "dan", sumABC1[1][1] == sumABC2[1][1])

	// 3. Sifat Komutatif Perkalian
	fmt.Println("\nSifat Komutatif Perkalian:")
	productAB := multiplyMatrices(A, B)
	productBA := multiplyMatrices(B, A)
	fmt.Println("A x B:")
	printMatrix(productAB)
	fmt.Println("B x A:")
	printMatrix(productBA)
	fmt.Println("Hasil sama:", productAB[0][0] == productBA[0][0], "dan", productAB[1][1] == productBA[1][1])

	// 4. Sifat Asosiatif Perkalian
	fmt.Println("\nSifat Asosiatif Perkalian:")
	productABC1 := multiplyMatrices(multiplyMatrices(A, B), C)
	productABC2 := multiplyMatrices(A, multiplyMatrices(B, C))
	fmt.Println("(A x B) x C:")
	printMatrix(productABC1)
	fmt.Println("A x (B x C):")
	printMatrix(productABC2)
	fmt.Println("Hasil sama:", productABC1[0][0] == productABC2[0][0], "dan", productABC1[1][1] == productABC2[1][1])

	// 5. Sifat Distributif
	fmt.Println("\nSifat Distributif:")
	sumBC := addMatrices(B, C)
productADistributive := multiplyMatrices(A, sumBC)
sumProductADistributive := addMatrices(multiplyMatrices(A, B), multiplyMatrices(A, C))
	fmt.Println("A x (B + C):")
	printMatrix(productADistributive)
	fmt.Println("A x B + A x C:")
	printMatrix(sumProductADistributive)
	fmt.Println("Hasil sama:", productADistributive[0][0] == sumProductADistributive[0][0], "dan", productADistributive[1][1] == sumProductADistributive[1][1])
}
