package main

import "testing"

// see https://habr.com/ru/company/mailru/blog/510200/
//
// $ sysctl -a | grep cacheline
// hw.cachelinesize: 64
//
// $ sysctl hw.l1dcachesize
// hw.l1dcachesize: 32768

var matrixLen = 512

func BenchmarkMatrixCombination(b *testing.B) {
	matrixA := createMatrix(matrixLen)
	matrixB := createMatrix(matrixLen)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < matrixLen; i++ {
			for j := 0; j < matrixLen; j++ {
				matrixA[i][j] = matrixA[i][j] + matrixB[i][j]
			}
		}
	}
}

func BenchmarkMatrixReversedCombination(b *testing.B) {
	matrixA := createMatrix(matrixLen)
	matrixB := createMatrix(matrixLen)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < matrixLen; i++ {
			for j := 0; j < matrixLen; j++ {
				matrixA[i][j] = matrixA[i][j] + matrixB[j][i]
			}
		}
	}
}

// loop nest optimization
func BenchmarkMatrixReversedCombinationPerBlock(b *testing.B) {
	matrixA := createMatrix(matrixLen)
	matrixB := createMatrix(matrixLen)
	blockSize := 8

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < matrixLen; i += blockSize {
			for j := 0; j < matrixLen; j += blockSize {
				for ii := i; ii < i+blockSize; ii++ {
					for jj := j; jj < j+blockSize; jj++ {
						matrixA[ii][jj] = matrixA[ii][jj] + matrixB[jj][ii]
					}
				}
			}
		}
	}
}

func createMatrix(len int) [][]int64 {
	res := make([][]int64, len)
	for i := 0; i < len; i++ {
		res[i] = make([]int64, len)
	}
	return res
}
