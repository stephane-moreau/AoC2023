package main

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

var (
	lightInput = []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}

	largeInput = []string{
		"..#............................#....................................................................................#..........#........#...",
		"..........#.........................................#.......................................................................................",
		"...................#..........................................................................#.............................................",
		".....................................................................................#.............................................#........",
		".......#.......#.............#..........#......#...................#.....#..................................................................",
		"................................................................................#..................#..............#.........................",
		"....................................#.......................................................................................................",
		"...........#...................................................................................#......................................#.....",
		".......................................................................................................................#.........#..........",
		".....................#.................................................#...................#.................#..............#...............",
		"....#...........................................#...............#...........................................................................",
		".........................#.......................................................................#..........................................",
		"...................................#......................#.......................................................#.......................#.",
		"..................#...........................................................#.......................#.................#...........#.......",
		"..........#........................................................#..........................................#.............................",
		".............................................#.........#.....#..............................................................................",
		".......................#.....................................................................#..............................#..........#....",
		".............................#.......................................................#...........................#..........................",
		"..#...........#.......................#.....................................#...............................................................",
		"...........................................................#.............................................#..................................",
		"......#..........................................................#......#........................................................#..........",
		"................................#.........#...........#....................................#........#...................#..................#",
		"...........................#..................................................#..............................#..............................",
		"..............................................#.................................................#...........................................",
		".#.................................................................................................................#........................",
		".......................................................................#....................................................................",
		"...............#.............................................................................................................#..........#...",
		".................................#.............................#............................................#...............................",
		".....#.................#...................#................................................................................................",
		"....................................................#...................................................#............#......................",
		"...........#.....#...................................................................#............................................#.........",
		"..............................................#...........................................................................#.................",
		"...#.........................#.....................................................................#........................................",
		".........................................................#...................................#........................................#.....",
		".........#....................................................#........#................#.....................#.............................",
		"........................#..........................................................................................#........................",
		"..............#...................................#............................#.............................................#.....#........",
		".................................#................................................................#.........................................",
		"..................#........................#...............#................................#.............#.............................#...",
		"..........#........................................................#............................................................#...........",
		"..#.....................................................................................#.............................#.....................",
		"..........................#..........#..................................#......................................#...........#................",
		"................................................................#...........................................................................",
		"................#................................#..........................#...............................................................",
		"...........................................#..................................................#.....#.......#...............................",
		"..........#............................................#.....................................................................#..............",
		".......................#....................................#...........................................................#..............#....",
		".....#.........................#........#..............................................#.................#........................#.........",
		".................#....................................................#............................................#........................",
		"....................................#...............#................................................#......................................",
		".........................#........................................#................#.......#...................#............................",
		"....................................................................................................................................#.......",
		"...........#...................................#...........................#......................#.........................#..............#",
		"..................................#......#..................................................................................................",
		"......#....................#.........................#..................................................#...................................",
		"......................#...................................#......#.....#.................................................#..................",
		"..............#..............................................................................#.................#.....................#......",
		"............................................................................................................................................",
		"#.............................#................#....................................................................#....................#..",
		"............................................................................................................................................",
		"........#...........................................#.......................................................................................",
		"...................................................................#..........#..................#..........................................",
		".....................#............#........#.............................................#............#................#....................",
		"................#............................................................................................................#..............",
		"...............................................................................................................#...................#........",
		"........................#......#..........................#.....#......#................................................................#...",
		"................................................#..........................................#.............#..................................",
		"......................................#..............................................................................#......................",
		"..#.................................................................#.........#.............................................................",
		"..........#.....#...............................................................................................#...........#...............",
		".................................#....................................................................#..............................#......",
		".....................................................#........................................#........................#....................",
		"..........................................#..................#............#.................................................................",
		"......................#...............................................................#..................#.........#............#...........",
		"......#........................................................................#............................................................",
		"...................................#...................................#..................................................#.................",
		"...................................................#.........................................#.................#............................",
		"...........................#...................................#.....................................................................#......",
		"........................................#...................................................................................................",
		"................#......................................#........................#.........#.....#...........................................",
		"...........#......................#.....................................................................#...................................",
		"....#....................#...................#...................#................................................#...........#...........#.",
		"............................................................................................................................................",
		"............................................................................................................................................",
		".....................................#....................................#.................................................................",
		".......................#..................#........#............................#................#.............#............................",
		"...............#...............................................#......................#..............................................#......",
		"............................................................................................................................................",
		".................................#...........#.....................#....................................................#...................",
		"............................................................................................................................................",
		"..#........#..........#...............#..................................................................#..........#......................#",
		"............................................................#............................#..................................................",
		".................................................................#..........#.................................#.................#...........",
		".............................#......................#.......................................................................................",
		"....#..................................................................................................................................#....",
		"...................#...............#..................................................#....................#................................",
		".........................................................#......................................#...................#..............#........",
		".....................................................................#...................................................#..................",
		"........................#...................................................................#..................#............................",
		"...#.....#...............................#....................#.............................................................................",
		"................#...............#...........................................................................................................",
		".....................#........................................................#.....................................................#.......",
		"..............................................#.......#................#...............................................#...................#",
		".........................................................................................#...............#...................#..............",
		".#...........................#.......#......................................................................................................",
		".............................................................................................................#........................#.....",
		".................................#..........................................................................................................",
		"...........#........................................................................................#.............................#.........",
		"................#...............................................#............................#...................#..........................",
		"...#............................................................................#...........................................................",
		"....................................................#.......#.......#......#...........................................................#....",
		".........................#..................................................................................................................",
		"....................#...............#.........................................................................................#.............",
		".........................................................#....................#......................#......................................",
		"....#........................#...............#........................#...............#.....................................................",
		"..........#......#....................................................................................................................#.....",
		"..............................................................................................#...................#.........................",
		".......................................................#...................................................#.................#..............",
		"..........................#....................#............................................................................................",
		"........................................#......................#.........#...............#.......#......................#.........#.....#...",
		"...............................#....................................#.......................................................................",
		"........#...........#.............................#.........................................................................................",
		"...#...........#...................................................................#..................#............#........................",
		"...........................#..............................#...................................#.............................................",
		"...................................#.............................#..........................................#...............................",
		"..............................................#........................#........#..........................................#................",
		".......................................................................................#.............................#..............#.......",
		".........................................#..................................................................................................",
		".....#............#......................................#.....#............................#...............................................",
		"............................#...............................................................................................................",
		".......................#.........#........................................#..............................#...............................#..",
		"..#...........#.......................................#.........................#...........................................................",
		".......................................#.........................#..........................................................................",
		"............................................................#...............................................................................",
		".......#......................#........................................#................................................#..............#....",
		"..............................................................................................................#.............................",
		"...............................................#.....................................................#......................................",
		"......................#..................................................................#......#.................................#.......#.",
		"..#.......#................#......#..................................#.......#.....................................#........................",
		"................#...............................................#..................#....................#..................#................",
	}
)

