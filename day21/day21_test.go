package day21

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

var (
	lightInput = []string{
		"...........",
		".....###.#.",
		".###.##..#.",
		"..#.#...#..",
		"....#.#....",
		".##..S####.",
		".##..#...#.",
		".......##..",
		".##.#.####.",
		".##..##.##.",
		"...........",
	}

	largeInput = []string{
		"...................................................................................................................................",
		"...............#.....#...#......#...............#.....#...#...........#..#....##.#..#.........#..............#...........#..#......",
		"........#.#..........#.##..#.......#.#.#.......#..#........................#............#...#....#.............#....#.........#....",
		".....#......#...................#...#......#.............#..............#.#....#.......#...#...#......##.#.#...........#..###......",
		".........#............#...##.........#.#....#........#..........................#....#.#......#......#....#....#............#.#....",
		"...#.....#................#.#..........#..........#.#.#...........#.......#..............#......#.......##.....#.........##.#......",
		"..........#....................#............#........#............#............#....#..........#.........................#.........",
		"...#.#...................#.#......##..#........#.##.....................................#........##............#.##....#...........",
		".........................................#...#.......................#...........................#..#........#....##.#...##........",
		"...#.......#...#.......#..#.#.......#.......#........................#........#....#....#...#.#.....#...#...........#........##.#..",
		"..#....#.............#........................#..##.............................#..#.................#....##.......#.......#....#..",
		".........................#.......#...#...........#..................#..............#.#....#.....#...#........#...........#......#..",
		"............#..................#...#.......#...#.#.........#.......................#............#.......#....#.....#..........#....",
		"..........#........#......##.#.....#...#................#....##...........#................#.......#........#......#...#....#..#...",
		".................#.......................##...............#.....#.....#............#.................#..........#......#.....#.....",
		".....#.###................#..........#...................#.##........#...............#....#...............#..#................#.#..",
		".#..#..........#..................#...................................#......#.................#.#..................#..............",
		"..........#..#...#.#...........#.#..#....................#....................#..........#.#.#.....#....#...........#........#.....",
		".#.............#.......##.##.........#.#..............#.......#.............#.........................................#....#.#.....",
		".............#.....#.#........#..........#..............#.....................#..................#..#........#.#.....#...#.........",
		"....#.......#..............#.....#.......#...........#.............#.##..................#.................#..#......#.#...#....#..",
		"..#....#.....#......................#.....................#..........#..##........#.........#.............#......#.................",
		"...##.......#.......#...#..##..................#.........##....................#.#...................................#..#..........",
		"........#.........#..##........................#......#........#...#...#........#......................................#.......#...",
		"....#..#....#.#.#.................#.............#.........#.............#....#................#..#.....#...##.....#.........##.....",
		"...................#................#...............................#.................#...........#.....##.#.......................",
		".....#........#........#........................#..............#.............................................#........#...#........",
		".......#............#.................................#.......#........#...#..............................#.....#.#.#......#.#.#...",
		".....................###....#.#................#.#....#........##.#..#.......#...#............................#..........###....#..",
		".#.....##.#......#.#........#.#..............#...............#.........#.........#.....#...........#.............................#.",
		"...#.....#...........#.#......#................#..#....#...#..#......#................#....#.............#.#....#..................",
		"..#....##......#.............#...............#.................................#.......#................#...........#.#............",
		"..........#....#....##.......................#.#.#..................#.....#..............#................#..................#.....",
		".....#......................................#....#......................#..........#........#........................#.............",
		"..#.#..#........#...#..#...................#............#..#...#......#.#...#............#.....#...............#..............#....",
		"..#....#..##..........#...........#......................##................#...#...............................#..#....#........#..",
		"......#........#...................#.#.......#..#.#....#..##...#.........##......#....#..........#....................#....#.......",
		".#.##.....#.....#..................#.....#..#..#....#....#......#...........#......#....#.........#........##..#.#..##.............",
		"........#..#.....#.....#.......#.#......................#..................#.#.......##...................................#..#.....",
		"..#........#.........................#.......#..............................#.....##.........#..#.##...............................",
		"..............................#....................#......................#.....#..#.............................#..........#......",
		"........#.#.....................#......#.....#.............##...............#....#.#...#..........................#...#....#.......",
		"........#...................#.#.##.##........................#.#........................#...............................#..#...#.#.",
		".............##.................#......#.#.##.......#.........................#.......................#.................#..........",
		"..##.....#..........................##..........#....#.............................##.....###......#...........................#...",
		"...#.......#...#..............#................#.#.....#..#..................#..#...##....#....#.......#...........#.........#.....",
		".....#.......#.........#..#.......##.............#.##...#.......#................#..........#............#.................#.......",
		"..............#..........##....................#......#...##..................................#.#.......#.#..............#.........",
		"..#.....#..#.#.............#................#......................#.#.................##.......###.#....#..............##.........",
		"....................#.........#...#............#.#......#..........#........#..........................#..#.#............##........",
		".......#...#..........#.........#........#....#......###.......#..#.........##....#..............#.......#..#..........#...........",
		".....##.#.................................#.....#.............#.#.......#..##.............#........................................",
		".......#..................##.##..#............#......#...#....##........##.#................#....#..#..##..........................",
		"..##.#..............###..#........#...#...#.........#.....................#....#.......##............#...#..................#......",
		"....#..............#...........#.....##................##...##.......#...#....#...#............#..#..............................#.",
		".........................#.........#..................##................#....#...#...........#......#..#....####.............#.....",
		".....................#...#....#........#...........#..........#....................#..#..#..........#..............................",
		".#..................#..................#........#............#.#..............#....#...............#........#....##............##..",
		"...............#....#.#.........#...#.......#...#.......#.......#................#.##................#.................#...........",
		"..............#.#.......#...#....#......#.............#...................##...#.......##......##..................................",
		"..........##.#..#.......##........#.##......#............#.#..##...............#...#.....#...#............#......#.................",
		"........#.#.#........................................##....#..#...............#.......##......#......#...........#..##.............",
		"......................................#........................#.........#.....#...................#...............................",
		".........................#....#.......#...................#..#..#.#.....#..................#......##.......................#.......",
		"......#....#.......#.............#.........................#........#.#...#.....#..#....#......#....#.#.......#...........#........",
		".................................................................S.................................................................",
		"...........#...#.....#......................#.......#.............#.#............#.##.#...#........#.......#........#.......##.....",
		"................#..#........#...#.....................#...............#...........................##.....#.........#...............",
		".............##.##.#...#........#..#...#....................#...#.......................#............#.............................",
		"................#....#.#.#...#...............#...........#..............................##........................#................",
		".......................................#.....#...........#.........#...................#..#..#.................#.........#.......#.",
		"..#...............#....#.#.#............................#.#..#............#.........#......#........#....#......................#..",
		"...#......................#....#.#....................##..........#......#...#........#....#.....................#.#...........#.#.",
		"...............#...#.................................##...#.#.......#...#...........#...................#...........#..............",
		"................#......#.............#.#...##...##......#......#..#..........#..#......#...........................................",
		"..................#.#.....#.##.....#....#..........................##.#.#....#.........#...#.....#........#...##.#...............#.",
		"........................#...#..............#....#.......#......#......#......#..#...............##.............##..................",
		"..#...##.............#..................#....#.......................#..........#...##......................#..#..#..........#...#.",
		"....##.#..........#...#..........#.......#...#............#.#.........#..#...#.#.......##......#..###..#.........#.................",
		".........#.............#...#..........#........#..##..................#..#.........................#.....#...#..............#......",
		".......#...#.............#.#....###................#...#.................#.#.............#..................##...............#.....",
		"....#....##..............#......#.....#.....#...#..#............#...##.......#.#....#...#........##.#..#...............#...#.......",
		".#..#....#............#.........#...........#......#...#............#...........#......##....#........#..........................#.",
		"....###...#.................#.......#............#.#.........#...................#.##............#....................#............",
		"........#...............................#..........#................#......#....#.............#..........##..........#.....#..##...",
		"....#.....###...................##..##.##...#..##...##...#....#.......#..................#........#..#...................#.........",
		"..#.#............#................#...#.#.......#.........#.....#.#.....#.............#.....#..#.....................#..##.........",
		"............#.#..#....................#.....#..........#.#...................#....................#..#.......................#..#..",
		"....#....#.....#..................##................#..#...#.......##...............#.....##..#................#..##..##...#.......",
		"..................#..............................##................#....#....#........................................#...#...##...",
		".............##........................#....#........##......#.............#.#...#................#..........#..#.#.....#.......#..",
		"....#.#..#...#..#..................#..#..#..#.#......#..#..#.#.....####......#...........##.#....#..#..............#............#..",
		".#...#.................................#............##...#....................#.#..#....#.......#................#..#..........#...",
		"..#.................#.................#...................#.#.......#....#...............#......#.#...............#................",
		"......###.....#.........#.................#.#.............#..#..#.......#.....#.............................#....##..##..#.........",
		"..........###....#.#.#....#........#............###..#.........#....#...........#.......#......#........#......#....#............#.",
		".....##...#.#.##....................##...#.#..#......#.......#......#...#........#.....................#.......#......#.....##.....",
		".#.....#..#...............................#....#.................................##...................#.......#...............#....",
		".....................##.....#...................#......................#.#..#................................##....................",
		"....#....#..........#........#.........#...##..........#.....#..#..#.............#..##....................................#........",
		"..........#.#.............#.................#............#.....#......#...............................#.................#....#.#...",
		"......#...#.#.....#.#.##..#.#............###....#.....#................##......#.......#..#....................#.....#.............",
		".....#....#....................#.....................#..........#.....#.............#..#.......................#.#.................",
		"......#....#...............................#...#........#.......................#....................##...#..#......#...##....#....",
		"....#..............#...........#................#.#......#........##...#..#....###..................#.........#.............#......",
		"......#....#....#.....#..##.#.#.............##.......#.......#....#...##..........................#....###...#...#.............#...",
		".#....##......#........#.#..........#..........................................#................#..................#...............",
		"..#......#...#....#.......##.....#.....................#.....#.#..#...#...#......#.......................................#..#......",
		"......#..............#.......#.#..........................#.............#.......#.#.............#......#....................#..#...",
		"..#...#.........................#.#....#............................#.....................#.....#.................#................",
		"...................##......#.....#.............................#......##....#............#...#......................#..............",
		"......#...#..#.##.............#........#..#.........##...#....#......#..#......#..........#...........#..........#.#..#.#..........",
		".#..#..#.........#.......#.......#.........#..........##...##...#.#...#........................................#....#..............",
		".#..#.#.#....#......#................#........................#.#.....#........................#...#......................#......#.",
		"...............##...............##....................#..#.#.........................#.......#.......#...........#............#....",
		"..........#.........#.#.............#....#..#.#.............##....#......#.#.........................#..#.....#.......#........#...",
		"....##......#...#.............................#...............#......#................#.....................................#......",
		"....###.#..#.......#...................#....#..........................#..#..................#.....#....#.....#................#.#.",
		"............#.#..#.#.###........##..#.##...........................#.................#.#.............#.............#.....##........",
		"..#......#................................#.#.....................#......................#........#.....#.....#.#...#.....#........",
		".............#..........#.#.##......#..#....................##................................#..............#......#...........##.",
		".....#...#..............#..#....#..#..............................#..........................#.###..............#...##.....##......",
		".#...#..#.......#.#...............#....#.....#..#....#..............#...................................................#.#.#......",
		".......##.........#.#....#...........#.....#....#...............#....................#..................##......#........#.........",
		"...#.#.....#....#.................#..#................#........#..................#......#............#....#...........#..#........",
		".....#...#..................#.......#.........#...#....#...................#...##..#.#....#....#.##..#..##.##..................#...",
		".....................#..#...#.....#............#..........................#.##.....................#..#...#......##.....#.....#....",
		"..#........#.......##.#..............#..##................................#..#.#....#.....#...............#.................#.#....",
		"...#.................................#.........#...........................#....##.#...................#......##.................#.",
		".........#..........#.#......#....##....#.#...#.........#................................#....#..#..............#....##...#........",
		"...................................................................................................................................",
	}
)

