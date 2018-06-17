package ammoshootplayer

import (
	"github.com/Jacobious52/ammoshootshield/ass"
)

// AmmoShootPlayer plays a random move everytime
type AmmoShootPlayer struct {
	LastMove ass.PlayerMove
}

// Name returns the name of this player
func (p *AmmoShootPlayer) Name() string {
	return "ammoshoot"
}

// Move returns the move of the next play
// players random number (move)
func (p *AmmoShootPlayer) Move(moveChan chan<- ass.PlayerMove) {
	if p.LastMove == ass.AmmoMove {
		p.LastMove = ass.ShootMove
	} else if p.LastMove == ass.ShootMove {
		p.LastMove = ass.AmmoMove
	}
	moveChan <- p.LastMove
}

// Feedback gets the last game results
// Do nothing as it's just random so we don't learn
func (p *AmmoShootPlayer) Feedback(opponentsMove ass.PlayerMove, lastOutcome ass.GameOutcome, doneChan chan<- struct{}) {
	if lastOutcome != ass.NoWinner {
		p.LastMove = ass.ShootMove
	}
	doneChan <- struct{}{}
}
