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
		Validate: func(ans interface{}) error {
			token := fmt.Sprintf("%v", ans)
			match, err := regexp.Match(`^(.+)\.(.+)\.(.+)$`, []byte(token))
			if err != nil {
				return err
			}
			if !match {
				return errors.New("invalid token")
			}
			return nil
		},
	},
	{
		Name: "channel",
		Prompt: &survey.Input{
			Message: "Channel ID:",
		},
		Validate: func(ans interface{}) error {
			token := fmt.Sprintf("%v", ans)
			match, err := regexp.Match(`^\d{18}$`, []byte(token))
			if err != nil {
				return err
			}
			if !match {
				return errors.New("invalid channel id")
			}
			return nil
		},
	},
	{
		Name: "message",
		Prompt: &survey.Input{
			Message: "Message (check formatting.txt):",
		},
		Validate: survey.Required,
	},
	{
		Name: "image",
		Prompt: &survey.Input{
			Message: "Attachment (image or gif url):",
			Default: "none",
		},
	},
}
