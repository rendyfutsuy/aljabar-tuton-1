package main

import (
	"fmt"
)

// Fungsi untuk membuat matriks dengan ukuran tertentu
func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	return matrix
}

// Fungsi untuk menjumlahkan dua matriks
func addMatrices(A, B [][]int) [][]int {
	rows, cols := len(A), len(A[0])
	C := createMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}
	return C
}

// Fungsi untuk mengalikan dua matriks
func multiplyMatrices(A, B [][]int) [][]int {
	rowsA, colsA := len(A), len(A[0])
	rowsB, colsB := len(B), len(B[0])
	if colsA != rowsB {
		fmt.Println("Ukuran matriks tidak valid untuk perkalian!")
		return nil
	}
	C := createMatrix(rowsA, colsB)
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}

// Fungsi untuk membuktikan sifat komutatif pada penjumlahan matriks
func testCommutativeAddition(A, B [][]int) {
	fmt.Println("Menguji sifat komutatif penjumlahan: A + B = B + A")
	sumAB := addMatrices(A, B)
	sumBA := addMatrices(B, A)

	if compareMatrices(sumAB, sumBA) {
		fmt.Println("Penjumlahan matriks bersifat komutatif")
	} else {
		fmt.Println("Penjumlahan matriks tidak bersifat komutatif")
	}
}

// Fungsi untuk membuktikan sifat komutatif pada perkalian matriks
func testCommutativeMultiplication(A, B [][]int) {
	fmt.Println("Menguji sifat komutatif perkalian: A * B = B * A")
	mulAB := multiplyMatrices(A, B)
	mulBA := multiplyMatrices(B, A)

	if compareMatrices(mulAB, mulBA) {
		fmt.Println("Perkalian matriks bersifat komutatif")
	} else {
		fmt.Println("Perkalian matriks tidak bersifat komutatif")
	}
}

// Fungsi untuk membuktikan sifat distributif
func testDistributive(A, B, C [][]int) {
	fmt.Println("Menguji sifat distributif: A * (B + C) = A * B + A * C")
	sumBC := addMatrices(B, C)
	mulA_sumBC := multiplyMatrices(A, sumBC)

	mulAB := multiplyMatrices(A, B)
	mulAC := multiplyMatrices(A, C)
	mulA_sumAB_AC := addMatrices(mulAB, mulAC)

	if compareMatrices(mulA_sumBC, mulA_sumAB_AC) {
		fmt.Println("Perkalian matriks bersifat distributif")
	} else {
		fmt.Println("Perkalian matriks tidak bersifat distributif")
	}
}

// Fungsi untuk membandingkan dua matriks
func compareMatrices(A, B [][]int) bool {
	for i := range A {
		for j := range A[i] {
			if A[i][j] != B[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	// Contoh matriks A, B, dan C
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

	// Pengujian sifat komutatif pada penjumlahan
	testCommutativeAddition(A, B)

	// Pengujian sifat komutatif pada perkalian
	testCommutativeMultiplication(A, B)

	// Pengujian sifat distributif
	testDistributive(A, B, C)
}
