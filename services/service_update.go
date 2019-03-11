package services

import (
	"TrelloReportTools/models"
	"TrelloReportTools/utils"

	"github.com/adlio/trello"
)

type ServiceUpdate struct{}

var UtilString utils.UtilString
var UtilTime utils.UtilTime

func (*ServiceUpdate) UpdateCard(card *trello.Card) {
	timeGuess := UtilString.GetTimeGuessForDone(card.Name)
	timeReal := UtilString.GetRealTimeOfDone(card.Name)

	db.InsertOrUpdate(models.NewCard(card, timeGuess, timeReal))
}
