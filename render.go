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
func Render(writer *bytes.Buffer, file [][]string, rowLimit int) {
	var frame string
	// how not to do this in n^2
	max_elements := GetFrameLengths(file)
	for i := 0; i < rowLimit; i++ {
		for j := 0; j < len(file[i]); j++ {
			var seperatingChar string
			if j != len(file[i])-1 {
				seperatingChar = "|"
			} else {
				//terminates the line
				seperatingChar = "|\n"
			}

			fmt.Fprint(writer, file[i][j]+strings.Repeat(" ", max_elements[j]-len(file[i][j]))+seperatingChar)
			frame += strings.Repeat("-", max_elements[j])
			frame += "+"
		}
		if i == 0 {
			fmt.Fprintln(writer, frame)
		}

	}

}
