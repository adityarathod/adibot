package main

import (
	"log"
	"math/rand/v2"
	"os"
	"os/signal"

	"github.com/adityarathod/adibot/config"
	"github.com/adityarathod/adibot/llm"
	"github.com/bwmarrin/discordgo"
)

func main() {
	config, err := config.LoadBotConfig("bot-config.json")
	if err != nil {
		log.Fatalf("could not load bot config: %s", err)
	}
	llmClient := llm.NewLLMClient(config.ModelEndpoint)
	session, _ := discordgo.New("Bot " + config.Token)

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as %s", r.User.String())
	})

	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore channels/users that are not allowlisted
		if !config.IsChannelAllowlisted(m.ChannelID) || !config.IsUserAllowlisted(m.Author.ID) {
			return
		}

		// Ignore all messages created by the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		if config.ReplyRatelimit.Enabled && rand.IntN(10) >= int(10*config.ReplyRatelimit.Proportion) {
			log.Printf("Ignoring message: %s", m.Content)
			s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
			return
		}

		s.ChannelTyping(m.ChannelID)
		response := llmClient.CallModelAPI(m.Content)

		if response == "" {
			s.MessageReactionAdd(m.ChannelID, m.ID, "💤")
			return
		}
		_, err := s.ChannelMessageSendReply(m.ChannelID, response, m.SoftReference())
		if err != nil {
			log.Printf("could not send message: %s", err)
		}
	})

	err = session.Open()
	if err != nil {
		log.Fatalf("could not open session: %s", err)
	}

	// Wait for an interrupt signal to gracefully close the session
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	<-sigch
	err = session.Close()
	if err != nil {
		log.Printf("could not close session gracefully: %s", err)
	}
}
