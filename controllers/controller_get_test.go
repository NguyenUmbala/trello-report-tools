package controllers_test

import (
	"TrelloReportTools/routers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllCardReview(t *testing.T) {
	// var tests = []struct {
	// 	input    *gin.Context
	// 	expected error
	// }{
	// 	{
	// 		input:    &c,
	// 		expected: nil,
	// 	},
	// }

	// for _, test := range tests {
	// 	output := GetAllCardReview(test.input)
	// 	if output != test.expected {
	// 		t.Error("Test failed:\n Input:", test.input, "\n Output:", output, "\n Expected:", test.expected)
	// 	}
	// }

	body := gin.H{
		"List card on review-me": "world",
		"List card on Done":      "as",
	}

	router := routers.SetupRouters()

	w := performRequest(router, "GET", "/b/cards/review")
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	listCardsOnReviewme, exists := response["List card on review-me"]
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["List card on review-me"], listCardsOnReviewme)

	listCardsOnDone, exists := response["List card on Done"]
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["List card on Done"], listCardsOnDone)
}

// func Test_GetAllCardChangeDue(t *testing.T) {
// 	var c gin.Context
// 	var tests = []struct {
// 		input    *gin.Context
// 		expected error
// 	}{
// 		{
// 			input:    &c,
// 			expected: nil,
// 		},
// 	}

// 	for _, test := range tests {
// 		output := GetAllCardChangeDue(test.input)
// 		if output != test.expected {
// 			t.Error("Test failed:\n Input:", test.input, "\n Output:", output, "\n Expected:", test.expected)
// 		}
// 	}
// }
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
