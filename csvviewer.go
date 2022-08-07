package main

import (
	"bufio"
	"fmt"
	"os"
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
	maxPage := len(content) - 1/config.Pagesize
	header, currentContent := GetData(content, page, config.Pagesize)

	Render(os.Stdout, header, currentContent)
	fmt.Println("F)irst page, P)revious page, N)ext page, L)ast page, E)xit")

	// read in the data

	// render the data

	for {
		input.Scan()
		switch strings.ToLower(input.Text()) {
		case "e":
			os.Exit(0)
		case "f":
			page = 1
			header, currentContent := GetData(content, page, config.Pagesize)
			Render(os.Stdout, header, currentContent)
		case "p":
			page--
			if page < 1 {
				page = 1
			}
			header, currentContent := GetData(content, page, config.Pagesize)
			Render(os.Stdout, header, currentContent)
		case "n":
			page++
			if page > maxPage {
				page = maxPage
			}
			header, currentContent := GetData(content, page, config.Pagesize)
			Render(os.Stdout, header, currentContent)
		case "l":
			page = maxPage
			header, currentContent := GetData(content, page, config.Pagesize)
			Render(os.Stdout, header, currentContent)
		}

	}

}
