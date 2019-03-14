package utils

import (
	"testing"
	"time"
)

func Test_Utils_CompareTime(t *testing.T) {
	time1, _ := time.Parse("2019-04-20 05:00:00+00:00", "2019-04-20 05:00:00+00:00")
	time2, _ := time.Parse("2019-04-20 05:00:00+00:00", "2019-04-20 05:00:00+00:00")

	var tests = []struct {
		input1   *time.Time
		input2   *time.Time
		expected bool
	}{
		{
			input1:   &time1,
			input2:   &time2,
			expected: true,
		},
	}

	var util Utils
	for _, test := range tests {
		output := util.CompareTime(test.input1, test.input2)
		if output != test.expected {
			t.Error("Test faled:\n Input 1:", test.input1, "\n Input 2:", test.input2, "\n Output:", output, "\n Expected:", test.expected)
		}

	}
}

func Test_Utils_GetTimeRealForDone(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{
			input:    "test [3](4)",
			expected: 3,
		},
	}

	for _, test := range tests {
		output := new(Utils).GetTimeRealForDone(test.input)
		if output != test.expected {
			t.Error("Test faled:\n Input:", test.input, "\n Output:", output, "\n Expected:", test.expected)
		}
	}
}

func Test_Utils_GetTimeGuessForDone(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{
			input:    "test [3](4)",
			expected: 4,
		},
	}

	for _, test := range tests {
		output := new(Utils).GetTimeGuessForDone(test.input)
		if output != test.expected {
			t.Error("Test faled:\n Input:", test.input, "\n Output:", output, "\n Expected:", test.expected)
		}
	}
}