func expandUniverse(lines []string) ([][]byte, []int, []int) {
	emptyLines := make([]int, 0, len(lines))
	for i, line := range lines {
		if strings.IndexByte(line, '#') == -1 {
			emptyLines = append(emptyLines, i)
		}
	}
	emptyCols := make([]int, 0, len(lines[0]))
	for j := 0; j < len(lines[0]); j++ {
		found := false
		for i := range lines {
			if lines[i][j] == '#' {
				found = true
				break
			}
		}
		if !found {
			emptyCols = append(emptyCols, j)
		}
	}
	converted := make([][]byte, 0, len(lines))
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		newLine := make([]byte, len(line))
		for j := 0; j < len(lines[0]); j++ {
			newLine[j] = line[j]
		}
		converted = append(converted, newLine)
	}
	return converted, emptyLines, emptyCols
}

func display(lines [][]byte) {
	for _, line := range lines {
		fmt.Println(string(line))
	}
}

type position struct {
	x, y int
}

func computeDistances(input []string, age int) {
	expanded, lines, cols := expandUniverse(input)
	fmt.Printf("%v\n%v\n", lines, cols)
	display(expanded)
	positions := make([]position, 0, 10)
	for i := 0; i < len(expanded); i++ {
		for j := 0; j < len(expanded[0]); j++ {
			if expanded[i][j] == '#' {
				positions = append(positions, position{i, j})
			}
		}
	}
	fmt.Printf("%v\n", positions)

	paths := 0
	for i := 0; i < len(positions); i++ {
		for j := 0; j < i; j++ {
			dist := int(math.Abs(float64(positions[i].x-positions[j].x)) + math.Abs(float64(positions[i].y-positions[j].y)))
			fmt.Printf("%d%v - %d%v:", i, positions[i], j, positions[j])
			xI, xJ := positions[i].x, positions[j].x
			if xI > xJ {
				xI, xJ = xJ, xI
			}
			for _, l := range lines {
				if xI < l && xJ > l {
					dist += age - 1
				}
			}
			yI, yJ := positions[i].y, positions[j].y
			if yI > yJ {
				yI, yJ = yJ, yI
			}
			for _, c := range cols {
				if yI < c && yJ > c {
					dist += age - 1
				}
			}
			fmt.Printf(" %d\n", dist)
			paths += int(dist)
		}
	}
	fmt.Printf("Paths : %d\n", paths)
}

func TestDay11Phase1(t *testing.T) {
	// 374
	computeDistances(lightInput, 2)
	// 10313550
	computeDistances(largeInput, 2)
}

func TestDay11Phase2(t *testing.T) {
	// 374
	computeDistances(lightInput, 2)
	computeDistances(lightInput, 10)
	computeDistances(lightInput, 100)
	// 10313550
	computeDistances(largeInput, 1000000)
}
