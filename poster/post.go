package poster

import (
	"fmt"
	"github.com/DeathVenom54/bot-poster/questions"
	"github.com/bwmarrin/discordgo"
	"regexp"
)

func Post(data *questions.Answers) error {
	data.Message = formatMessage(data.Message)
	s, err := discordgo.New("Bot " + data.Token)
	if err != nil {
		return err
	}

	if data.Image != "none" {
		_, err := s.ChannelMessageSend(data.Channel, data.Image)
		if err != nil {
			return err
		}
	}

	_, err = s.ChannelMessageSend(data.Channel, data.Message)
	if err != nil {
		return err
	}

	channel, err := s.Channel(data.Channel)
	if err != nil {
		return err
	}

	fmt.Printf("Posted message to #%s successfully", channel.Name)
	return nil
}

func formatMessage(m string) string {
	// user mentions
	r := regexp.MustCompile(`u(\d{18})`)
	m = r.ReplaceAllString(m, "<@$1>")

	// role mentions
	r = regexp.MustCompile(`r(\d{18})`)
	m = r.ReplaceAllString(m, "<@&$1>")

	// channel mentions
	r = regexp.MustCompile(`c(\d{18})`)
	m = r.ReplaceAllString(m, "<#$1>")

	return m
}
