package main

import "errors"

func JumpToPage(jumpPage int, maxPage int) (int, error) {
	if jumpPage > maxPage || jumpPage < 1 {
		return 0, errors.New("jumped beyond the page limit")
	}
	return jumpPage, nil
}
