package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}



	for index, o := range arr{
		fmt.Print(index)
		if o == 9 {
			arr = append(arr[:index], arr[index +1:]...)
			fmt.Print(index)
			break
		}
	}

	fmt.Print(arr)
}
