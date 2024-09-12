package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"DiscordBot/handlers"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// Function to handle messages
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "!ping", respond with "Pong!"
	if m.Content == "!help" {
		s.ChannelMessageSend(m.ChannelID, "These are the commands you can use:\n!ping - Pong!\n!hello - Greeting \n!hi -savage greeting \n!whoami - your role \n!trivia - trivia question \n!describeme - describe yourself")
	}
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
	if m.Content == "!hello" {
		s.ChannelMessageSend(m.ChannelID, "Hello Muji")
	}
	if m.Content == "!hi" {
		s.ChannelMessageSend(m.ChannelID, "Hi Muji k xaa tero khabaar aaja gula chilaide")
	}
	if m.Content == "!whoami" {
		s.ChannelMessageSend(m.ChannelID, "You are a good lady boy from the Bang la Desh")
	}
	if m.Content == "!describeme" {
		s.ChannelMessageSend(m.ChannelID, "What is your name?")
	}
	if m.Content == "Aashish" {
		s.ChannelMessageSend(m.ChannelID, "तपाईं Aashish चञ्चलानीको सस्तो संस्करण हुनुहुन्छ। तपाईंसँग एक अनौठो क्षमता छ—सस्तो रचनात्मकता प्रदर्शन गर्नु र आफ्नो अल्छि स्वभावसँगै धेरै खाने सन्दर्भमा चर्चित हुनु। तपाईंलाई ब्वासो भनेर चिनिन्छ र तपाईंका साथीहरूलाई हाँसो र रमाइलो ल्याउनका लागि जानिन्छ।")
	}
	// Handle the !trivia command
	if m.Content == "!trivia" {
		question, answer, err := handlers.FetchTrivia()
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Sorry, I couldn't fetch a trivia question.")
			return
		}

		// Send the trivia question and provide the answer after 10 seconds
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Here's your trivia question:\n**%s**", question))

		// Send the answer after a short delay
		go func() {
			<-time.After(10 * time.Second) // 10-second delay
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("The correct answer is: **%s**", answer))
		}()
	}
}

func main() {
	// Create a new Discord session using the bot token
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	Token := os.Getenv("DISCORD_BOT_TOKEN")
	if Token == "" {
		fmt.Println("Error: DISCORD_BOT_TOKEN environment variable not set.")
		return
	}
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for the MessageCreate events
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session,", err)
		return
	}
	fmt.Println("Bot is now running. Press CTRL+C to exit.")

	// Wait here until CTRL+C or other term signal is received
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	// Cleanly close down the Discord session
	dg.Close()
}