type position struct {
	x, y int
}

type move struct {
	num       int
	direction byte
}

const (
	MOVE_UP    = 'U'
	MOVE_DOWN  = 'D'
	MOVE_RIGHT = 'R'
	MOVE_LEFT  = 'L'
)

func bytes(input []string) [][]byte {
	res := make([][]byte, len(input))
	for i := range input {
		res[i] = []byte(input[i])
	}
	return res
}

func display(gardenMap [][]byte, pos map[position]bool) {
	xMin := math.MaxInt
	xMax := math.MinInt
	yMin := math.MaxInt
	yMax := math.MinInt
	for p := range pos {
		if p.x < xMin {
			xMin = p.x
		}
		if p.y < yMin {
			yMin = p.y
		}
		if p.x > xMax {
			xMax = p.x
		}
		if p.y > yMax {
			yMax = p.y
		}
	}
	yStart := 0
	ySize := len(gardenMap)
	for yStart > yMin {
		yStart -= ySize
	}
	xStart := 0
	xSize := len(gardenMap[0])
	for xStart > xMin {
		xStart -= ySize
	}
	yStop := ySize
	for yStop <= yMax {
		yStop += ySize
	}
	xStop := xSize
	for xStop <= xMax {
		xStop += xSize
	}
	var str strings.Builder
	for y := yStart; y < yStop; y++ {
		for x := xStart; x < xStop; x++ {
			if pos[position{x, y}] {
				fmt.Fprint(&str, "O")
			} else {
				yNdx := y
				for yNdx < 0 {
					yNdx += ySize
				}
				for yNdx >= ySize {
					yNdx -= ySize
				}
				xNdx := x
				for xNdx < 0 {
					xNdx += xSize
				}
				for xNdx >= xSize {
					xNdx -= xSize
				}
				fmt.Fprintf(&str, "%c", gardenMap[yNdx][xNdx])
			}
		}
		fmt.Fprint(&str, "\n")
	}
	fmt.Println(str.String() + "\n----------------\n")
}

