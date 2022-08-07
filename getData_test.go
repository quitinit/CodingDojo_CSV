package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetData(t *testing.T) {
	t.Run("pass in and get the whole thing out", func(t *testing.T) {
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}}
		header, data := GetData(rows, 1, 1)
		assert.Equal(t, [][]string{{"Peter", "42", "New York"}}, data)
		assert.Equal(t, []string{"Name", "Age", "City"}, header)
	})

	t.Run("pass in and get the whole thing out", func(t *testing.T) {
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}}
		header, data := GetData(rows, 1, 2)
		assert.Equal(t, [][]string{{"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}}, data)
		assert.Equal(t, []string{"Name", "Age", "City"}, header)
	})

	t.Run("pass in and get the whole thing out", func(t *testing.T) {
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}}
		header, data := GetData(rows, 2, 1)
		assert.Equal(t, [][]string{{"Michael", "26", "Berlin"}}, data)
		assert.Equal(t, []string{"Name", "Age", "City"}, header)
	})
	t.Run("pass in and get the whole thing out", func(t *testing.T) {
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}}
		header, data := GetData(rows, 2, 2)
		assert.Equal(t, [][]string{}, data)
		assert.Equal(t, []string{"Name", "Age", "City"}, header)
	})

	t.Run("pass in and get the whole thing out", func(t *testing.T) {
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}, {"Nikita", "34", "Munich"}}
		header, data := GetData(rows, 1, 4)
		assert.Equal(t, [][]string{{"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}, {"Nikita", "34", "Munich"}}, data)
		assert.Equal(t, []string{"Name", "Age", "City"}, header)
	})

}
