package main

import (
	"crypto/tls"
	"fmt"

	scripts "github.com/ecdsa521/newbot/scripts"
	irc "github.com/thoj/go-ircevent"
)

func onConnect(event *irc.Event, irc *irc.Connection) {
	irc.Join("#moobot")
	irc.Join("#cat")

}
func onPing(event *irc.Event, irc *irc.Connection) {
	fmt.Printf("PING: %v\n", event.Message())
}

const serverssl = "127.0.0.1:7000"

func main() {
	irccon := irc.IRC("moobot", "moobot")
	irccon.UseTLS = true
	irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := irccon.Connect(serverssl)
	if err != nil {
		fmt.Printf("Err %s\n", err)
		return
	}

	irccon.AddCallback("001", func(e *irc.Event) {
		go onConnect(e, irccon)
	})
	irccon.AddCallback("PING", func(e *irc.Event) {
		go onPing(e, irccon)
	})
	irccon.AddCallback("PRIVMSG", func(e *irc.Event) {
		go scripts.Call(e, irccon, "PRIVMSG")
	})
	irccon.AddCallback("JOIN", func(e *irc.Event) {
		go scripts.Call(e, irccon, "JOIN")
	})
	irccon.AddCallback("PART", func(e *irc.Event) {
		go scripts.Call(e, irccon, "PART")
	})
	irccon.AddCallback("KICK", func(e *irc.Event) {
		go scripts.Call(e, irccon, "KICK")
	})
	irccon.AddCallback("MODE", func(e *irc.Event) {
		go scripts.Call(e, irccon, "MODE")
	})

	irccon.Loop()
}
