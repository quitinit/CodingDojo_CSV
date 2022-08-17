package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpToPage(t *testing.T) {
	t.Run("jump to maxpage", func(t *testing.T) {
		var maxPage uint = 2
		got, _ := JumpToPage(2, maxPage)
		want := 2
		assert.Equal(t, want, got)
	})
	t.Run("Jump to page in the middle", func(t *testing.T) {
		var maxPage uint = 10
		got, _ := JumpToPage(5, maxPage)
		want := 5
		assert.Equal(t, want, got)
	})
	t.Run("Jump beyond the maxpage", func(t *testing.T) {
		var maxPage uint = 10
		_, err := JumpToPage(11, maxPage)
		//want := 5
		assert.EqualError(t, err, "jumped beyond the page limit")
	})

}

func TestGetJumpPageInput(t *testing.T) {
	t.Run("Recieves the right input", func(t *testing.T) {

	})
	//fmt.P

}
