package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	day01 "github.com/pdemirdjian/adventofcode22/day.01"
	day02 "github.com/pdemirdjian/adventofcode22/day.02"
	"github.com/pdemirdjian/adventofcode22/util"
)

func main() {

	inputPathFlagPtr := flag.String("inputPath", "./inputs/", "path to input text files for the problems")

	specificDayPtr := flag.Int("day", 0, "specific day you want to run")

	startDayPtr := flag.Int("start", 1, "first day to process")
	endDayPtr := flag.Int("end", 25, "last day to process")

	flag.Parse()
	inputPathFlag := *inputPathFlagPtr
	specificDay := *specificDayPtr
	startDay := *startDayPtr
	endDay := *endDayPtr

	if _, err := os.Stat(inputPathFlag); os.IsNotExist(err) {
		log.Fatalln(inputPathFlag+" does not exist.", err)
	}

	if specificDay != 0 {
		startDay = specificDay
		endDay = specificDay
	}

	rootInputPath := inputPathFlag

	fmt.Println("Welcome to pdemirdjian's advent of code 2022!")
	days := []util.Problem{}
	for i := startDay; i <= endDay; i++ {
		var dayA util.Problem
		var dayB util.Problem
		switch i {
		case 1:
			dayA = day01.Day{
				Number:    i,
				PartA:     true,
				InputPath: rootInputPath + "01a.txt",
				TestPath:  rootInputPath + "01a.test.txt",
			}
			dayB = day01.Day{
				Number:    i,
				PartA:     false,
				InputPath: rootInputPath + "01b.txt",
				TestPath:  rootInputPath + "01b.test.txt",
			}
		case 2:
			dayA = day02.Day{
				Number:    i,
				PartA:     true,
				InputPath: rootInputPath + "02a.txt",
				TestPath:  rootInputPath + "02a.test.txt",
			}
			dayB = day02.Day{
				Number: i,
				PartA:  false,
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
