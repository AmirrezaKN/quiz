package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

// Question is the struct for questions
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
	var limit time.Duration = 10
	timer := time.NewTimer(limit * time.Second)
Loop:
	for !flag {
		if len(questions) == 0 {
			break
		}

		randomNumber := randGen(0, len(questions)-1)
		fmt.Println("Type The Answer Of The Question:", questions[randomNumber].Quest)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scan(&answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			flag = true
			break Loop
		case answer := <-answerCh:
			if answer == questions[randomNumber].Answer {
				score++
				fmt.Println("Correct! Well done Sekiro")
				delSliceItem(&questions, randomNumber)
			} else {
				flag = true
				break Loop
			}
		}

	}

	if flag {
		fmt.Println("\nGame Over!\nYour Score Is:", score)
	} else {
		fmt.Println("\nYou Have Answerd All Of Our Questions!!\nYou Are Our CHAMPION!!!\nThanks for playing.\nYour Score Is:", score)
	}
}

func randGen(min int, max int) int {
	time.Sleep(10)
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(max-min+1) + min)
}

func delSliceItem(slice *[]*Question, index int) {
	(*slice)[index] = (*slice)[len(*slice)-1]
	(*slice)[len(*slice)-1] = &Question{}
	(*slice) = (*slice)[:len(*slice)-1]
}
