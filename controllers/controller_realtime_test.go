package controllers

import (
	"testing"
	"time"
)

func Test_UpdateCardRealTime(t *testing.T) {
	var tests = []struct {
		input    error
		expected error
	}{
		{nil, nil},
	}

	for _, test := range tests {
		var output error
		go func(err *error) {
			*err = UpdateCardRealTime()
		}(&output)

		time.Sleep(5 * time.Second)
		if output != test.expected {
			t.Error("Test failed:\n Input:", test.input, "\n Output:", output, "\n Expected:", test.expected)
		}
	}
}
