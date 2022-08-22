package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Command interface {
	execute()
}
type State struct {
	Page    int
	MaxPage int
}

func (state *State) New(maxPage int) {
	state.Page = 1
	state.MaxPage = maxPage

}
func (state *State) reset() {
	state.Page = 1
}

func (state *State) upPage() {
	if state.Page+1 > state.MaxPage {
		state.Page = state.MaxPage
	} else {
		state.Page++
	}

}
func (state *State) downPage() {
	if state.Page < 1 {
		state.Page = 1
	} else {
		state.Page--
	}

}
func (state *State) maxPage() {
	state.Page = state.MaxPage
}
func (state *State) setPage(page int) error {
	if page > state.MaxPage || page < 1 {
		return errors.New("set beyond page limit")
	}
	state.Page = page
	return nil
}

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

	state := State{MaxPage: (len(content) - 1) / config.Pagesize}

	writer := os.Stdout
	for {
		CompleteRender(writer, content, state.Page, state.MaxPage, config)
		input.Scan()
		switch strings.ToLower(input.Text()) {
		// inputs a letter -> does something in between -> outputs a page and perhaps renders something more to the screen

		case "e":
			os.Exit(0)
		case "f":
			state.reset()

		case "p":
			state.downPage()

		case "n":
			state.upPage()

		case "l":
			state.maxPage()

		case "j":
			fmt.Fprintln(writer, "Page Number:")
			input.Scan()
			jumpPage, err := strconv.Atoi(input.Text())
			if err != nil {
				fmt.Fprintf(writer, "Not a page number %v %v", jumpPage, err)
				break
			}
			jumpPage, jumpErr := JumpToPage(jumpPage, state.MaxPage)
			if jumpErr != nil {
				fmt.Fprintln(writer, jumpErr)
				break
			}
			state.setPage(jumpPage)
			//page = jumpPage
		}

	}

}

func CompleteRender(writer io.Writer, content [][]string, page int, maxPage int, config *Config) {
	header, currentContent := GetData(content, page, config.Pagesize)
	Render(writer, header, currentContent)
	fmt.Fprintf(writer, "Page %d/%d\n", page, maxPage)
	fmt.Fprintln(writer, "F)irst page, P)revious page, N)ext page, L)ast page, J) Jump to Page, E)xit")

}
