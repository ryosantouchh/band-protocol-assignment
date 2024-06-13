package main

import "fmt"

func isBossBabyAGoodBoy(shoots string) string {
	if shoots == "" {
		return "That kid doesn's shoot the boss baby"
	}
	shot := 0
	revenge := 0
	for _, shoot := range shoots {
		if revenge > shot {
			return "Bad Boy"
		}
		switch shoot {
		case 'R':
			revenge++
			if shot == 0 || shot < revenge {
				return "Bad boy"
			}
		case 'S':
			if shot > 0 && revenge > 0 && shot >= revenge {
				return "Good Boy"
			}
			shot++
		}
	}
	return "Good Boy"
}

func main() {
	output1 := isBossBabyAGoodBoy("SRSSRRR")
	output2 := isBossBabyAGoodBoy("RSSRR")
	output3 := isBossBabyAGoodBoy("SSSRRRRS")
	output4 := isBossBabyAGoodBoy("SRRSSR")
	output5 := isBossBabyAGoodBoy("SSRSRRR")
	fmt.Println(output1)
	fmt.Println(output2)
	fmt.Println(output3)
	fmt.Println(output4)
	fmt.Println(output5)
}
