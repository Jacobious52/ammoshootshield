package randomplayer

import (
	"math/rand"

	"github.com/Jacobious52/ammoshootshield/ass"
)

// RandomPlayer plays a random move everytime
type RandomPlayer struct{}

// Name returns the name of this player
func (p *RandomPlayer) Name() string {
	return "random"
}

// Move returns the move of the next play
// players random number (move)
func (p *RandomPlayer) Move() ass.PlayerMove {
	return ass.PlayerMove(rand.Intn(3))
}

// Feedback gets the last game results
// Do nothing as it's just random so we don't learn
func (p *RandomPlayer) Feedback(yourMove, opponentsMove ass.PlayerMove) {
}

// BeginMatch match is about to start
func (p *RandomPlayer) BeginMatch(rounds int, winCondition ass.GameOutcome) {
}

// EndMatch match has ended. final result
func (p *RandomPlayer) EndMatch(finalOutcome ass.GameOutcome) {
}

// BeginGame match is about to start
func (p *RandomPlayer) BeginGame() {
}

// EndGame Game has ended. Result of game
func (p *RandomPlayer) EndGame(outcome ass.GameOutcome) {
}
