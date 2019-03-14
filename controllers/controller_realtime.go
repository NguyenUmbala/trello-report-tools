package controllers

func UpdateCardRealTime() (err error) {
	idBoard := conf.IDBoard

	for {
		cardsOnTrelloDB, err := DBTrelloCard.GetMany("board:" + idBoard)
		if err != nil {
			return err
		}

		cardsOnDB, err := DBCard.GetAll()
		if err != nil {
			return err
		}

		err = DBCard.UpdateCardsChangedDueDate(cardsOnTrelloDB, cardsOnDB)
		if err != nil {
			return err
		}

	}
	return nil
}
