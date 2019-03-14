package models

import (
	"testing"
)

func Test_DBTrelloCard_GetMany(t *testing.T) {
	var tests = []struct {
		input               string
		expectedCardsNumber int
		expectedErr         error
	}{
		{
			input:               "board:iCBtQXmr",
			expectedCardsNumber: 10,
			expectedErr:         nil,
		},
	}

	var dbTrello DBTrelloCard
	for _, test := range tests {
		output, err := dbTrello.GetMany(test.input)

		if err != test.expectedErr {
			t.Error("Test failed:\n Input:", test.input, "\n Output:", output, "\n Output Number:", len(output), "\n Expected Data Number:", test.expectedCardsNumber, "\n Expected Error:", test.expectedErr)
		} else if len(output) != test.expectedCardsNumber {
			t.Error("Test failed:\n Input:", test.input, "\n Output:", output, "\n Output Number:", len(output), "\n Expected Data Number:", test.expectedCardsNumber, "\n Expected Error:", test.expectedErr)
		}
	}
}
