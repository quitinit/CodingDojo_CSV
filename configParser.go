package main

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Filename string
	Pagesize int
}

func ParseCommandlineArgs(args []string) (c *Config, err error) {
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
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return data
}
