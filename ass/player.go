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
	// I/O: first output requested, after sending number of rounds to input
	Name() string
	// Move outputs the next move of the player through thr channel
	// I/O: called after name and before each round
	Move(chan<- PlayerMove)
	// Feedback gives the last opponents move and outcome of the game
	// Must send struct{}{} through channel when finished processing
	// I/O: input sent after every read from move
	Feedback(PlayerMove, GameOutcome, chan<- struct{})
}
