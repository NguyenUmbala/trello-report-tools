package models

import (
	"testing"

	"github.com/jinzhu/gorm"
)

func Test_openDB(t *testing.T) {
	var tests = []struct {
		input    *gorm.DB
		expected error
	}{
		{
			input:    db,
			expected: nil,
		},
	}

	for _, test := range tests {
		output := tests[0].expected
		if output != test.expected {
			t.Error("Test failed:\n Input:", test.input, "\n Output:", output, "\n Expected:", test.expected)
		}
	}

}
