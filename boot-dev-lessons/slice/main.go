package main

import (
	"errors"
	"fmt"
)

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	str := [3]string{primary, secondary, tertiary}
	var costs [3]int
	sum := 0
	for i := 0; i < 3; i++ {
		sum += len(str[i])
		costs[i] = sum
	}
	return str, costs
}

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	switch plan {
	case planFree:
		return messages[:2], nil
	case planPro:
		return messages[:], nil
	default:
		return nil, errors.New("unsupported plan")
	}
}

// Slices hold references to an underlying array
func manipulate(slice []int, num int) {
	fmt.Println(len(slice), cap(slice))
	slice[1] = num
}

func getMessageCosts(messages []string) []float64 {
	length := len(messages)                    // number of elements it contains
	capacity := cap(messages)                  // reports the maximum length the slice may assume. The number of elements in the underlying array, counting from the first element in the slice
	costs := make([]float64, length, capacity) // the capacity argument is usually omitted and defaults to the length
	for i := 0; i < length; i++ {
		costs[i] = float64(len(messages[i])) * 0.01
	}
	return costs
}

// ---------------- How Append method implemented --------------

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
	//pri := "Hello sir/madam can I interest you in a yacht?"
	//sec := "Please I'll even give you an Amazon gift card?"
	//ter := "You're missing out big time"
	//
	//strArr, costArr := getMessageWithRetries(pri, sec, ter)
	//
	//fmt.Println(strArr, costArr)

	//slice, err := getMessageWithRetriesForPlan("", [3]string{
	//	"Hello sir/madam can I interest you in a yacht?",
	//	"Please I'll even give you an Amazon gift card?",
	//	"You're missing out big time",
	//})
	//
	//fmt.Println(slice, err)

	//ar := [5]int{1, 2, 3, 4, 5}
	//fmt.Println(ar)
	//manipulate(ar[2:4], 10)
	//fmt.Println(ar)

	fmt.Println(getMessageCosts([]string{"asdfasdfa", "asdfawefwaxcvasd", "fashcoa hiosrfuwyifa", "asdfedcv"}))
}
