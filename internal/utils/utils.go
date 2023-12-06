package utils

import "os"

func ReadInput(path string) *os.File {
	pwd, _ := os.Getwd()
	input, err := os.Open(pwd + path)
	Check(err)
	return input
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
