package main

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

var (
	lightInput = []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	testInput = []string{
		"Time:        56     97     77     93",
		"Distance:   499   2210   1097   1440",
	}

	testSingleInput = []string{
		"Time:        56     97     77     93",
		"Distance:   499   2210   1097   1440",
	}
)

type record struct {
	time     int
	distance int
}

func parseLine(s string) []int {
	vals := strings.Split(strings.Split(s, ":")[1], " ")
	var values []int
	for _, v := range vals {
		if v != "" {
			val, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			values = append(values, val)
		}
	}
	return values
}

func parseInput(time, distance string) []record {
	times := parseLine(time)
	distances := parseLine(distance)
	if len(times) != len(distances) {
		panic("incorrect parsing")
	}
	var scores []record
	for i := range times {
		scores = append(scores, record{
			times[i],
			distances[i],
		})
	}
	return scores
}

func compute(input []string) {
	scores := parseInput(input[0], input[1])
	result := 1
	for _, score := range scores {
		numSolution := 0
		for a := 1; a < score.time; a++ {
			if (a * (score.time - a)) > score.distance {
				numSolution++
			}
		}
		delta := math.Sqrt(float64(score.time*score.time - 4*score.distance))
		s1 := int(-0.5 * (-float64(score.time) - delta))
		s2 := int(-0.5 * (-float64(score.time) + delta))
		if s1*(score.time-s1) <= score.distance {
			s1--
		}
		if s2*(score.time-s2) <= score.distance {
			s2++
		}
		println("Delta solution: ", s1-s2+1, "  -  ", numSolution)
		result *= numSolution
	}
	println("Num solutions: ", result)
}

// a+m = t AND m*a>d
// a(t-a) > d
// -a2+ta-d > 0
// D = t2-4d
// (-t +- V(t2-4d))/2
// Spread : V(t2-4d)
func TestDay6Phase1(t *testing.T) {
	compute(lightInput)
	compute(testInput)
}

func computeSingle(input []string) {
	time, err := strconv.Atoi(strings.Split(strings.ReplaceAll(input[0], " ", ""), ":")[1])
	if err != nil {
		panic(err)
	}
	distance, err := strconv.Atoi(strings.Split(strings.ReplaceAll(input[1], " ", ""), ":")[1])
	if err != nil {
		panic(err)
	}
	numSolution := 0
	for a := 0; a <= time-0; a++ {
		if (a * (time - a)) > distance {
			numSolution++
		}
	}
	score := record{time, distance}
	delta := math.Sqrt(float64(score.time*score.time - 4*score.distance))
	s1 := int(-0.5 * (-float64(score.time) - delta))
	s2 := int(-0.5 * (-float64(score.time) + delta))
	if s1*(score.time-s1) <= score.distance {
		s1--
	}
	if s2*(score.time-s2) <= score.distance {
		s2++
	}
	println("Delta solution: ", s1-s2+1)
	println("num of solutions: ", numSolution)
}

func TestDay6Phase2(t *testing.T) {
	computeSingle(lightInput)
	computeSingle(testInput)
}
