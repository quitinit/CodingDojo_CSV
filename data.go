package main

import (
	"errors"
	"sort"
)

type Data struct {
	header []string
	body   [][]string
}

func NewData(input_data [][]string) Data {
	copy_data := make([][]string, len(input_data))
	copy(copy_data, input_data)
	if len(copy_data) == 0 {
		return Data{}
	}
	if len(copy_data) == 1 {
		return Data{header: copy_data[0]}
	}
	return Data{header: copy_data[0], body: copy_data[1:]}
}
func (data *Data) Sort(header string) error {
	// array of length data.body
	var index int
	if header == data.header[0] {
		index = 0
	} else {
		index = sort.StringSlice(data.header).Search(header)
	}

	// if it did not find anythig, it will return an index over the bound => equal to length
	if index == len(data.header) {
		return errors.New("wrong header passed")
	}

	// get data for the header
	sort.SliceStable(data.body, func(i, j int) bool { return data.body[i][index] < data.body[j][index] })

	return nil
}

func (data Data) SliceData(page int, step int) Data {
	// it is page -1 because page starts at 1 while index at 0
	lower_limit := (page - 1) * step
	upper_limit := page * step
	if upper_limit > len(data.body) {
		upper_limit = len(data.body)
	}
	return Data{header: data.header, body: data.body[lower_limit:upper_limit]}
}
