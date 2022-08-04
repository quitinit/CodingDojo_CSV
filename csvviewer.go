package main

import (
	"bufio"
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
	if _, fileErr := os.Stat(filename); fileErr == nil {
		return nil, fileErr
	}
	pagesize, conversionErr := strconv.Atoi(args[1])
	if conversionErr != nil {
		return nil, err
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

	for {
		input.Scan()
		if strings.ToLower(input.Text()) == "e" {
			os.Exit(0)
		}
	}

}
