package main

import (
	"fmt"
	"gopkg.in/AlecAivazis/survey.v1"
)

func main() {
	fmt.Println(banner)

	data := Answers{}
	err := survey.Ask(questions, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Post(data)
}
