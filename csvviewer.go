package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
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
func ReadFile(path string) [][]string {
	var result [][]string
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		r := csv.NewReader(strings.NewReader(line))

		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			result = append(result, record)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	/* --


	 */

	return result
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

	_, err := parseCommandlineArgs(os.Args[1:])

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
