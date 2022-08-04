package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommandline(t *testing.T) {
	t.Run("correct config with page size 3 and relative file", func(t *testing.T) {
		args := []string{"./sample_files/example.csv", "3"}
		got, _ := ParseCommandlineArgs(args)
		assert.Equal(t, &Config{"./sample_files/example.csv", 3}, got)
	})
	t.Run("parser fails on negative page size", func(t *testing.T) {
		args := []string{"./sample_files/example.csv", "-3"}
		_, err := ParseCommandlineArgs(args)
		assert.EqualErrorf(t, err, "page size needs to be a positive integer", "args parser let negative number through")
	})
	t.Run("parser fails on wrong format page size", func(t *testing.T) {
		args := []string{"./sample_files/example.csv", "something"}
		_, err := ParseCommandlineArgs(args)
		assert.EqualErrorf(t, err, "page size needs to be an integer", "args parser argument fail")
	})
	t.Run("parser fails on wrong non existing file", func(t *testing.T) {
		args := []string{"./sample_files/notexistent.csv", "something"}
		_, err := ParseCommandlineArgs(args)
		assert.ErrorContains(t, err, "The system cannot find the file specified.")
	})

}
