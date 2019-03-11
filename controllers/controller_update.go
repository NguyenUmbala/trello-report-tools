package controllers

func init() {
	ServiceUpdate.SaveAllCards(conf.IDBoard)
}

type UpdateCard struct{}

func (*UpdateCard) UpdateCardRealTime() {
	idBoard := conf.IDBoard

	for {
		ServiceUpdate.UpdateCards(idBoard)
	}
}
