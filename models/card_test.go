package models

import (
	"testing"
	"time"

	"github.com/adlio/trello"
)

func Test_DBCard_InsertOrUpdate(t *testing.T) {
	var date = time.Now()
	var tests = []struct {
		input    Card
		expected error
	}{
		{
			input: Card{
				ID:                "4AmMG2on",
				Name:              "Demo 5 (5)[6]",
				DateLastChangeDue: &date,
			},
			expected: nil,
		},
	}

	var dbCard DBCard
	for _, test := range tests {
		output := dbCard.InsertOrUpdate(test.input)
		if output != test.expected {
			t.Error("Test failed:\n Input:", test.input, "\n Output:", output, "\n Expected:", test.expected)
		}
	}
}

func Test_DBCard_GetAll(t *testing.T) {
	var tests = []struct {
		input              error
		expectedCardNumber int
		expectedErr        error
	}{
		{
			input:              nil,
			expectedCardNumber: 13,
			expectedErr:        nil,
		},
	}

	var dbCard DBCard
	for _, test := range tests {
		output, err := dbCard.GetAll()
		if err != test.expectedErr {
			t.Error("Test failed:\n Input:", test.input, "\n Output:", output, "\n Output Number:", len(output), "\n Expected Error:", test.expectedErr, "\n Expected Data Number:", test.expectedCardNumber)
		} else if len(output) != test.expectedCardNumber {
			t.Error("Test failed:\n Input:", test.input, "\n Output:", output, "\n Output Number:", len(output), "\n Expected Error:", test.expectedErr, "\n Expected Data Number:", test.expectedCardNumber)
		}

	}
}

func Test_DBCard_SaveTrelloCards(t *testing.T) {
	var tests = []struct {
		input              error
		expectedCardNumber int
	}{
		{
			input:              nil,
			expectedCardNumber: 13,
		},
	}

	var dbCard DBCard
	var dbTrello DBTrelloCard
	for _, test := range tests {
		trelloCards, _ := dbTrello.GetMany("board:iCBtQXmr")
		dbCard.SaveTrelloCards(trelloCards)

		cards, _ := dbCard.GetAll()
		if len(cards) != test.expectedCardNumber {
			t.Error("Test failed:\n Input:", test.input, "\n Output Number:", len(cards), "\n Expected Data Number:", test.expectedCardNumber)
		}
	}
}

func Test_DBCard_GetCardsChangedDueByTime(t *testing.T) {
	var tests = []struct {
		input              int
		expectedCardNumber int
		expectedErr        error
	}{
		{
			input:              1,
			expectedCardNumber: 6,
			expectedErr:        nil,
		},
	}

	var dbCard DBCard
	for _, test := range tests {
		output := dbCard.GetCardsChangedDueByTime(test.input)
		if len(output) != test.expectedCardNumber {
			t.Error("Test failed:\n Input:", test.input, "\n Output:", output, "\n Output Number:", len(output), "\n Expected Data Number:", test.expectedCardNumber, "\n Expected Error:", test.expectedErr)
		}
	}
}

func Test_DBCard_UpdateCardsChangedDueDate(t *testing.T) {
	var dbCard DBCard
	var dbTrelloCard DBTrelloCard

	cards, _ := dbCard.GetAll()
	trelloCards, _ := dbTrelloCard.GetMany("board:iCBtQXmr")

	var tests = []struct {
		inputCard       []Card
		inputTrelloCard []*trello.Card
		expectedErr     error
	}{
		{
			inputCard:       cards,
			inputTrelloCard: trelloCards,
			expectedErr:     nil,
		},
	}

	for _, test := range tests {
		output := dbCard.UpdateCardsChangedDueDate(test.inputTrelloCard, test.inputCard)
		if output != test.expectedErr {
			t.Error("Test failed:\n Input Cards:", test.inputCard, "\n Input Trello Cards:", test.inputTrelloCard, "\n Output Error:", output, "\n Expected Error:", test.expectedErr)
		}
	}
}
