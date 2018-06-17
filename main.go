package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Jacobious52/ammoshootshield/players/ammoshootplayer"
	"github.com/Jacobious52/ammoshootshield/players/ioplayer"

	"github.com/Jacobious52/ammoshootshield/ass"
	"github.com/Jacobious52/ammoshootshield/players/randomplayer"
	"gopkg.in/alecthomas/kingpin.v2"
)

// commandline args
var (
	player1Source = kingpin.Flag("player1", "player1 source").Default("player_random").String()
	player2Source = kingpin.Flag("player2", "player2 source").Default("player_random").String()
	rounds        = kingpin.Flag("rounds", "rounds to play in a match").Default("1000").Int()
)

// createPlayersMap is where packaged players are registered for use from the commandline
// *** add your player in here ***
// map key MUST start with 'player_' to work with cmd parser
var createPlayersMap = map[string]func() ass.Player{
	"player_random": func() ass.Player {
		return new(randomplayer.RandomPlayer)
	},
	"player_ammoshoot": func() ass.Player {
		return &ammoshootplayer.AmmoShootPlayer{LastMove: ass.ShootMove}
	},
	/*
		"player_example": func() ass.Player {
			return exampleplayer.NewExamplePlayer(*rounds)
		},
	*/
}

// parsePlayer reads the player source string and creates either a registered player
// or I/O device player if does not start with 'player_'
func parsePlayer(playerStr string) *ass.PlayerController {
	if strings.HasPrefix(playerStr, "player_") {
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
