package scripts

import (
	"fmt"

	irc "github.com/thoj/go-ircevent"
)

//ScriptPing ...
type ScriptPing struct {
	author      string
	description string
	command     string
}

func init() {
	ScriptHooks["PRIVMSG"] = append(ScriptHooks["PRIVMSG"], ScriptPing{
		command:     "!ping",
		author:      "moo",
		description: "Replies with obscenities when you ping it",
	}.OnMessage)

}

//OnMessage ...
func (h ScriptPing) OnMessage(event *irc.Event, irc *irc.Connection) {
	if h.command == event.Message() {
		irc.Privmsg(event.Arguments[0], "pong")
	}
	fmt.Printf("%v | %v\n", h.command, event.Message())
}
