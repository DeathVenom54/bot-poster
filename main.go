package main

import (
	"fmt"
	"github.com/DeathVenom54/bot-poster/poster"
	"github.com/DeathVenom54/bot-poster/questions"
	"gopkg.in/AlecAivazis/survey.v1"
)

func main() {
	fmt.Println(questions.Banner)

	data := questions.Answers{}
	err := survey.Ask(questions.Questions, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = poster.Post(&data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
