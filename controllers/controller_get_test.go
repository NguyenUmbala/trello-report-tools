package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"trello-report-tools/models"
	"trello-report-tools/routers"

	"github.com/adlio/trello"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router = routers.SetupRouters()

// Count number of cards on response
func Test_GetAllCardReview(t *testing.T) {
	tests := []struct {
		input    string
		expected gin.H
	}{
		{
			input: "/api/review",
			expected: gin.H{
				"Number of cards on review-me": 0,
				"Number of cards on Done":      1,
			},
		},
	}

	for _, test := range tests {

		w := performRequest(router, "GET", test.input)
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string][]*trello.Card
		err := json.Unmarshal([]byte(w.Body.String()), &response)
		assert.Nil(t, err)

		listCardsOnReviewme, exists := response["List card on review-me"]
		assert.True(t, exists)
		assert.Equal(t, test.expected["Number of cards on review-me"], len(listCardsOnReviewme))

		listCardsOnDone, exists := response["List card on Done"]
		assert.True(t, exists)
		assert.Equal(t, test.expected["Number of cards on Done"], len(listCardsOnDone))
	}
}

func Test_GetAllCardChangeDue(t *testing.T) {
	tests := []struct {
		input    string
		expected gin.H
	}{
		{
			input: "/api/changedue",
			expected: gin.H{
				"Number of cards changed due date": 1,
			},
		},
	}

	for _, test := range tests {

		w := performRequest(router, "GET", test.input)
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string][]models.Card
		err := json.Unmarshal([]byte(w.Body.String()), &response)
		assert.Nil(t, err)

		listCardsChangedDueDate, exists := response["Cards changed due date"]
		assert.True(t, exists)
		assert.Equal(t, test.expected["Number of cards changed due date"], len(listCardsChangedDueDate))
	}
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
