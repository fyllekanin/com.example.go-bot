package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
	"net/http"
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
	d := net.Dialer{Timeout: 30 * time.Second}
	conn, err := d.Dial("tcp", "balancer:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	b, err := json.Marshal(payload {
		"worker",
		"declaring",
		"",
	})
	conn.Write(b)
	fmt.Println("Sent first event to server")

	start(conn)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}



func start(conn net.Conn) {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	for {
		d := json.NewDecoder(conn)
		var msg payload
		d.Decode(&msg)

		if msg.DATA != "" {
			http.Get("https://webhook.site/e1edf837-8e7d-45a4-a72b-0714610a45f7?worker=two&content=" + msg.DATA)
			log.Print("Server relay:", msg)
		}
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
}

