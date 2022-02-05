package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"main/handlers"
	"os"
	"os/signal"
	"syscall"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	start()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}



func start() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(handlers.OnMessage)
	dg.AddHandler(messageDelete)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	message, arr := dg.ChannelMessageSend("934497562369613824", "hejsan")
	if message == nil {
		fmt.Println("gg", arr)
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
}

func messageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {
	fmt.Println("Deleted message")
}
