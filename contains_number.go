//Program to check if the list of integers are repeated. Has duplicated

package main

import (
	"fmt"
)

func main() {

	listOfInput := []int{1, 5, 6, 8, 2, 6, 8, 42, 69, 8, 5, 9, 56, 6, 54, 4, 69, 9, 6, 54, 5, 5, 69, 6, 65, 4, 4, 5, 5, 9, 6, 66, 7}

	res := findDuplicates(listOfInput)

	fmt.Println(res)

}

func findDuplicates(num []int) bool {
	s := map[int]bool{}

	for i := 0; i < len(num); i++ {

		s[num[i]] = true

	}
	return len(s) != len(num) == true
}
