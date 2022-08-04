package main

import (
	"bytes"
	"fmt"
	"strings"
)

func GetFrameLengths(rows [][]string) []int {
	length := len(rows[0])
	max_elements := make([]int, length)
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {

			if len(rows[i][j]) > max_elements[j] {
				max_elements[j] = len(rows[i][j])

			}
			temp := max_elements[j]
			fmt.Println(temp)
		}
	}
	return max_elements
}
func Render(writer *bytes.Buffer, rows [][]string, rowLimit int) {
	var frame string
	// how not to do this in n^2
	max_elements := GetFrameLengths(rows)
	for i := 0; i < rowLimit; i++ {
		for j := 0; j < len(rows[i]); j++ {
			var seperatingChar string
			if j != len(rows[i])-1 {
				seperatingChar = "|"
			} else {
				//terminates the line
				seperatingChar = "|\n"
			}
			entry := rows[i][j]
			fmt.Fprint(writer, entry+strings.Repeat(" ", max_elements[j]-len(entry))+seperatingChar)
			frame += strings.Repeat("-", max_elements[j])
			frame += "+"
		}
		if i == 0 {
			fmt.Fprintln(writer, frame)
		}

	}

}
