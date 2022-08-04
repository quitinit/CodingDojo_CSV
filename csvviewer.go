package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
		argument one is the file
		argument two is page length

		rules for myself --> everything that needs to change needs to go inside config file
		SOLID Priciples


		controls:
		F)irst page, P)revious page, N)ext page, L)ast page, E)xit


	*/

	_, err := ParseCommandlineArgs(os.Args[1:])

	if err != nil {
		//TODO have some logging in here
		os.Exit(1)
	}
	// I do this so that the file is already in memory and does not need to be loaded from disc again
	//content := ReadFile(config.Filename)

	input := bufio.NewScanner(os.Stdin)
	fmt.Println("F)irst page, P)revious page, N)ext page, L)ast page, E)xit")

	// read in the data

	// render the data

	for {
		input.Scan()
		switch strings.ToLower(input.Text()) {
		case "e":
			os.Exit(0)

		}

	}

}
