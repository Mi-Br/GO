package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	moods := [...]string{
		0: "great 👍",
		1: "bit sad 😭",
		2: "very angry 💥",
		3: "focused and determined 👀",
	}
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(moods))
	fmt.Printf("\nSacrates feels %s today!", moods[r])
}
