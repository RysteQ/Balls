package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)
var token string
func main () {
	token := "Nzg1NTUyNTUyOTUwNzU5NDc0.X85g0w.a-bXdhQdMQ5UqXFvuCXKSV-NhwY"

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("something bad happened :(", err)
		return
	}

	discord.AddHandler(onMessage)

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("THE BOT IS WORKING I THINK")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}




func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.Contains(m.Content, "balls") {
		s.ChannelMessageSend(m.ChannelID, "BALLS BALLS BALLS BALLLSL BALLS!")
	}
}