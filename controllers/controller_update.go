package controllers

type UpdateCard struct{}

func (*UpdateCard) UpdateCardRealTime() {
	idBoard := conf.IDBoard
	for {
		cardsOnBoard := ServiceGet.GetCardsChangedDueDateOnBoard(idBoard)
		for _, value := range cardsOnBoard {
			ServiceUpdate.UpdateCard(value)
		}
	}
}
