package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var token =  "BOT-TOKEN"
var botChID = "BOT-RSS-SHARE-CHANNEL-ID"
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

func handleHelp(s *discordgo.Session, m *discordgo.MessageCreate)  {
	s.ChannelMessageSend(m.ChannelID, "Available commands:\n----------------------------\n!dcscragor add_blog <URL>       Adds your RSS feed URL to the scrape wish list.")
}

func handleAddBlogUrltoList(s *discordgo.Session, m *discordgo.MessageCreate) {
	splits := strings.Split(m.Content, " ")
	url := splits[2]
	if len(splits) != 3 {
		s.ChannelMessageSend(m.ChannelID, "You are prompted wrong command to add blog request. Usage: !dcscragor add_blog <URL>")
		l.Printf("DCScragor got a wrong blog adding request from %s, Command: %s", m.Author, m.Content)
	} else{
		file, err := os.OpenFile("blog_request.list", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
		l.Println(err)
		}

		defer file.Close()
		if !strings.Contains(readFile("blog_request.list"), url) {
			if _, err := file.WriteString(url + "\n"); err != nil {
				l.Fatal(err)
			} else{
				s.ChannelMessageSend(m.ChannelID, "Blog URL is added to wish list.")
				l.Printf("DCScragor added a new blog to the wish list. Author: %s, URL: %s", m.Author, url)
			}
			
		} else{
			s.ChannelMessageSend(m.ChannelID, "You've already requested for this URL!")
		}
		
	}
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
		s.ChannelMessageSend(m.ChannelID, "Hey! I'm here. Use !dcscragor help")
		l.Printf("DCScragor got a command from %s, Command: !dcscragor", m.Author)
	}

	// !dcscragor add_blog <URL>
	if strings.HasPrefix(m.Content, "!dcscragor") && len(strings.Split(m.Content, " ")) == 3{

		if strings.Split(m.Content, " ")[1] == "add_blog" {
			if strings.Contains(strings.Split(m.Content, " ")[2], "http://") ||  strings.Contains(strings.Split(m.Content, " ")[2], "https://") {
				handleAddBlogUrltoList(s,m)
			l.Printf("DCScragor got a blog adding request from %s, Command: %s", m.Author, m.Content)
			} else{
				s.ChannelMessageSend(m.ChannelID, "Please use a proper URL!")
			}
			
		} else {
			s.ChannelMessageSend(m.ChannelID, "Please use !dcscragor add_blog <URL> command to request your URL to add wish list.")
		}
	}

	// !dcscragor help
	if strings.HasPrefix(m.Content, "!dcscragor") && len(strings.Split(m.Content, " ")) == 2{
		if strings.Split(m.Content, " ")[1] == "help"{
			handleHelp(s,m)
		}
	}


}
