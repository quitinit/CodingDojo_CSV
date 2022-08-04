package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {

	t.Run("Render only the header", func(t *testing.T) {
		writer := bytes.Buffer{}
		rowLimit := 1
		rows := [][]string{{"Name", "Age", "City"}}
		// the render function takes csv da
		Render(&writer, rows, rowLimit)
		expected := `
Name|Age|City|
----+---+----+
`
		expected = strings.TrimLeft(expected, "\n")
		expected = strings.TrimLeft(expected, "\t")
		got := writer.String()
		assert.Equal(t, expected, got)

	})
	//TODO make a test for trailing space
}
