package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var token = "YOUR-BOT-TOKEN"
var botChID = "CHANNEL-ID"
var Dg *discordgo.Session

func initSession() {
	// Create a new DC session using the bot's token
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	err = dg.Open()

	if err != nil {
		fmt.Println("Error openning connection, ", err)
		l.Fatalf("Error openning connection, ", err)
		return
	}
	Dg = dg

}

func ConnectToDC() {
	initSession()
	go ParseRSS()
	time.Sleep(time.Second)

	// Register the messageCreate func as a callback for MessageCreate events.
	Dg.AddHandler(messageCreate)

	Dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running, press CTRL+C to exit.")
	l.Println("[INFO] DCScragor started to running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	Dg.Close()
	l.Println("[INFO] DCScragor stopped.")
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!dcscragor" {
		s.ChannelMessageSend(m.ChannelID, "Hey! I'm here.")
		l.Printf("DCScragor got a command from %s, Command: !dcscragor", m.Author)
	}

}
