package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var guessNum int
var feedback string

func main() {
	fmt.Println("Guess Number Project")
	max := 9999
	min := 0
	rand.Seed(time.Now().UnixNano())

	secretNum := min + rand.Intn(max-min+1)

	// Create map for storing digits and position of digits for guess number and secret number
	// Convert numbers to string for indexing
	secretNumMap := make(map[string][]int)
	secretNumStr := strconv.Itoa(secretNum)

	for i := 0; i < len(secretNumStr); i++ {
		secretNumMap[secretNumStr[i:i+1]] = append(secretNumMap[secretNumStr[i:i+1]], i)
	}
	feedbackMap := make(map[int]string)

	for true {
		// Input your guess number
		//fmt.Printf("Secret number: %v\n", secretNum)
		fmt.Println("Guess your number:\t")
		fmt.Scan(&guessNum)

		// Input zero to exit
		if guessNum == 0 {
			break
		}

		guessNumMap := make(map[string][]int)
		guessNumStr := strconv.Itoa(guessNum)

		for i := 0; i < len(guessNumStr); i++ {
			guessNumMap[guessNumStr[i:i+1]] = append(guessNumMap[guessNumStr[i:i+1]], i)
		}
		position := 0
		contains := 0
		2
		for keyGuess := range guessNumMap {

			comparisonMap := map[int]int{}

			// Find intersection and difference of index of letters of guess word
			intersections := Intersection(guessNumMap[keyGuess], secretNumMap[keyGuess])
			diff := Difference(guessNumMap[keyGuess], secretNumMap[keyGuess])
			diff2 := Difference(secretNumMap[keyGuess], guessNumMap[keyGuess])

			// Add intersection values
			for i := 0; i < len(intersections); i++ {
				comparisonMap[intersections[i]] = intersections[i]
			}

			// Add difference values
			for i := 0; i < len(diff); i++ {
				if len(diff2) == 0 || i >= len(diff2) {
					comparisonMap[diff[i]] = -1
				} else {
					comparisonMap[diff[i]] = diff2[i]
				}
			}

			for guessIndex, wordIndex := range comparisonMap {
				if wordIndex == -1 {
					// If not contains
					feedbackMap[guessIndex] = "_"
					continue
				} else if wordIndex != guessIndex {
					// Contains but position wrong
					feedbackMap[guessIndex] = "?"
					contains--
				} else if wordIndex == guessIndex {
					// Contains and position are correct
					feedbackMap[guessIndex] = "0"
					position++
				}
			}
		}

		for i := 0; i <= 4; i++ {
			// Create colorful feedback
			feedback += feedbackMap[i]
		}

		// Print feedback and restart for next round
		fmt.Printf("%+v\t%+d%+d\n", feedback, position, contains)
		feedback = ""
		// Create print conditions

		// When find number correctly exit the program
		if position == 4 {
			fmt.Printf("Congratulations! You find secret number!")
			break
		}
	}
}

func Intersection(first, second []int) []int {
	var intersections []int

	for _, i := range first {
		for _, j := range second {
			if i == j {
				intersections = append(intersections, i)
			}
		}
	}
	return intersections
}

// Difference returns the elements in `a` that aren't in `b`.
func Difference(a, b []int) []int {
	mb := make(map[int]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []int
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