func isValid(x, y int, modulo bool, gardenMap [][]byte) bool {
	xMax := len(gardenMap[0])
	yMax := len(gardenMap)
	if !modulo {
		return x >= 0 && y >= 0 && x < xMax && y < yMax && gardenMap[y][x] != '#'
	}
	for x < 0 {
		x += xMax
	}
	for y < 0 {
		y += yMax
	}
	for x >= xMax {
		x -= xMax
	}
	for y >= yMax {
		y -= yMax
	}
	return gardenMap[y][x] != '#'
}

func positions(start position, gardenMap [][]byte, modulo bool, count int) map[int]int {
	pos := map[position]bool{start: true}
	res := make(map[int]int)
	for i := 0; i < count; i++ {
		moved := make(map[position]bool, count*count)

		for p := range pos {
			if isValid(p.x, p.y+1, modulo, gardenMap) {
				moved[position{p.x, p.y + 1}] = true
			}
			if isValid(p.x, p.y-1, modulo, gardenMap) {
				moved[position{p.x, p.y - 1}] = true
			}
			if isValid(p.x+1, p.y, modulo, gardenMap) {
				moved[position{p.x + 1, p.y}] = true
			}
			if isValid(p.x-1, p.y, modulo, gardenMap) {
				moved[position{p.x - 1, p.y}] = true
			}
		}
		if false {
			display(gardenMap, moved)
		}
		pos = moved
		res[i+1] = len(pos)
	}
	return res
}

