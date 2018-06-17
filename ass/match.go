package ass

import (
	"fmt"
	"os"
	"time"
)

// GameOutcome is a value representing the outcome of a single round
type GameOutcome int

const (
	// NoWinner means that the game is still going
	NoWinner GameOutcome = iota
	// Player1Wins means that player1 won the game
	Player1Wins
	// Player2Wins means that player2 won the game
	Player2Wins
)

// Match defines a match that plays many rounds
type Match struct {
	Player1, Player2 *PlayerController
	SleepTime        time.Duration
	BarWidth         int
}

// RunRounds runs n rounds and displays the output as a animated bar
func (m *Match) RunRounds(rounds int) {
	name1 := m.Player1.Name()
	name2 := m.Player2.Name()
	m.Player1.BeginMatch(rounds, Player1Wins)
	m.Player2.BeginMatch(rounds, Player2Wins)

	for i := 1; i < rounds+1; i++ {

		m.Player1.BeginGame()
		m.Player2.BeginGame()

		// run this round until it is finished
		var lastOutcome GameOutcome
		for lastOutcome == NoWinner {
			lastOutcome = m.runRound()
		}

		m.Player1.EndGame(lastOutcome)
		m.Player2.EndGame(lastOutcome)

		// update stats and display output
		if lastOutcome == Player1Wins {
			m.Player1.Wins++
		} else if lastOutcome == Player2Wins {
			m.Player2.Wins++
		}
		percentGraphColored(
			os.Stdout,
			name1,
			name2,
			m.BarWidth,
			float64(m.Player1.Wins),
			float64(m.Player2.Wins),
			float64(i),
			float64(rounds),
		)

		// for better visuals of progress bar
		time.Sleep(m.SleepTime)
	}

	var finalOutcome GameOutcome

	// game finished
	if m.Player1.Wins > m.Player2.Wins {
		finalOutcome = Player1Wins
		fmt.Println("Player1 wins!")
	} else if m.Player1.Wins < m.Player2.Wins {
		finalOutcome = Player2Wins
		fmt.Println("Player2 wins!")
	} else {
		fmt.Println("It's a draw!")
	}

	m.Player1.EndMatch(finalOutcome)
	m.Player2.EndMatch(finalOutcome)
}

// runRound runs one round
// does moves then gives feedback to players
func (m *Match) runRound() GameOutcome {
	m1, m2, lastOutcome := m.doMoves()

	// give feedback concurrently
	p1DoneChan := make(chan struct{})
	p2DoneChan := make(chan struct{})

	go func(p1DoneChan chan struct{}) {
		m.Player1.Feedback(m1, m2)
		p1DoneChan <- struct{}{}
	}(p1DoneChan)
	go func(p2DoneChan chan struct{}) {
		m.Player2.Feedback(m2, m1)
		p2DoneChan <- struct{}{}
	}(p2DoneChan)

	<-p1DoneChan
	<-p2DoneChan

	return lastOutcome
}

// doMoves gets moves from players and simulates a result
func (m *Match) doMoves() (PlayerMove, PlayerMove, GameOutcome) {

	// do concurrently
	m1Chan := make(chan PlayerMove)
	m2Chan := make(chan PlayerMove)

	go func(m1Chan chan PlayerMove) {
		m1Chan <- m.Player1.Move()
	}(m1Chan)
	go func(m2Chan chan PlayerMove) {
		m2Chan <- m.Player2.Move()
	}(m2Chan)

	m1 := <-m1Chan
	m2 := <-m2Chan

	// lookup the result of the encoded move combinations in the table
	res := moveMap[encodeMoves(m1, m2)](m.Player1, m.Player2)
	return m1, m2, res
}
