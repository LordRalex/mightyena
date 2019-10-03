package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Deop struct {
	Connection *irc.Connection
	Channel    core.Channel
	Target     core.User
	Sender     core.User
}

func (d *Deop) EventName() string {
	return "deop"
}
