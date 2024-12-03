package main

import "fmt"

// 7 6 4 2 1

func main() {
	arr := []int{7, 6, 4, 2, 1}
	temp := []int{1, 2}

	temp = append(temp, arr[:2]...)
	temp = append(temp, arr[2:]...)
	fmt.Println(temp)

}
