package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CheckRender(t *testing.T, sample_data [][]string, limit int, expected string) {
	t.Helper()
	writer := bytes.Buffer{}
	// the render function takes csv da
	Render(&writer, sample_data, limit)
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
		CheckRender(t, rows, rowLimit, expected)

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
		CheckRender(t, rows, rowLimit, expected)
	})

	//TODO make a test for trailing space but that goes inside end to end tests
}
