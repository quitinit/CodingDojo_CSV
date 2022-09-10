package main

import (
	"bufio"
	"errors"
	"fmt"
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

// builder constructor
func NewState(maxPage int) *State {
	state := State{}
	state.Page = 1
	state.MaxPage = maxPage
	return &state
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
	if state.Page <= 1 {
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

	state := NewState((len(content) - 1) / config.Pagesize)

	writer := os.Stdout

	//data := GetData(content, state.Page, config.Pagesize)

	input_data := NewData(content)

	for {
		data := input_data.SliceData(state.Page, config.Pagesize)
		Render(writer, state, &data)
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
		case "s":
			fmt.Fprintln(writer, "Column name:")
			input.Scan()
			column_name := input.Text()
			input_data.Sort(column_name)
		case "j":
			fmt.Fprintln(writer, "Page Number:")
			input.Scan()
			jumpPage, err := strconv.Atoi(input.Text())
			if err != nil {
				fmt.Fprintf(writer, "Not a page number %v %v", jumpPage, err)
				break
			}
			jumpError := state.setPage(jumpPage)
			if jumpError != nil {
				fmt.Fprintln(writer, jumpError.Error())
			}

			//page = jumpPage
		}

	}

}
