package day01

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/pdemirdjian/adventofcode22/util"
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

type elf struct {
	calories []int
}

func (day Day) Load() *os.File {
	utilDay := util.Day{
		Number:    1,
		InputPath: day.InputPath,
	}
	return utilDay.Load()
}

func (day Day) GetDay() int {
	return 1
}

func (day Day) GetPartA() bool {
	return day.PartA
}

func (day Day) LoadTest() *os.File {
	utilDay := util.Day{
		Number:   1,
		TestPath: day.TestPath,
	}
	return utilDay.LoadTest()
}

func (day Day) Execute(file *os.File) int {
	elves := day01Data(file)
	if day.PartA {
		return calculateResults(elves, 1)
	} else {
		return calculateResults(elves, 3)
	}
}

func (day Day) Test(file *os.File) bool {
	if day.PartA {
		return day.Execute(file) == 24000
	} else {
		return day.Execute(file) == 45000
	}
}

func calculateResults(elves []elf, maxToCount int) int {
	max := make([]int, maxToCount)
	for _, elf := range elves {
		sort.Ints(max)
		total := 0
		for _, snack := range elf.calories {
			total += snack
		}
		for i, num := range max {
			if num < total {
				max[i] = total
				break
			}
		}
	}
	sum := 0
	for _, num := range max {
		sum += num
	}
	return sum
}

func day01Data(file *os.File) []elf {
	elves := []elf{}
	fileScanner := bufio.NewScanner(file)
	elf := elf{}
	for fileScanner.Scan() {
		text := fileScanner.Text()
		if text == "" {
			elves = append(elves, elf)
			elf.calories = nil
		} else {
			num, err := strconv.Atoi(text)
			if err != nil {
				log.Fatalln("Couldn't convert " + text + " to int.")
			}
			elf.calories = append(elf.calories, num)
		}
	}
	elves = append(elves, elf)
	defer file.Close()
	return elves
}
