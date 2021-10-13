package hkidchecker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHKIDChecker(t *testing.T) {
	type testData struct {
		id       string
		expected bool
	}
	hkids := []testData{
		{"E364912(5)", true},
		{" E364912(5)", true},
		{"  E364912(5)", true},
		{"  E364912(5)   ", true},
		{"\n\tE364912(5)   \r", true},
		{"AB987654(3)", true},
		{"E364912(5)", true},
		{"E364912{5}", false},
		{"E364912(5", false},
		{"E364912(", false},
		{"E364912", false},
		{"E364912", false},
		{"AAB987654(3)", false},
	}
	for i, hkid := range hkids {
		assert.Equal(t, hkid.expected, CheckHKID(hkid.id), "Test %d: %s", i+1, hkid.id)
	}
}
