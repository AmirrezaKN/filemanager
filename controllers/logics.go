package controllers

import (
	"io/ioutil"
	"os"
)

func CreateFile(filepath string) {
	file, err := os.Create(filepath)
	CheckError(err)
	defer file.Close()
}

func ReadFile(filepath string) []byte {
	file, err := ioutil.ReadFile(filepath)
	CheckError(err)
	return file
}

func WriteFile(filepath string, data []byte) {
	err := ioutil.WriteFile(filepath, data, 0644)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
