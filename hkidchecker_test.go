package hkidchecker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHKIDChecker(t *testing.T) {
	type testData struct {
		id        string
		expectedA bool
		expectedB bool
	}
	hkids := []testData{
		{"E364912(5)", true, true},
		{" E364912(5)", true, true},
		{"  E364912(5)", true, true},
		{"  E364912(5)   ", true, true},
		{"\n\tE364912(5)   \r", true, true},
		{"AB987654(3)", true, true},
		{"E364912(5)", true, true},
		{"E364912(6)", true, false},
		{"E364912{5}", false, false},
		{"E364912(5", false, false},
		{"E364912(", false, false},
		{"E364912", false, false},
		{"E364912", false, false},
		{"AAB987654(3)", false, false},
	}
	for i, hkid := range hkids {
		assert.Equal(t, hkid.expectedA, CheckHKIDFormat(hkid.id), "Test %dA: %s", i+1, hkid.id)
		assert.Equal(t, hkid.expectedB, CheckHKID(hkid.id), "Test %dB: %s", i+1, hkid.id)
	}
}
