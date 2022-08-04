package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Filename string
	Pagesize int
}

func parseCommandlineArgs(args []string) (c *Config, err error) {
	filename := args[0]
	if _, fileErr := os.Stat(filename); fileErr != nil {
		return nil, fileErr
	}
	pagesize, conversionErr := strconv.Atoi(args[1])
	if conversionErr != nil {
		return nil, errors.New("page size needs to be an integer")
	}
	if pagesize < 1 {
		return nil, errors.New("page size needs to be a positive integer")
	}

	return &Config{Filename: filename, Pagesize: pagesize}, nil
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
	fmt.Println("F)irst page, P)revious page, N)ext page, L)ast page, E)xit")

	input := bufio.NewScanner(os.Stdin)
	_, err := parseCommandlineArgs(os.Args[1:])
	if err != nil {
		//TODO have some logging in here
		os.Exit(1)
	}

	for {
		input.Scan()
		switch strings.ToLower(input.Text()) {
		case "e":
			os.Exit(0)
		}

	}

}