func computePositions(input []string, modulo bool, count int) int {
	gardenMap := bytes(input)
	start := position{}
	for y := 0; y < len(gardenMap); y++ {
		for x := 0; x < len(gardenMap[0]); x++ {
			if gardenMap[y][x] == 'S' {
				start = position{x, y}
				break
			}
		}
	}
	res := positions(start, gardenMap, modulo, count)
	return res[count]
}

func TestDay21Phase1(t *testing.T) {
	// 16
	c := computePositions(lightInput, false, 6)
	if c != 16 {
		t.Fail()
	}
	//
	c = computePositions(largeInput, false, 64)
	if c != 3776 {
		t.Fail()
	}
}

// guess positions
func guessPositions(base int, center, other,
	lArrow, rArrow, uArrow, dArrow,
	brCorner, tlCorner, trCorner, blCorner,
	brTrunc, tlTrunc, trTrunc, blTrunc int) int {

	// Sample strucut for 2
	// E	TLC		UA	TRC 	E
	// TLC	TLT		O	TRT		TRC
	// LA	O		C	O		RA
	// BLC  BLT		O	BRT		BRC
	// E	BRC		DA	BRC		E

	return lArrow + rArrow + uArrow + dArrow + // Each arrouw once
		(brCorner+tlCorner+trCorner+blCorner)*base + // Triangle appears <base> times
		(brTrunc+tlTrunc+blTrunc+trTrunc)*(base-1) + // truncated rectanble appear base -1 (between triangles)
		other*base*base + // other appears base times in the middle and ... up to 1 on both side
		center*(base-1)*(base-1) // center apres base-1 ... up to 1  on both side
}

