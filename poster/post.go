package poster

import (
	"fmt"
	"github.com/DeathVenom54/bot-poster/questions"
	"github.com/bwmarrin/discordgo"
)

func Post(data questions.Answers) error {
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
