package main

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/exp/slog"
)

const discordToken = ""

func main() {
	api, err := discordgo.New(discordToken)
	if err != nil {
		slog.Error("Unable to init Discord API", "error", err)
		os.Exit(1)
	}

	if err = api.Open(); err != nil {
		slog.Error("Unable to connect to Discord API", "error", err)
		os.Exit(1)
	}
	defer api.Close()

	slog.Info(api.State.SessionID)
}
