package trelloAPI

import (
	"fmt"

	"github.com/adlio/trello"
)

func GetCardsIsOpenOnWeek(idBoard, nameList string) []*trello.Card {
	cards, err := client.SearchCards("board:"+idBoard+" is:open sort:created created:week list:"+nameList, trello.Defaults())
	if err != nil {
		fmt.Println(err)
	}
	return cards
}

func GetCardsOnBoard(idBoard string) []*trello.Card {
	board, err := client.GetBoard(idBoard, trello.Defaults())
	if err != nil {
		fmt.Println(err)
	}

	cards, err := board.GetCards(trello.Defaults())
	if err != nil {
		fmt.Println(err)
	}

	return cards
}
