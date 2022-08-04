package main

import (
	"bytes"
	"fmt"
	"strings"
)

func Render(writer *bytes.Buffer, file [][]string, rowLimit int) {
	var frame string
	for i := 0; i < rowLimit; i++ {
		for j := 0; j < len(file[i]); j++ {
			var seperatingChar string
			if j != len(file[i])-1 {
				seperatingChar = "|"
			} else {
				//terminates the line
				seperatingChar = "|\n"
			}

			fmt.Fprint(writer, file[i][j]+seperatingChar)
			frame += strings.Repeat("-", len(file[i][j]))
			frame += "+"
		}
		fmt.Fprintln(writer, frame)

	}
}
