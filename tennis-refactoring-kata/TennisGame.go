package tennis

type Game interface {
	WonPoint(playerName string)
	GetScore() string
}

type game struct {
	mScore1     int
	mScore2     int
	player1Name string
	player2Name string
}

func NewGame(player1Name string, player2Name string) Game {
	game := &game{
		player1Name: player1Name,
		player2Name: player2Name,
	}

	return game
}

func (game *game) WonPoint(playerName string) {
	if playerName == "player1" {
		game.mScore1 += 1
	} else {
		game.mScore2 += 1
	}
}

func (game *game) GetScore() string {
	score := ""
	tempScore := 0
	if game.mScore1 == game.mScore2 {
		switch game.mScore1 {
		case 0:
			score = "Love-All"
		case 1:
			score = "Fifteen-All"
		case 2:
			score = "Thirty-All"
		default:
			score = "Deuce"
		}
	} else if game.mScore1 >= 4 || game.mScore2 >= 4 {
		minusResult := game.mScore1 - game.mScore2
		if minusResult == 1 {
			score = "Advantage player1"
		} else if minusResult == -1 {
			score = "Advantage player2"
		} else if minusResult >= 2 {
			score = "Win for player1"
		} else {
			score = "Win for player2"
		}
	} else {
		for i := 1; i < 3; i++ {
			if i == 1 {
				tempScore = game.mScore1
			} else {
				score += "-"
				tempScore = game.mScore2
			}
			switch tempScore {
			case 0:
				score += "Love"
			case 1:
				score += "Fifteen"
			case 2:
				score += "Thirty"
			case 3:
				score += "Forty"
			}
		}
	}
	return score
}
