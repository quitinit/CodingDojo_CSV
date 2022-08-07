package main

import (
	"fmt"
	"io"
	"strings"
)

func GetFrameLengths(rows [][]string) []int {
	if len(rows) == 0 {
		max_elements := make([]int, 0)
		return max_elements
	}
	length := len(rows[0])
	startRows := 0
	endRows := len(rows)

	max_elements := make([]int, length)
	for i := startRows; i < endRows; i++ {
		for j := 0; j < len(rows[i]); j++ {

			if len(rows[i][j]) > max_elements[j] {
				max_elements[j] = len(rows[i][j])

			}
		}
	}
	return max_elements
}
func Render(writer io.Writer, header []string, rows [][]string) {
	var frame string
	headerArray := make([][]string, 0)
	// how not to do this in n^2
	if len(header) != 0 {
		headerArray = [][]string{header}
	}
	full_data := append(headerArray, rows...)
	max_elements := GetFrameLengths(full_data)

	for i := 0; i < len(full_data); i++ {
		for j := 0; j < len(full_data[i]); j++ {
			var seperatingChar string
			if j != len(full_data[i])-1 {
				seperatingChar = "|"
			} else {
				//terminates the line
				seperatingChar = "|\n"
			}
			entry := full_data[i][j]
			fmt.Fprint(writer, entry+strings.Repeat(" ", max_elements[j]-len(entry))+seperatingChar)
			frame += strings.Repeat("-", max_elements[j])
			frame += "+"
		}
		if i == 0 {
			fmt.Fprintln(writer, frame)
		}

	}

}
