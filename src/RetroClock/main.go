package main

import (
	"fmt"
	"time"
)

func main() {

	// type board struct {
	// 	digit placeholder
	// 	order int
	// }
	// var panel [8]board

	var index = [8]int{6, 7, 8, 9, 10, 11, 12, 13}
	for i, _ := range index {
		index[i] = i
	}

	blink := true

	for {
		now := time.Now()
		time.Sleep(1 * time.Second)
		fmt.Print("\033[H\033[2J")
		hours, minutes, seconds := now.Hour(), now.Minute(), now.Second()
		blink = !blink

		// get latest time and put to sockets
		board := [...]placeholder{
			digits[hours/10],
			digits[hours%10],
			colon,
			digits[minutes/10],
			digits[minutes%10],
			colon,
			digits[seconds/10],
			digits[seconds%10],
		}

		shift := [...]placeholder{
			empty,
			empty,
			empty,
			empty,
			empty,
			empty,
			empty,
			empty,
		}

		// calculate digit placement based on index
		for i, v := range index {
			switch {
			case v >= 0 && v < 8:
				shift[v] = board[i]
			}
		}

		// shift index by - 1
		for i, v := range index {
			index[i] = v - 1
			// if index[i] < -1*(len(index)-1) {
			// 	index[i] = len(index) - 1
			// }
		}

		clock := [...]placeholder{
			shift[0],
			shift[1],
			shift[2],
			shift[3],
			shift[4],
			shift[5],
			shift[6],
			shift[7],
		}

		for l := range clock[0] {
			for d := range clock {
				fmt.Print(clock[d][l], "  ")
			}
			fmt.Println()
		}

		//ofset index to start from the right
		if index[0] == -8 {
			index = [8]int{7, 8, 9, 10, 11, 12, 13, 14}
		}

	}
}
