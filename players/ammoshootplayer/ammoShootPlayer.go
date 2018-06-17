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
func (p *AmmoShootPlayer) Move() ass.PlayerMove {
	if p.LastMove == ass.AmmoMove {
		p.LastMove = ass.ShootMove
	} else if p.LastMove == ass.ShootMove {
		p.LastMove = ass.AmmoMove
	}
	return p.LastMove
}

// Feedback gets the last game results
// Do nothing as it's just random so we don't learn
func (p *AmmoShootPlayer) Feedback(yourMove, opponentsMove ass.PlayerMove) {
}

// BeginMatch match is about to start
func (p *AmmoShootPlayer) BeginMatch() {
}

// EndMatch match has ended. final result
func (p *AmmoShootPlayer) EndMatch(finalOutcome ass.GameOutcome) {
}

// BeginGame match is about to start
func (p *AmmoShootPlayer) BeginGame() {
}

// EndGame Game has ended. Result of game
func (p *AmmoShootPlayer) EndGame(outcome ass.GameOutcome) {
	p.LastMove = ass.ShootMove
}
