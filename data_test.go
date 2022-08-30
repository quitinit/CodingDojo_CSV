package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetData(t *testing.T) {
	/* t.Run("rows gets passed by reference and not by value", func(t *testing.T) {
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}}
		data := GetData(rows, 1, 1)
		data.body = append(data.body, []string{"some", "other", "name"})
		if &rows != &data.body {
			t.Errorf("%v expected but got %v", [][]string{{"Peter", "42", "New York"}}, data.body)
		}
	}) */

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
	//data := GetData(rows, 1, 4)

	t.Run("sorting alphabetically with a cortrectly given header name", func(t *testing.T) {
		data := GetData(rows, 1, 4)
		_ = data.Sort("City")
		assert.Equal(t, [][]string{{"Michael", "26", "Berlin"}, {"Peter", "42", "New York"}}, data.body)
	})
	t.Run("sorting alphabetically with a incorrectly given header name", func(t *testing.T) {
		data := GetData(rows, 1, 4)
		err := data.Sort("something")
		assert.EqualError(t, err, "wrong header passed")
		assert.Equal(t, [][]string{{"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}}, data.body)
	})
}
