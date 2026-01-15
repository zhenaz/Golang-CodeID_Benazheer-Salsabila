package main

import "fmt"

func main() {
	// Quiz Time (4) 6
	/* var matrix [5][5]int
	matrixCounter := matrixRowTriangle1(matrix)
	displayMatrix(matrixCounter) */

	// Quiz Time (4) 6
	/* var matrix [5][5]int
	matrixCounter := matrixRowTriangle2(matrix)
	displayMatrix(matrixCounter) */

	// Quiz Time (5) 8
	/* var matrix [7][7]int
	matrixCounter := matrixHollowCounter(matrix)
	displayMatrix7(matrixCounter) */

	// Quiz Time (5) 9
	/* var matrix [7][7]int
	matrixCounter := matrixSumCounter(matrix)
	displayMatrix8(matrixCounter) */

	// Quiz Time (7)
	countCorrectAnswers()

}

// display matrix 5x5 (soal 4)
func displayMatrix(matrix [5][5]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			fmt.Printf("%5d", matrix[row][col])
		}
		fmt.Println()
	}

}

// display matrix 7x7 (soal 5)
func displayMatrix7(matrix [7][7]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			fmt.Printf("%5d", matrix[row][col])
		}
		fmt.Println()
	}

}

// display matrix 8 (soal 5.2)
func displayMatrix8(matrix [8][8]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			fmt.Printf("%5d", matrix[row][col])
		}
		fmt.Println()
	}

}

// Quiz Time (4) 6
func matrixRowTriangle1(matrix [5][5]int) [5][5]int {
	var counterMatrix int = 1
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if col == row {
				matrix[row][col] = counterMatrix
				counterMatrix++
			} else if col >= row {
				matrix[row][col] = 10
			} else {
				matrix[row][col] = 20
			}

		}
	}
	return matrix
}

// Quiz Time (4) 7

func matrixRowTriangle2(matrix [5][5]int) [5][5]int {
	var counterMatrix int = len(matrix)
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if col == row {
				matrix[row][col] = counterMatrix
				counterMatrix--
			} else if col >= row {
				matrix[row][col] = 20
			} else {
				matrix[row][col] = 10
			}

		}
	}
	return matrix
}

// Quiz Time (5) 8
func matrixHollowCounter(matrix [7][7]int) [7][7]int {
	var counterMatrixRow int = 1
	var counterMatrixCol int = 1
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if col == 0 && row == 0 {
				matrix[row][col] = 0
			} else if col == 0 || col == len(matrix) || row == len(matrix)-1 {
				matrix[row][col] = counterMatrixCol
				counterMatrixCol++
			} else if row == 0 || row == len(matrix) || col == len(matrix)-1 {
				matrix[row][col] = counterMatrixRow
				counterMatrixRow++
			} else {
				matrix[row][col] = 0
			}
		}
	}
	return matrix
}

func matrixSumCounter(matrix [7][7]int) [8][8]int {
	var result [8][8]int

	inputSize := len(matrix)
	outputSize := len(result)
	diagonalSum := 0
	lastIndex := outputSize - 1

	for row := 0; row < inputSize; row++ {
		rowSum := 0
		for col := 0; col < inputSize; col++ {
			result[row][col] = row + col

			rowSum += row + col
			result[lastIndex][col] += row + col

			if row == col {
				diagonalSum += row + col
			}
		}
		result[row][lastIndex] = rowSum
	}

	total := 0
	for i := 0; i < inputSize; i++ {
		total += result[lastIndex][i]
	}
	result[lastIndex][lastIndex] = diagonalSum

	return result
}

func countCorrectAnswers() {
	answers := [8][10]string{
		{"A", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"D", "B", "A", "B", "C", "A", "E", "E", "A", "D"},
		{"E", "D", "D", "A", "C", "B", "E", "E", "A", "D"},
		{"C", "B", "A", "E", "D", "C", "E", "E", "A", "D"},
		{"A", "B", "D", "C", "C", "D", "E", "E", "A", "D"},
		{"B", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
		{"B", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"E", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
	}
	kunciJawaban := [10]string{"D", "B", "D", "C", "C", "D", "A", "E", "A", "D"}

	// hitung jawaban benar tiap siswa
	for i := 0; i < len(answers); i++ {
		benar := 0
		for j := 0; j < len(kunciJawaban); j++ {
			if answers[i][j] == kunciJawaban[j] {
				benar++
			}
		}
		fmt.Printf("Jawaban Siswa %d yang benar : %d\n", i, benar)
	}
}
