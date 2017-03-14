package scripts

import (
	"fmt"

	irc "github.com/thoj/go-ircevent"
)

//ScriptHooks ...
var ScriptHooks = make(map[string][]func(event *irc.Event, irc *irc.Connection))

//ScriptInterface ...
type ScriptInterface interface {
	OnMessage(event *irc.Event, irc *irc.Connection)
	OnJoin(event *irc.Event, irc *irc.Connection)
	OnPart(event *irc.Event, irc *irc.Connection)
	OnMode(event *irc.Event, irc *irc.Connection)
	OnKick(event *irc.Event, irc *irc.Connection)
}

func init() {
	fmt.Println("Hello from _scripts.go#init")
}

//Call ..
func Call(event *irc.Event, irc *irc.Connection, et string) {
	for i := range ScriptHooks[et] {
		ScriptHooks[et][i](event, irc)
	}
}
