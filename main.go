package main

import "fmt"

func main() {
	a := []int{1000000000, 900000000, 900000000, 800000000, 700000000, 800000000, 1000000000}
	loner := Solution(a)
	fmt.Println(loner)

}

// Мне стыдно...Но пусть будет так
func Solution(inputArray []int) int {
	var result int
	for i := 0; i < len(inputArray); i++ {

		if i == len(inputArray)-1 {
			return inputArray[i]
		}

		if inputArray[i] == 0 {
			continue
		}

		isFind := false
		temp := inputArray[i]
		for j := i + 1; j < len(inputArray); j++ {
			if inputArray[j] == temp {
				isFind = true
				inputArray[j] = 0
			}
		}
		if !isFind {
			result = inputArray[i]
			break
		}
	}
	return result
}
