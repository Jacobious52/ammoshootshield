package ass

// Add your players as a new folder under the ./players/ directory'
// package name should be in the form `<name>player` where name is the name of the player
// package should be all lowercase
// register player in the main.go createPlayersMap for use from commandline

// PlayerController contains the stats for a player and embeds a Player logic interface
type PlayerController struct {
	Ammo, Wins int
	Player
}

// Player interface is the logic for a certain player
// implement this in a new package to use in main as a player
type Player interface {
	// Name returns the printable name of the player
	Name() string

	// BeginMatch is called before the entire match is started
	// param 1 is number of rounds to be played
	// param 2 is which player you are (1 or 2) (your win condition)
	// Used for setup.. loading neuralnets and such
	BeginMatch(int, GameOutcome)
	// EndMatch is called after the entire match is finished
	// Used for saving what you learned
	// param 1 is the outcome of the entire match
	EndMatch(GameOutcome)
	// BeginGame runs before a game is played
	BeginGame()
	// EndGame is played after a game is finished
	// Game result returns the winner of the game
	EndGame(GameOutcome)

	// Move should return the next move of the player
	Move() PlayerMove
	// Feedback returns your last move and the last opponents move
	// param 1 is your move, param 2 is their move
	Feedback(PlayerMove, PlayerMove)
}
