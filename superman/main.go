package main

import "fmt"

func superman(chickens []int, k int) int {
	if k <= 0 {
		fmt.Println("Hi Clark Kent, bring the roof before rescue the chicken please")
		return 0
	}

	amount := len(chickens)
	if amount == 0 {
		fmt.Println("invalid chickens input")
		return 0
	}
	start := 0
	end := 0
	save := 0

	for end < amount {
		for end < amount && chickens[end]-chickens[start]+1 <= k {
			// move pointer
			end++
		}
		// compare new max
		newSave := end - start
		if newSave > save {
			save = newSave
		}
		// reset start and end pointer
		newStart := start + 1
		start = newStart
		end = newStart
	}
	return save
}

func main() {
	output1 := superman([]int{2, 5, 10, 12, 15}, 5)
	output2 := superman([]int{1, 11, 30, 34, 35, 37}, 10)
	output3 := superman([]int{1, 2, 5, 31, 32, 33}, 3)
	fmt.Println(output1)
	fmt.Println(output2)
	fmt.Println(output3)
}
