package models

import (
	"TrelloReportTools/config"

	"github.com/adlio/trello"
)

var client *trello.Client

type DBTrelloCard struct{}

func init() {
	conf := config.Setup()
	client = trello.NewClient(conf.KeyApp, conf.Token)
}

func (*DBTrelloCard) GetMany(query string) (cards []*trello.Card, err error) {
	cards, err = client.SearchCards(query, trello.Defaults())
	if err != nil {
		return nil, err
	}

	return
}
