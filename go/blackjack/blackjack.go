package blackjack

var m = map[string]int{
	"ace":   11,
	"jack":  10,
	"queen": 10,
	"king":  10,
	"ten":   10,
	"nine":  9,
	"eight": 8,
	"seven": 7,
	"six":   6,
	"five":  5,
	"four":  4,
	"three": 3,
	"two":   2,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	v, ok := m[card]
	if !ok {
		return 0
	}
	return v
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return (ParseCard(card1) + ParseCard(card2)) == 21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	switch {
	case !isBlackjack:
		return "P"
	case dealerScore < 10:
		return "W"
	default:
		return "S"
	}
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	switch {
	case handScore >= 17:
		return "S"
	case handScore <= 11:
		return "H"
	case dealerScore >= 7:
		return "H"
	default:
		return "S"
	}
}
