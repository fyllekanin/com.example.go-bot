package handlers

import "github.com/bwmarrin/discordgo"

func OnMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content != "do-embed" {
		return
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Type: "rich",
		Title: "Awesome",
		Color: 1,
		Description: "This is a description",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "I am a field",
				Value:  "I am a value",
				Inline: true,
			},
		},
	})
}