package main

import "sort"

type Data struct {
	header []string
	body   [][]string
}

func (data *Data) Sort(header string) {
	// array of length data.body
	index := sort.StringSlice(data.header).Search(header)

	// get data for the header
	sort.SliceStable(data.body, func(i, j int) bool { return data.body[i][index] < data.body[j][index] })
	/* 	sort.SliceStable(data.body, func(i, j int) bool {
		return data.body[i][index] < data.body[j][index]
	}) */

}

func GetData(input_data [][]string, page int, step int) *Data {
	/*
		This function takes the entire dataset and returns a sliced up part for the render

	*/
	// it makes no sense the way I do it here, I should have a page devider function
	if len(input_data) == 0 {
		return &Data{}
	}
	if len(input_data) == 1 {
		return &Data{header: input_data[0]}
	}
	lower_limit := (page-1)*step + 1
	upper_limit := page * step
	if upper_limit > len(input_data)-1 {
		upper_limit = len(input_data) - 1
	}
	return &Data{header: input_data[0], body: input_data[lower_limit : upper_limit+1]}

}
