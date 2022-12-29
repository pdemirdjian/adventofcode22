package day02

import (
	"bufio"
	"os"
	"strings"

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

type pick int

const (
	rock pick = iota + 1
	paper
	sciscors
)

type result int

const (
	loss result = 0
	draw result = 3
	win  result = 6
)

type round struct {
	input       string
	winScore    result
	totalPoints int
	MyPick      pick
	OppPick     pick
}

func (round round) parseInput() round {
	picks := strings.SplitN(round.input, " ", 2)
	oppPick := picks[0]
	myPick := picks[1]
	round.OppPick = pickParse(oppPick)
	round.MyPick = pickParse(myPick)
	return round
}

func (round round) pointParse() round {
	myPick := round.MyPick
	oppPick := round.OppPick
	if myPick == oppPick {
		round.winScore = draw
	} else if myPick == rock && oppPick != paper {
		round.winScore = win
	} else if myPick == paper && oppPick != sciscors {
		round.winScore = win
	} else if myPick == sciscors && oppPick != rock {
		round.winScore = win
	} else {
		round.winScore = loss
	}
	return round
}

func (round round) totalParse() round {
	round.totalPoints = int(round.winScore) + int(round.MyPick)
	return round
}

func pickParse(input string) pick {
	switch input {
	case "A", "X":
		return rock
	case "B", "Y":
		return paper
	case "C", "Z":
		return sciscors
	}
	return 0
}

func day02Data(file *os.File) []round {
	rounds := []round{}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		round := round{}
		round.input = fileScanner.Text()
		rounds = append(rounds, round)
	}
	return rounds
}

func (day Day) Execute(file *os.File) int {
	if day.PartA {
		rounds := day02Data(file)
		for i, round := range rounds {
			parsedRound := round.parseInput().pointParse().totalParse()
			rounds[i] = parsedRound
		}
		total := 0
		for _, round := range rounds {
			total += round.totalPoints
		}
		return total
	} else {
		return util.Day{
			Number: 2,
			PartA:  day.PartA,
		}.Execute(file)
	}
}

func (day Day) Test(file *os.File) bool {
	if day.PartA {
		return day.Execute(file) == 15
	}
	return util.Day{
		Number: 2,
		PartA:  day.PartA,
	}.Test(file)
}

func (day Day) Load() *os.File {
	return util.Day{
		Number:    2,
		InputPath: day.InputPath,
	}.Load()
}

func (day Day) LoadTest() *os.File {
	return util.Day{
		Number:   2,
		TestPath: day.TestPath,
	}.LoadTest()
}

func (day Day) GetDay() int {
	return 2
}

func (day Day) GetPartA() bool {
	return day.PartA
}
