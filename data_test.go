package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetData(t *testing.T) {

	t.Run("pass in and get the whole thing out", func(t *testing.T) {
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}}
		input_data := NewData(rows)
		data := input_data.SliceData(1, 1)
		assert.Equal(t, [][]string{{"Peter", "42", "New York"}}, data.body)
		assert.Equal(t, []string{"Name", "Age", "City"}, data.header)
	})

	t.Run("pass in and get the whole thing out", func(t *testing.T) {
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}}
		input_data := NewData(rows)
		data := input_data.SliceData(1, 2)
		assert.Equal(t, [][]string{{"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}}, data.body)
		assert.Equal(t, []string{"Name", "Age", "City"}, data.header)
	})

	t.Run("pass in and get the whole thing out", func(t *testing.T) {
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}}
		input_data := NewData(rows)
		data := input_data.SliceData(2, 1)
		assert.Equal(t, [][]string{{"Michael", "26", "Berlin"}}, data.body)
		assert.Equal(t, []string{"Name", "Age", "City"}, data.header)
	})
	t.Run("pass in and get the whole thing out", func(t *testing.T) {
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}}
		input_data := NewData(rows)
		data := input_data.SliceData(2, 2)
		assert.Equal(t, [][]string{}, data.body)
		assert.Equal(t, []string{"Name", "Age", "City"}, data.header)
	})

	t.Run("pass in and get the whole thing out", func(t *testing.T) {
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}, {"Nikita", "34", "Munich"}}
		input_data := NewData(rows)
		data := input_data.SliceData(1, 4)
		assert.Equal(t, [][]string{{"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}, {"Nikita", "34", "Munich"}}, data.body)
		assert.Equal(t, []string{"Name", "Age", "City"}, data.header)
	})

}

func TestSort(t *testing.T) {
	rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}}
	//data := GetData(rows, 1, 4)
	data := NewData(rows)
	t.Run("sorting alphabetically with a cortrectly given header name", func(t *testing.T) {

		_ = data.Sort("City")
		assert.Equal(t, [][]string{{"Michael", "26", "Berlin"}, {"Peter", "42", "New York"}}, data.body)
	})
	t.Run("sorting alphabetically with a incorrectly given header name", func(t *testing.T) {
		prev_state_data := data
		err := data.Sort("something")
		assert.EqualError(t, err, "wrong header passed")
		assert.Equal(t, prev_state_data.body, data.body)
	})
	t.Run("sorting data by number", func(t *testing.T) {

		err := data.Sort("Age")
		if err != nil {
			t.Fatalf("panicked on sorting by age")
		}
		assert.Equal(t, [][]string{{"Michael", "26", "Berlin"}, {"Peter", "42", "New York"}}, data.body)
	})

}

func AcceptanceTestsSorting(t *testing.T) {
	content := ReadFile("./sample_files/example.csv")
	input_data := NewData(content)
	data := input_data.SliceData(1, 3)
	t.Run("sorting on name field", func(t *testing.T) {
		err := data.Sort("Name")
		if err != nil {
			assert.Equal(t, [][]string{{"Mary", "35", "Munich"}, {"Paul", "57", "London"}, {"Peter", "42", "New York"}}, data)
		}
	})

}
