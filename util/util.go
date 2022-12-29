package util

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Day struct {
	Number    int
	PartA     bool
	InputPath string
	TestPath  string
	Input     string
	Result    int
	Passed    bool
}

type Problem interface {
	Execute(*os.File) int
	Test(*os.File) bool
	Load() *os.File
	LoadTest() *os.File
	GetDay() int
	GetPartA() bool
}

func fileLoad(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("Couldn't open file at "+path, err)
	}
	return file
}

func (day Day) Load() *os.File {
	if day.InputPath != "" {
		return fileLoad(day.InputPath)
	} else {
		fmt.Println("Load:" + day.notImplemented())
		return nil
	}
}

func (day Day) LoadTest() *os.File {
	if day.TestPath != "" {
		return fileLoad(day.TestPath)
	} else {
		fmt.Println("LoadTest:" + day.notImplemented())
		return nil
	}
}

func (day Day) Execute(file *os.File) int {
	fmt.Println("Execute:" + day.notImplemented())
	return day.Result
}

func (day Day) Test(file *os.File) bool {
	fmt.Println("Test:" + day.notImplemented())
	return day.Passed
}

func (day Day) GetDay() int {
	return day.Number
}

func (day Day) GetPartA() bool {
	return day.PartA
}

func (day Day) notImplemented() string {
	part := "A"
	if !day.PartA {
		part = "B"
	}
	return " Day " + strconv.Itoa(day.Number) + " Part " + part + " not implemented yet."
}
