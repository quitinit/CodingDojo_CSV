package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetData(t *testing.T) {
	t.Run("pass in and get the whole thing out", func(t *testing.T) {
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}}
		data := GetData(rows, 1, 1)
		assert.Equal(t, [][]string{{"Peter", "42", "New York"}}, data.body)
		assert.Equal(t, []string{"Name", "Age", "City"}, data.header)
	})

	t.Run("pass in and get the whole thing out", func(t *testing.T) {
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}}
		data := GetData(rows, 1, 2)
		assert.Equal(t, [][]string{{"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}}, data.body)
		assert.Equal(t, []string{"Name", "Age", "City"}, data.header)
	})

	t.Run("pass in and get the whole thing out", func(t *testing.T) {
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}}
		data := GetData(rows, 2, 1)
		assert.Equal(t, [][]string{{"Michael", "26", "Berlin"}}, data.body)
		assert.Equal(t, []string{"Name", "Age", "City"}, data.header)
	})
	t.Run("pass in and get the whole thing out", func(t *testing.T) {
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}}
		data := GetData(rows, 2, 2)
		assert.Equal(t, [][]string{}, data.body)
		assert.Equal(t, []string{"Name", "Age", "City"}, data.header)
	})

	t.Run("pass in and get the whole thing out", func(t *testing.T) {
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}, {"Nikita", "34", "Munich"}}
		data := GetData(rows, 1, 4)
		assert.Equal(t, [][]string{{"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}, {"Nikita", "34", "Munich"}}, data.body)
		assert.Equal(t, []string{"Name", "Age", "City"}, data.header)
	})

}

func TestSort(t *testing.T) {
	rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}}
	data := GetData(rows, 2, 2)

	t.Run("sorting alphabetically with a cortrectly given header name", func(t *testing.T) {
		data.Sort("City")
		assert.Equal(t, [][]string{{"Michael", "26", "Berlin"}, {"Peter", "42", "New York"}}, data.body)
	})
}
