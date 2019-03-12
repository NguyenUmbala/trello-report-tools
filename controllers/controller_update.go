package controllers

type UpdateCard struct{}

func (*UpdateCard) UpdateCardRealTime() {
	idBoard := conf.IDBoard

	for {
		cardsOnTrelloDB := ServiceCard.GetCardsOnBoard(idBoard)
		cardsOnDB := ServiceCard.GetCardsOnDB()
		ServiceCard.UpdateCardsChangedDueDate(cardsOnTrelloDB, cardsOnDB)
	}
}
