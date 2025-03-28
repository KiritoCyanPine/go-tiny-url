package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatString(t *testing.T) {
	s, err := formatUrlHash("qwertyuiop")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, s, "qwe-rty-uiop")
}
