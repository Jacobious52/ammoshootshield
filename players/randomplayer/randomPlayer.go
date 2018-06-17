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
func (p *RandomPlayer) Move(moveChan chan<- ass.PlayerMove) {
	m := ass.PlayerMove(rand.Intn(3))
	moveChan <- m
}

// Feedback gets the last game results
// Do nothing as it's just random so we don't learn
func (p *RandomPlayer) Feedback(opponentsMove ass.PlayerMove, lastOutcome ass.GameOutcome, doneChan chan<- struct{}) {
	doneChan <- struct{}{}
}
