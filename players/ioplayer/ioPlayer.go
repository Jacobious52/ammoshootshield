package ioplayer

import (
	"github.com/Jacobious52/ass/ass"
)

// IOPlayer is a player that reads and writes from an io device
// Is an external program reading and writing from stdin/stdout
// Will be executed on creatino and closed when game ended
type IOPlayer struct {
}

// NewIOPlayer creates a new IO player and sends the number of rounds to stdin
// the will read player name from stdout
func NewIOPlayer(source string, rounds int) ass.Player {
	return nil
}

// write to the players stdin
func (p *IOPlayer) Write(b []byte) (int, error) {
	return 0, nil
}

// read from the players stdout
func (p *IOPlayer) Read(b []byte) (int, error) {
	return 0, nil
}

// Move returns the move of the next play
// Read from stdout
func (p *IOPlayer) Move(moveChan chan<- ass.PlayerMove) {
	moveChan <- 0
}

// Feedback gets the last game results
// Send to stdin
func (p *IOPlayer) Feedback(opponentsMove ass.PlayerMove, lastOutcome ass.GameOutcome, doneChan chan<- struct{}) {
	doneChan <- struct{}{}
}
