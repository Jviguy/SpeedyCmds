package main

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

type Ping struct {
	name string
}

func (p Ping) Setname(newname string) {
	p.name = newname
}

func (p Ping) GetName() string {
	return p.name
}

func (p Ping) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	_, err := session.ChannelMessageSend(ctx.GetChannel().ID, "Pong at " + strconv.Itoa(int(session.HeartbeatLatency().Milliseconds())) + "ms !")
	return err
}

