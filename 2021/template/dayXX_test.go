package year2021

import (
	file "advent-of-code/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayXXa(t *testing.T) {
	t.Run("TestInput", func(t *testing.T) {
		input := []string{"hello"}
		assert.Equal(t, 4512, dayXXa(input))
	})

	t.Run("RealInput", func(t *testing.T) {
		input := file.ReadPuzzleInputStrings("\n")
		t.Logf("")
		t.Logf("--------------------")
		t.Logf("Answer to Puzzle XXa")
		t.Logf("%d", dayXXa(input))
		t.Logf("--------------------")
	})
}

func TestDayXXb(t *testing.T) {
	t.Run("TestInput", func(t *testing.T) {
		input := []string{"hello"}
		assert.Equal(t, 1924, dayXXb(input))
	})

	t.Run("RealInput", func(t *testing.T) {
		input := file.ReadPuzzleInputStrings("\n")
		t.Logf("")
		t.Logf("--------------------")
		t.Logf("Answer to Puzzle XXb")
		t.Logf("%d", dayXXb(input))
		t.Logf("--------------------")
	})
}
