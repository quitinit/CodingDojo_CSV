package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/term"
)

type Config struct {
	Filename string
	Pagesize int
}

func ParseCommandlineArgs(args []string) (c *Config, err error) {
	fmt.Println(args)
	var pageSize int
	var fileName string
	//pagesize, conversionErr := strconv.Atoi(args[1])
	flag.IntVar(&pageSize, "p", 0, "number of entries on one page")
	flag.StringVar(&fileName, "f", "", "filename of csv")
	// if conversionErr != nil {
	// 	return nil, errors.New("page size needs to be an integer")
	// }
	flag.Parse()
	if _, fileErr := os.Stat(fileName); fileErr != nil {
		return nil, fileErr
	}
	if pageSize == 0 {
		_, height, err := getTerminalSize()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		pageSize = height
		fmt.Println(pageSize)
	}
	return &Config{Filename: fileName, Pagesize: pageSize}, nil
}
func getTerminalSize() (int, int, error) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	return width, height, err

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
