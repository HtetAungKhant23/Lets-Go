package main

import "fmt"

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

func main() {
	pri := "Hello sir/madam can I interest you in a yacht?"
	sec := "Please I'll even give you an Amazon gift card?"
	ter := "You're missing out big time"

	strArr, costArr := getMessageWithRetries(pri, sec, ter)

	fmt.Println(strArr, costArr)
}