func TestDay21PhaseExplo(t *testing.T) {
	computePositions(lightInput, true, 50)
}

func TestDay21Phase2(t *testing.T) {
	testCases := []int{6, 10, 50, 100}       // , 500, 1000, 5000}
	testResults := []int{16, 50, 1594, 6536} // , 167004 , 668697, 16733044}

	for i, s := range testCases {
		c := computePositions(lightInput, true, s)
		if c != testResults[i] {
			t.Fail()
		}
	}

	gardenMap := bytes(largeInput)

	// arrows
	lArrow := positions(position{-1, 65}, gardenMap, false, 131)
	rArrow := positions(position{131, 65}, gardenMap, false, 131)
	uArrow := positions(position{65, 131}, gardenMap, false, 131)
	dArrow := positions(position{65, -1}, gardenMap, false, 131)
	fmt.Printf("\t^\n\t%d\n", uArrow[131])
	fmt.Printf("< : %d - %d >\n", rArrow[131], lArrow[131])
	fmt.Printf("\t%d\n\tv\n", dArrow[131])

	brCorner := positions(position{0, -1}, gardenMap, false, 65)
	tlCorner := positions(position{131, 130}, gardenMap, false, 65)
	trCorner := positions(position{0, 131}, gardenMap, false, 65)
	blCorner := positions(position{131, 0}, gardenMap, false, 65)
	fmt.Printf("Triangles:\ntl: %d\nbl: %d\ntr: %d\nbr: %d\n", tlCorner[65], blCorner[65], trCorner[65], brCorner[65])

	brTrunc := positions(position{0, -1}, gardenMap, false, 197)
	tlTrunc := positions(position{131, 130}, gardenMap, false, 197)
	trTrunc := positions(position{0, 131}, gardenMap, false, 197)
	blTrunc := positions(position{131, 0}, gardenMap, false, 197)
	fmt.Printf("Truncated:\ntl: %d\nbl: %d\ntr: %d\nbr: %d'n", tlTrunc[196], blTrunc[196], trTrunc[196], brTrunc[196])

	fullMap := positions(position{len(gardenMap[0]) / 2, len(gardenMap) / 2}, gardenMap, false, len(gardenMap)+1)
	other := fullMap[130]
	center := fullMap[129]
	fmt.Printf("%d = %d\n", center, other)

	iterations := 26501365
	fmt.Printf("Full count : %d\n", iterations)
	fmt.Printf("Number of full maps : %d\n", iterations/131)
	fmt.Printf("Remaining : %d\n", iterations%131)

	c := guessPositions(2, center, other,
		lArrow[131],
		rArrow[131],
		uArrow[131],
		dArrow[131],
		brCorner[65],
		tlCorner[65],
		trCorner[65],
		blCorner[65],
		brTrunc[196],
		tlTrunc[196],
		trTrunc[196],
		blTrunc[196])
	if c != 95816 {
		t.Fail()
	}
	fmt.Printf("Results for 2*131+65 : %d\n", c)

	c = guessPositions(4, center, other,
		lArrow[131],
		rArrow[131],
		uArrow[131],
		dArrow[131],
		brCorner[65],
		tlCorner[65],
		trCorner[65],
		blCorner[65],
		brTrunc[196],
		tlTrunc[196],
		trTrunc[196],
		blTrunc[196])
	if c != 310036 {
		t.Fail()
	}
	fmt.Printf("Results for 4*131+65 : %d\n", c)

	c = guessPositions(6, center, other,
		lArrow[131],
		rArrow[131],
		uArrow[131],
		dArrow[131],
		brCorner[65],
		tlCorner[65],
		trCorner[65],
		blCorner[65],
		brTrunc[196],
		tlTrunc[196],
		trTrunc[196],
		blTrunc[196])
	if c != 646544 {
		t.Fail()
	}
	fmt.Printf("Results for 6*131+65 : %d\n", c)

	c = guessPositions(iterations/131, center, other,
		lArrow[131],
		rArrow[131],
		uArrow[131],
		dArrow[131],
		brCorner[65],
		tlCorner[65],
		trCorner[65],
		blCorner[65],
		brTrunc[196],
		tlTrunc[196],
		trTrunc[196],
		blTrunc[196])
	if c != 625587097150084 {
		t.Fail()
	}
	fmt.Printf("Results for %d*131+65 : %d\n", iterations/131, c)
}
