package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		var temp []int
		for j := 0; j < cols; j++ {
			temp = append(temp, i*j)
		}
		matrix = append(matrix, temp)
	}
	return matrix
}

func main() {
	res := createMatrix(0, 0)
	fmt.Println(res)
}

//  ---------- note ----------

//  var tt []int              => equal nil
//  tt := make([]int, 0)      => not equal nil
