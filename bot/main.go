package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	Token string
)

type payload struct {
	TYPE string `json:"type"`
	EVENT string `json:"event"`
	DATA string `json:"data"`
}

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

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	b, err := json.Marshal(payload {
		"bot",
		"",
		"",
	})
	conn.Write(b)

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		b, err := json.Marshal(payload {
			"bot",
			"MESSAGE_CREATE",
			m.Content,
		})
		if err != nil {
			return
		}

		conn.Write(b)
	})

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
