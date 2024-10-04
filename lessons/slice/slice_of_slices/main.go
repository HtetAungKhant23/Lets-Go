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

// ---- append method ------- don't use append in other slice -------

func Append(slice, data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) { // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data))*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	return slice
}

func main() {
	//res := createMatrix(0, 0)
	//fmt.Println(res)

	//------ track how is append method work ---------
	s := make([]byte, 3, 8)
	s[0] = 1
	s[1] = 2

	fmt.Println(len(s), cap(s), s, &s[0])
	data := []byte{3, 4, 5}

	l := len(s)
	if l+len(data) > cap(s) {
		newSlice := make([]byte, (l+len(data))*2)
		fmt.Println(len(newSlice), cap(newSlice), &newSlice[0])
		copy(newSlice, s)
		fmt.Println("new slice ", newSlice, &newSlice[0], &s[0])
		s = newSlice
		fmt.Println("s is ", s, &s[0])

	}
	fmt.Println("ok", cap(s))
	s = s[0 : l+len(data)]
	fmt.Println(len(s))
	copy(s[l:], data)
	fmt.Println(s, &s[0])

}

//  ---------- note ----------

//  var tt []int              => equal nil
//  tt := make([]int, 0)      => not equal nil
