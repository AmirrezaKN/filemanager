package controllers

import (
	"io/ioutil"
	"os"
)

// CreateFile will create a file
func CreateFile(filepath string) {
	file, err := os.Create(filepath)
	CheckError(err)
	defer file.Close()
}

// ReadFile will read a file
func ReadFile(filepath string) []byte {
	file, err := ioutil.ReadFile(filepath)
	CheckError(err)
	return file
}

// WriteFile will write a file
func WriteFile(filepath string, data []byte) {
	err := ioutil.WriteFile(filepath, data, 0644)
	CheckError(err)
}

// CheckError will check an error
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
