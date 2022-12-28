package main

import (
	"fmt"
	"strconv"

	day01 "github.com/pdemirdjian/adventofcode22/day.01"
	"github.com/pdemirdjian/adventofcode22/util"
)

func main() {
	fmt.Println("Welcome to pdemirdjian's advent of code 2022!")
	days := []util.Problem{}
	for i := 1; i <= 1; i++ {
		var dayA util.Problem
		var dayB util.Problem
		switch i {
		case 1:
			dayA = day01.Day{
				Number:    i,
				PartA:     true,
				InputPath: util.RootInputPath + "01a.txt",
				TestPath:  util.RootInputPath + "01a.test.txt",
			}
			dayB = day01.Day{
				Number:    i,
				PartA:     false,
				InputPath: util.RootInputPath + "01b.txt",
				TestPath:  util.RootInputPath + "01b.test.txt",
			}
		default:
			dayA = util.Day{
				Number: i,
				PartA:  true,
			}
			dayB = util.Day{
				Number: i,
				PartA:  false,
			}
		}
		days = append(days, dayA, dayB)
	}
	for _, day := range days {
		input := day.Load()
		testInput := day.LoadTest()
		test := day.Test(testInput)
		result := day.Execute(input)
		partA := day.GetPartA()
		part := "A"
		if !partA {
			part = "B"
		}
		fmt.Println("Test Day " + strconv.Itoa(day.GetDay()) + " Part " + part + ": " + strconv.FormatBool(test))
		fmt.Println("Excute Day " + strconv.Itoa(day.GetDay()) + " Part " + part + ": " + strconv.Itoa(result))
	}
}
