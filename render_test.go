package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CheckRender(t *testing.T, sample_data [][]string, start int, limit int, expected string) {
	t.Helper()
	writer := bytes.Buffer{}
	// the render function takes csv da
	header, data := GetData(sample_data, start, limit)
	Render(&writer, header, data)
	expected = strings.TrimLeft(expected, "\n")
	expected = strings.TrimLeft(expected, "\t")
	got := writer.String()
	assert.Equal(t, expected, got)

}
func TestRender(t *testing.T) {

	t.Run("Render only the header", func(t *testing.T) {
		rowLimit := 1
		rows := [][]string{{"Name", "Age", "City"}}
		// the render function takes csv da
		expected := `
Name|Age|City|
----+---+----+
`
		CheckRender(t, rows, 1, rowLimit, expected)

	})
	t.Run("Pass in two rows and render two rows", func(t *testing.T) {
		rowLimit := 2
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}}
		// the render function takes csv da
		expected := `
Name |Age|City    |
-----+---+--------+
Peter|42 |New York|
`
		CheckRender(t, rows, 1, rowLimit, expected)
	})
	t.Run("Pass in two rows and render one rows", func(t *testing.T) {
		rowLimit := 1
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}}
		// the render function takes csv da
		expected := `
Name |Age|City    |
-----+---+--------+
Peter|42 |New York|
`
		CheckRender(t, rows, 1, rowLimit, expected)
	})
	t.Run("Pass in two rows and render one rows", func(t *testing.T) {
		rowLimit := 1
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}}
		// the render function takes csv da
		expected := `
Name |Age|City    |
-----+---+--------+
Peter|42 |New York|
`
		CheckRender(t, rows, 1, rowLimit, expected)
	})
	t.Run("Pass in one row and render one row", func(t *testing.T) {
		rowLimit := 2
		rows := [][]string{{"Name", "Age", "City"}}
		// the render function takes csv da
		expected := `
Name|Age|City|
----+---+----+
`
		CheckRender(t, rows, 1, rowLimit, expected)
	})
	t.Run("Nothing gets printed empty data and higher row limit", func(t *testing.T) {
		rowLimit := 2
		rows := [][]string{}
		// the render function takes csv da
		expected := ``
		CheckRender(t, rows, 1, rowLimit, expected)
	})
	t.Run("Pass in two rows and render one row on page two", func(t *testing.T) {
		rowStart := 2
		rowLimit := 1
		rows := [][]string{{"Name", "Age", "City"}, {"Peter", "42", "New York"}, {"Michael", "26", "Berlin"}}
		// the render function takes csv da
		expected := `
Name   |Age|City  |
-------+---+------+
Michael|26 |Berlin|
`

		CheckRender(t, rows, rowStart, rowLimit, expected)
	})

	//TODO make a test for trailing space but that goes inside end to end tests
	//TODO make a test for empty data
}
