package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func GetData(input_data [][]string, page int, step int) (header []string, data [][]string) {
	/*
		This function takes the entire dataset and returns a sliced up part for the render

	*/
	// it makes no sense the way I do it here, I should have a page devider function

	if len(input_data) == 0 {
		return nil, nil
	}

	if len(input_data) == 1 {
		return input_data[0], nil
	}
	lower_limit := (page-1)*step + 1
	upper_limit := page * step
	if upper_limit > len(input_data)-1 {
		upper_limit = len(input_data) - 1
	}
	header, rest := input_data[0], input_data[lower_limit:upper_limit+1]
	return header, rest

}

func main() {
	/*
		argument one is the file
		argument two is page length

		rules for myself --> everything that needs to change needs to go inside config file
		SOLID Priciples


		controls:
		F)irst page, P)revious page, N)ext page, L)ast page, E)xit


	*/

	config, err := ParseCommandlineArgs(os.Args[1:])

	if err != nil {
		//TODO have some logging in here
		os.Exit(1)
	}
	// I do this so that the file is already in memory and does not need to be loaded from disc again
	content := ReadFile(config.Filename)
	// transform the data so it only shows the
	input := bufio.NewScanner(os.Stdin)

	page := 1
	maxPage := (len(content) - 1) / config.Pagesize
	writer := os.Stdout
	for {
		CompleteRender(writer, content, page, maxPage, config)
		input.Scan()
		switch strings.ToLower(input.Text()) {
		case "e":
			os.Exit(0)
		case "f":
			page = 1

		case "p":
			page--
			if page < 1 {
				page = 1
			}

		case "n":
			page++
			if page > maxPage {
				page = maxPage
			}
		case "l":
			page = maxPage

		case "j":
			fmt.Fprintln(writer, "Page Number:")
			input.Scan()
			jumpPage, err := strconv.Atoi(input.Text())
			if err != nil {
				fmt.Fprintf(writer, "Not a page number %v %v", jumpPage, err)
				break
			}
			jumpPage, jumpErr := JumpToPage(jumpPage, maxPage)
			if jumpErr != nil {
				fmt.Fprintln(writer, jumpErr)
				break
			}
			page = jumpPage
		}

	}

}

func CompleteRender(writer io.Writer, content [][]string, page int, maxPage int, config *Config) {
	header, currentContent := GetData(content, page, config.Pagesize)
	Render(writer, header, currentContent)
	fmt.Fprintf(writer, "Page %d/%d\n", page, maxPage)
	fmt.Fprintln(writer, "F)irst page, P)revious page, N)ext page, L)ast page, J) Jump to Page, E)xit")

}
