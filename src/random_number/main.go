package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const numberOfTries = 5
const specialMsg = `Bonus message !!! congrats on winning form the 1st time`
const ussage = `
			Welcome to the lucky guess game 👀
The program willl pickup %d random numbers. 
Your mission is to guess on those numbers.
		
The greater number is - the harder it gets

Wanna play ?
`

func main() {
	// get args
	if len(os.Args) < 2 {
		fmt.Println("please provide 1 number to guesss")
		return
	} else if len(os.Args) == 1 {
		fmt.Printf(ussage, numberOfTries)
	}

	// validate
	guess, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("please provide valid number: ", err)
		return
	}
	if guess <= 0 {
		fmt.Println("Please provide positive number")
		return
	}
	//initiate random seed using time's miliseconds
	t := time.Now().UnixNano()
	rand.Seed(t)

	var winMsg []string = []string{"Wow. You nailed it!", "Great job", "Winning streak !"}
	var lostMsg []string = []string{"Sorry you lost. Try again", "Next day- new hope", "Not lucky, sorry"}
	var i int
	//run loop and see if guessed
	for turn := 0; turn < numberOfTries; turn++ {
		n := rand.Intn(guess + 1)
		// fmt.Println(turn, " ", n)
		if n == guess {
			if turn == 0 {
				fmt.Println(specialMsg)
			} else {
				i = rand.Intn(len(lostMsg))
				fmt.Println(winMsg[i], guess)
			}
			return
		}
	}
	i = rand.Intn(len(lostMsg))
	fmt.Println(lostMsg[i])
	//report game status
}
