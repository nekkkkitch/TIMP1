package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	goodInput = "2000.10.10 12:30 \"Vasiliy Viktor\""

	badDate = "0001.13.48 12:30 \"Vasiliy Viktor\""
	noDate  = "200 12:30 \"Vasiliy Viktor\""

	badTime = "2000.10.10 12:61 \"Vasiliy Viktor\""
	noTime  = "2000.10.10 14312 \"Vasiliy Viktor\""

	noName  = "2000.10.10 12:30 \"Vasiliy "
	badName = "2000.10.10 12:30 \"Vasiliy123 bbb\""
)

func TestParser(t *testing.T) {
	t.Log("Good input test")
	_, err := ProcessInput(goodInput)

	assert.Nil(t, err)

	t.Log("Bad date check")
	_, err = ProcessInput(badDate)
	require.ErrorIs(t, err, ErrBadDate)
	_, err = ProcessInput(noDate)
	require.ErrorIs(t, err, ErrNoDate)

	t.Log("Bad time check")
	_, err = ProcessInput(badTime)
	require.ErrorIs(t, err, ErrBadTime)
	_, err = ProcessInput(noTime)
	require.ErrorIs(t, err, ErrNoTime)

	t.Log("Bad name check")
	_, err = ProcessInput(noName)
	require.ErrorIs(t, err, ErrNoName)
	_, err = ProcessInput(badName)
	require.ErrorIs(t, err, ErrNoName)
}
