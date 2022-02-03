package questions

import (
	"errors"
	"fmt"
	"gopkg.in/AlecAivazis/survey.v1"
	"regexp"
)

type Answers struct {
	Token   string
	Message string
	Channel string
	Image   string
}

var Banner = `
██████╗░░█████╗░████████╗  ██████╗░░█████╗░░██████╗████████╗███████╗██████╗░
██╔══██╗██╔══██╗╚══██╔══╝  ██╔══██╗██╔══██╗██╔════╝╚══██╔══╝██╔════╝██╔══██╗
██████╦╝██║░░██║░░░██║░░░  ██████╔╝██║░░██║╚█████╗░░░░██║░░░█████╗░░██████╔╝
██╔══██╗██║░░██║░░░██║░░░  ██╔═══╝░██║░░██║░╚═══██╗░░░██║░░░██╔══╝░░██╔══██╗
██████╦╝╚█████╔╝░░░██║░░░  ██║░░░░░╚█████╔╝██████╔╝░░░██║░░░███████╗██║░░██║
╚═════╝░░╚════╝░░░░╚═╝░░░  ╚═╝░░░░░░╚════╝░╚═════╝░░░░╚═╝░░░╚══════╝╚═╝░░╚═╝

`

var Questions = []*survey.Question{
	{
		Name: "token",
		Prompt: &survey.Input{
			Message: "Bot token:",
		},
		Validate: shouldMatchRegex(`^(.+)\.(.+)\.(.+)$`),
	},
	{
		Name: "channel",
		Prompt: &survey.Input{
			Message: "Channel ID:",
		},
		Validate: shouldMatchRegex(`^\d{18}$`),
	},
	{
		Name: "message",
		Prompt: &survey.Input{
			Message: "Message (formatting: https://bit.ly/3roY4oQ):",
		},
	},
	{
		Name: "image",
		Prompt: &survey.Input{
			Message: "Attachment (image or gif url):",
			Default: "none",
		},
		Validate: shouldMatchRegex(`(none)|(https://(.+))`),
	},
}

func shouldMatchRegex(regex string) func(interface{}) error {
	return func(ans interface{}) error {
		answer := fmt.Sprintf("%v", ans)
		match, err := regexp.Match(regex, []byte(answer))
		if err != nil {
			return err
		}
		if !match {
			return errors.New("invalid")
		}
		return nil
	}
}
