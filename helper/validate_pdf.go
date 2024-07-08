package helper

import "path/filepath"

func ValidatePdf(filename string) bool {

	if filepath.Ext(filename) == ".pdf" {
		return true
	}
	return false

}
