package store

import (
	"TrelloReportTools/config"

	"github.com/adlio/trello"
)

type TrelloDatabase struct {
	client *trello.Client
}

func (trelloDB TrelloDatabase) Start() TrelloDatabase {
	conf := config.Setup()
	trelloDB.client = trello.NewClient(conf.KeyApp, conf.Token)

	return trelloDB
}

func (trelloDB *TrelloDatabase) SelectManyTrello(query string) (cards []*trello.Card, err error) {
	cards, err = trelloDB.client.SearchCards(query, trello.Defaults())
	if err != nil {
		return nil, err
	}

	return
}
