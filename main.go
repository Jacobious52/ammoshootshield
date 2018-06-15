package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Jacobious52/ass/players/ioplayer"

	"github.com/Jacobious52/ass/ass"
	"github.com/Jacobious52/ass/players/randomplayer"
	"gopkg.in/alecthomas/kingpin.v2"
)

// commandline args
var (
	player1Source = kingpin.Flag("player1", "player1 source").Default("$random").String()
	player2Source = kingpin.Flag("player2", "player2 source").Default("$random").String()
	rounds        = kingpin.Flag("rounds", "rounds to play in a match").Default("1000").Int()
)

// createPlayersMap is where packaged players are registered for use from the commandline
// *** add your player in here ***
// map key MUST start with '$' to work with cmd parser
var createPlayersMap = map[string]func() ass.Player{
	"$random": func() ass.Player {
		return new(randomplayer.RandomPlayer)
	},
	/*
		"$example": func() ass.Player {
			return exampleplayer.NewExamplePlayer(*rounds)
		},
	*/
}

// parsePlayer reads the player source string and creates either a registered player
// or I/O device player if does not start with '$'
func parsePlayer(playerStr string) *ass.PlayerController {
	if strings.HasPrefix(playerStr, "$") {
		if f, ok := createPlayersMap[playerStr]; ok {
			return &ass.PlayerController{Player: f()}
		}
		fmt.Println("cannot find player", playerStr)
		os.Exit(1)
	}
	return &ass.PlayerController{Player: ioplayer.NewIOPlayer(playerStr, *rounds)}
}

func main() {
	kingpin.Parse()
	rand.Seed(time.Now().Unix())

	p1 := parsePlayer(*player1Source)
	p2 := parsePlayer(*player2Source)

	m := &ass.Match{
		Player1:   p1,
		Player2:   p2,
		SleepTime: 5 * time.Millisecond,
		BarWidth:  60,
	}

	m.RunRounds(*rounds)
}
