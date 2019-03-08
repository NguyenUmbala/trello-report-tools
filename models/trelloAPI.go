package models

import (
	"TrelloReportTools/config"

	"github.com/adlio/trello"
)

type TrelloAPI struct {
	client *trello.Client
}

func (trelloAPI TrelloAPI) GetCards(query string) ([]*trello.Card, error) {
	trelloAPI.client = trello.NewClient(config.KeyApp, config.Token)
	cards, err := trelloAPI.client.SearchCards(query, trello.Defaults())

	return cards, err
}
