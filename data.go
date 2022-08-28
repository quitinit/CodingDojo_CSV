package main

type Data struct {
	header []string
	body   [][]string
}

func (d *Data) Sort(header string) {
	// find the index of the header
	// var index int
	// for i, element := range d.header {
	// 	if header == element {
	// 		index = i
	// 		break
	// 	}
	// }

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
