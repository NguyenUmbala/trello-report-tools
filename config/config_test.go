package config

import (
	"testing"
)

func TestSetup(t *testing.T) {
	var tests = []struct {
		expected Config
	}{
		{Config{"fa6b1a601cfd6559cc134d0055507cc2", "e0a44f959cde9dfb0883e2865d5632232b0b3ac93900263d22d6e7f84a1d0017", "iCBtQXmr"}},
	}
	var output = Setup()
	for _, test := range tests {
		if output.IDBoard != test.expected.IDBoard || output.KeyApp != test.expected.KeyApp || output.Token != test.expected.Token {
			t.Error("Test failed: Output: {}, Expected: {}", output)
		} else {

		}
	}
}
