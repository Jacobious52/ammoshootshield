package ass

// PlayerMove defines the value of a move the player can take
type PlayerMove int

const (
	// AmmoMove is to increase ammo by one
	AmmoMove PlayerMove = iota
	// ShootMove means to shoot other player. ammo must be > 0 to work
	ShootMove
	// ShieldMove means to put a shield up and block a shot
	ShieldMove
)

// encoded move combinations
var (
	playP2AmmoP1Ammo     = encodeMoves(AmmoMove, AmmoMove)
	playP2AmmoP1Shoot    = encodeMoves(ShootMove, AmmoMove)
	playP2AmmoP1Shield   = encodeMoves(ShieldMove, AmmoMove)
	playP2ShootP1Ammo    = encodeMoves(AmmoMove, ShootMove)
	playP2ShootP1Shoot   = encodeMoves(ShootMove, ShootMove)
	playP2ShootP1Shield  = encodeMoves(ShieldMove, ShootMove)
	playP2ShieldP1Ammo   = encodeMoves(AmmoMove, ShieldMove)
	playP2ShieldP1Shoot  = encodeMoves(ShootMove, ShieldMove)
	playP2ShieldP1Shield = encodeMoves(ShieldMove, ShieldMove)
)

// encodeMoves encodes a combination of moves into a single int
func encodeMoves(m1, m2 PlayerMove) PlayerMove {
	return m1 | (m2 << 2)
}

// moveMap contains the dispatch for all combinations of moves and their logic
var moveMap = map[PlayerMove]func(*PlayerController, *PlayerController) GameOutcome{
	playP2AmmoP1Ammo: func(p1, p2 *PlayerController) GameOutcome { // 0000 : ammo ammo
		p1.Ammo++
		p2.Ammo++
		return NoWinner
	},
	playP2AmmoP1Shoot: func(p1, p2 *PlayerController) GameOutcome { // 0001 : ammo shoot
		if p1.Ammo > 0 {
			return Player1Wins
		}
		p2.Ammo++
		return NoWinner
	},
	playP2AmmoP1Shield: func(p1, p2 *PlayerController) GameOutcome { // 0010 : ammo sheild
		p2.Ammo++
		return NoWinner
	},
	playP2ShootP1Ammo: func(p1, p2 *PlayerController) GameOutcome { // 0100 : shoot ammo
		if p2.Ammo > 0 {
			return Player2Wins
		}
		p1.Ammo++
		return NoWinner
	},
	playP2ShootP1Shoot: func(p1, p2 *PlayerController) GameOutcome { // 0101 : shoot shoot
		if p1.Ammo == p2.Ammo {
			if p1.Ammo > 0 {
				p1.Ammo--
				p2.Ammo--
			}
			return NoWinner
		}
		if p1.Ammo > p2.Ammo {
			return Player1Wins
		}
		return Player2Wins
	},
	playP2ShootP1Shield: func(p1, p2 *PlayerController) GameOutcome { // 0110 : shoot sheild
		if p2.Ammo > 0 {
			if p2.Ammo >= 5 {
				return Player2Wins
			}
			p2.Ammo--
		}
		return NoWinner
	},
	playP2ShieldP1Ammo: func(p1, p2 *PlayerController) GameOutcome { // 1000 : sheild ammo
		p1.Ammo++
		return NoWinner
	},
	playP2ShieldP1Shoot: func(p1, p2 *PlayerController) GameOutcome { // 1001 : sheild shoot
		if p1.Ammo > 0 {
			if p1.Ammo >= 5 {
				return Player1Wins
			}
			p1.Ammo--
		}
		return NoWinner
	},
	playP2ShieldP1Shield: func(p1, p2 *PlayerController) GameOutcome { // 1010 : shield shield
		return NoWinner
	},
}
