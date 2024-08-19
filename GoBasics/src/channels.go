package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Backend programming management system
const PROBLEMS_SOLVED_TO_WIN float32 = 5

var PARTICIPANTS = []string{"L", "A", "B", "Z", "X", "P"}

func channels() {
	var problemsSolvedChannel = make(chan string, len(PARTICIPANTS))
	for i := range PARTICIPANTS {
		wg.Add(1)
		go getScore(PARTICIPANTS[i], problemsSolvedChannel)
	}
	wg.Wait()
	close(problemsSolvedChannel)
	sendMessage(problemsSolvedChannel)
}

// Each participant gets 3 tries to score over 5 out of 6
func getScore(participant string, problemsSolvedChannel chan string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		var problemsSolved = rand.Float32() * 10
		if problemsSolved > PROBLEMS_SOLVED_TO_WIN {
			problemsSolvedChannel <- participant
			wg.Done()
			return
		}
	}
	wg.Done()
}

func sendMessage(channel chan string) {
	fmt.Println("And the winners are...")
	for p := range channel {
		fmt.Println(p)
	}
}
