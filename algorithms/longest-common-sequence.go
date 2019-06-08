package main

import "fmt"

func main() {
	result := findLongestCommonSequenceSize("fish", "fosh")
	fmt.Println(result)

	result = findLongestCommonSubstringSize("fish", "fosh")
	fmt.Println(result)
}

// Если переменные одного типа в параметрах, то можно указать тип единожды
func findLongestCommonSequenceSize(a, b string) int {
	matrix := createMatrix(len(a), len(b))

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if i == 0 || j == 0 {
				matrix[i][j] = 1
				continue
			}

			if a[i] == b[j] {
				matrix[i][j] = matrix[i - 1][j - 1] + 1
			} else {
				matrix[i][j] = computeMax(matrix[i - 1][j], matrix[i][j - 1]);
			}
		}
	}

	return matrix[len(a) - 1][len(b) - 1]
}

func findLongestCommonSubstringSize(a, b string) int {
	matrix := createMatrix(len(a), len(b))
	lastRowIndex := 0
	lastColumnIndex := 0

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				if i == 0 || j == 0 {
					matrix[i][j] = 1
					continue
				}

				matrix[i][j] = matrix[i - 1][j - 1] + 1

				lastRowIndex = i
				lastColumnIndex = j
			} else {
				matrix[i][j] = 0
			}
		}
	}

	return matrix[lastRowIndex][lastColumnIndex]
}

func createMatrix(rows, columns int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < columns; i++ {
		matrix[i] = make([]int, columns)
	}

	return matrix
}

func computeMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}