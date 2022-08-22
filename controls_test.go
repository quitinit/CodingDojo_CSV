package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpToPage(t *testing.T) {
	t.Run("jump to maxpage", func(t *testing.T) {
		//var maxPage int = 2
		state := State{MaxPage: 2}
		state.maxPage()
		//got, _ := JumpToPage(2, maxPage)
		want := 2
		assert.Equal(t, want, state.Page)
	})
	t.Run("Jump to page in the middle", func(t *testing.T) {
		//var maxPage int = 10
		state := State{MaxPage: 10}
		state.setPage(5)
		//got, _ := JumpToPage(5, maxPage)
		want := 5
		assert.Equal(t, want, state.Page)
	})
	t.Run("Jump beyond the maxpage", func(t *testing.T) {
		//var maxPage int = 10
		state := State{MaxPage: 5}
		err := state.setPage(10)
		//got, _ := JumpToPage(5, maxPage)

		assert.EqualError(t, err, "set beyond page limit")
	})

}
