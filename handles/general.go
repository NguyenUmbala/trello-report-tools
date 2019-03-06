package handles

import (
	"TrelloReportTools/modules"

	"github.com/adlio/trello"
)

// if trello.Card != modules.Card => return false
func CompareTwoCard(cardsBoard *trello.Card, cardsDB modules.Card) bool {
	if cardsBoard.ID != cardsDB.ID {
		return false
	}
	if cardsBoard.Due == nil && cardsDB.Due == nil {
		return false
	}
	if cardsBoard.Due == nil && cardsDB.Due != nil {
		return true
	}
	if cardsBoard.Due != nil && cardsDB.Due == nil {
		return true
	}
	if cardsBoard.Due.String() == cardsDB.Due.String() {
		return false
	}
	return true
}
