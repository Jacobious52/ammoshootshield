package ioplayer

import (
	"io"
	"log"
	"math/rand"
	"os/exec"

	"github.com/Jacobious52/ammoshootshield/ass"
)

// IOPlayer is a player that reads and writes from an io device
// Is an external program reading and writing from stdin/stdout
// Will be executed on creatino and closed when game ended
type IOPlayer struct {
	Cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
	name   string
}

// NewIOPlayer creates a new IO player and sends the number of rounds to stdin
// the will read player name from stdout
func NewIOPlayer(source string, rounds int) ass.Player {
	cmd := exec.Command(source)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalf("error obtaining stdin from %v: %v", source, err)
		return nil
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("error obtaining stdout from %v: %v", source, err)
		return nil
	}

	p := &IOPlayer{
		Cmd:    cmd,
		stdin:  stdin,
		stdout: stdout,
	}

	if err = cmd.Start(); err != nil {
		log.Fatalf("error faile to start %v: %v", source, err)
		return nil
	}

	return p
}

// Name returns the io players name
func (p *IOPlayer) Name() string {
	return p.name
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
// players random number (move)
func (p *IOPlayer) Move() ass.PlayerMove {
	return ass.PlayerMove(rand.Intn(3))
}

// Feedback gets the last game results
// Do nothing as it's just random so we don't learn
func (p *IOPlayer) Feedback(yourMove, opponentsMove ass.PlayerMove) {
}

// BeginMatch match is about to start
func (p *IOPlayer) BeginMatch() {
}

// EndMatch match has ended. final result
func (p *IOPlayer) EndMatch(finalOutcome ass.GameOutcome) {
}

// BeginGame match is about to start
func (p *IOPlayer) BeginGame() {
}

// EndGame Game has ended. Result of game
func (p *IOPlayer) EndGame(outcome ass.GameOutcome) {
}
