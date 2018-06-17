package ioplayer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strconv"

	"github.com/Jacobious52/ammoshootshield/ass"
)

type command string

const (
	cmdName       = "name"
	cmdMatchBegin = "match_begin"
	cmdMatchEnd   = "match_end"
	cmdGameBegin  = "game_begin"
	cmdGameEnd    = "game_end"
	cmdMove       = "move"
	cmdFeedback   = "feedback"
)

// IOPlayer is a player that reads and writes from an io device
// Is an external program reading and writing from stdin/stderr
// Will be executed on creatino and closed when game ended
type IOPlayer struct {
	Cmd     *exec.Cmd
	stdin   io.WriteCloser
	stderr  io.ReadCloser
	inBuff  *bufio.Writer
	outBuff *bufio.Scanner
}

// NewIOPlayer creates a new IO player and sends the number of rounds to stdin
// the will read player name from stderr
func NewIOPlayer(source string, rounds int) ass.Player {
	cmd := exec.Command(source)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalf("error obtaining stdin from %v: %v", source, err)
		return nil
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("error obtaining stderr from %v: %v", source, err)
		return nil
	}

	if err = cmd.Start(); err != nil {
		log.Fatalf("error failed to start %v: %v", source, err)
		return nil
	}

	return &IOPlayer{
		Cmd:     cmd,
		stdin:   stdin,
		stderr:  stderr,
		inBuff:  bufio.NewWriter(stdin),
		outBuff: bufio.NewScanner(stderr),
	}
}

// Name returns the io players name
func (p *IOPlayer) Name() string {
	err := p.writeRequest(cmdName)
	if err != nil {
		fmt.Println(err)
	}
	name, err := p.readResponse()
	if err != nil {
		log.Println("no name returned from IOPlayer")
		return ""
	}
	return name
}

// writeRequest to the players stdin
func (p *IOPlayer) writeRequest(cmd command, args ...string) error {
	if _, err := p.inBuff.WriteString(string(cmd)); err != nil {
		return err
	}
	for _, arg := range args {
		if _, err := p.inBuff.WriteString(" "); err != nil {
			return err
		}
		if _, err := p.inBuff.WriteString(arg); err != nil {
			return err
		}
	}
	p.inBuff.WriteString("\n")
	if err := p.inBuff.Flush(); err != nil {
		return err
	}
	return nil
}

// readResponse from the players stderr
func (p *IOPlayer) readResponse() (string, error) {
	if ok := p.outBuff.Scan(); !ok {
		return "", errors.New("no command returned from IOPlayer")
	}
	return p.outBuff.Text(), nil
}

// Move returns the move of the next play
// players random number (move)
func (p *IOPlayer) Move() ass.PlayerMove {
	p.writeRequest(cmdMove)
	moveStr, _ := p.readResponse()
	move, err := strconv.Atoi(moveStr)
	if err != nil {
		log.Println("incorrect move format from IOPlayer:", err)
		return ass.AmmoMove
	}
	return ass.PlayerMove(move)
}

// Feedback gets the last game results
// Do nothing as it's just random so we don't learn
func (p *IOPlayer) Feedback(yourMove, opponentsMove ass.PlayerMove) {
	p.writeRequest(cmdFeedback, fmt.Sprintf("%d", int(yourMove)), fmt.Sprintf("%d", int(opponentsMove)))
}

// BeginMatch match is about to start
func (p *IOPlayer) BeginMatch(rounds int, winCondition ass.GameOutcome) {
	p.writeRequest(cmdMatchBegin, fmt.Sprintf("%d", rounds), fmt.Sprintf("%d", int(winCondition)))
}

// EndMatch match has ended. final result
func (p *IOPlayer) EndMatch(finalOutcome ass.GameOutcome) {
	p.writeRequest(cmdMatchEnd, fmt.Sprintf("%d", int(finalOutcome)))
	err := p.Cmd.Wait()
	if err != nil {
		log.Println("failed to shutdown:", err)
	}
}

// BeginGame match is about to start
func (p *IOPlayer) BeginGame() {
	p.writeRequest(cmdGameBegin)
}

// EndGame Game has ended. Result of game
func (p *IOPlayer) EndGame(outcome ass.GameOutcome) {
	p.writeRequest(cmdGameEnd, fmt.Sprintf("%d", int(outcome)))
}
