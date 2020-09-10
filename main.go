package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

type Question struct {
	Quest  string `csv:"question"`
	Answer string `csv:"answer"`
}

func main() {
	input, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	questions := []*Question{}

	if err := gocsv.UnmarshalFile(input, &questions); err != nil {
		panic(err)
	}

	flag := false
	score := 0
	for !flag {
		var answer string
		randomNumber := randGen(0, len(questions)-1)
		fmt.Println("Type The Answer Of The Question:", questions[randomNumber].Quest)
		fmt.Scan(&answer)

		if answer == questions[randomNumber].Answer {
			score++
			fmt.Println("Correct! Well done Sekiro\n")
		} else {
			break
		}
	}

	fmt.Println("\nGame Over!\nYour Score Is:", score)

	// for index, data := range questions {
	// 	fmt.Println(index, data)
	// }
}

func randGen(min int, max int) int {
	time.Sleep(10)
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(max-min+1) + min)
}
