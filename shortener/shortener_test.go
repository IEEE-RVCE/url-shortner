package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortLinkGenerator(t *testing.T) {
	initialLink := "https://go.dev/doc/tutorial/getting-started"
	userId := "UwQPr3aIf9dM5x7r"
	trialText := "Trial_Text"
	shortLink := GenerateShortURL(initialLink, userId, trialText)

	assert.Equal(t, shortLink, "dysg5Fas")
}
