package main

import (
	"errors"
	"sort"
)

type Data struct {
	header []string
	body   [][]string
}

func (data *Data) Sort(header string) error {
	// array of length data.body
	index := sort.StringSlice(data.header).Search(header)
	// if it did not find anythig, it will return an index over the bound => equal to length
	if index == len(data.header) && header != data.header[0] {
		return errors.New("wrong header passed")
	}
	// get data for the header
	sort.SliceStable(data.body, func(i, j int) bool { return data.body[i][index] < data.body[j][index] })

	return nil
}

func GetData(input_data [][]string, page int, step int) Data {
	/*
		This function takes the entire dataset and returns a sliced up part for the render

	*/

	if len(input_data) == 0 {
		return Data{}
	}
	if len(input_data) == 1 {
		return Data{header: input_data[0]}
	}
	lower_limit := (page-1)*step + 1
	upper_limit := page * step
	if upper_limit > len(input_data)-1 {
		upper_limit = len(input_data) - 1
	}
	return Data{header: input_data[0], body: input_data[lower_limit : upper_limit+1]}

}
