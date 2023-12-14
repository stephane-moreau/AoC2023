package day14

import (
	"fmt"
	"strings"
	"testing"
)

var (
	lightInput = []string{
		"O....#....",
		"O.OO#....#",
		".....##...",
		"OO.#O....O",
		".O.....O#.",
		"O.#..O.#.#",
		"..O..#O..O",
		".......O..",
		"#....###..",
		"#OO..#....",
	}

	largeInput = []string{
		"O...O..###.#..#O...###.......#....O#.O...O....#.O...O..O..##.O.O#...O........O...#.#.O...O#O.......#",
		"...O#..#O...#....O..###.O.....#......#.OO.O#..#.O..##...#....##...#.O....##O##OOO..#.O..OOOO..#..#.#",
		"...##.OOOOO#.OO....OO#O.###.O..O...OO.......O....O...OO##....#..O.....#O#.O.O.#O..O.##....##.#.OOOO#",
		"O..#O..O.....#.#...#....#O.O..#..O.O..#O#..O....O..O.#.O.#..O.O#...#O#...##O.#O.##.OO...#.#.#OOO....",
		".#O#O.O.......#O.#.O#.O..##.#.O.....#..#.....O..#..###..OO.#.O.OO#..O.O.........O..O...#.O.#....OO#.",
		"..........#...O.OO....#.O.O##OO#......#OO.#.O...O....O..#OO#..O...#..#.#.....O#O#..OO#...#OOO.O.#O..",
		"#O#OO.O.#.O.O...O.###O.........#O.O.#...O.#..O.O..#..O##.#O.O.OO.....OO.#.#.#..#.O..O#......O..O....",
		".#....O#O.OO..#..O#..O..O..O.#..O..#.O#....#.O....O.O.....#O#....#.#.#..O...O#OO#.#O.#O.#...OO..OO#O",
		".O.........#.#O#....O.O#OOO......O..#O#...#.....#.......#...#....#.##.O...O..#O..#O...O.#..O.OO..#..",
		"O...O...O.....O..O...........OO.##..OO.OO#.O.##..O#OO#......O..O#.O#.OO..#####......O....#....O...#.",
		".....#O.....#....#O......#..OO.O.O.#..O..O....O..#.#..#.......O..O.O....O.OO......OOO#.OOO......O...",
		"..O#.O.O.#O#O...#..........#..O##...O..O#OOO.OO..##O.#OOOO...#..O....#.....O...#....#.#.#...OO.O.#..",
		"#.......#.O....O#.#O..#OOO..O..###O#..#.O...OO.O.......##O#..O#...OO.........OOO.#O...O...#OO#O#...O",
		"#O#...#O.##..O.........O#..O..O..........#.##..OOOO...O..#.OO......O#....##....O.O...O.OO...O..OOO.O",
		".OO..#.....O#O..O..OOOO.#.#..#.O......#..#..O..O..O.#O..OO....#...O##O.O.O.O.#......O...O..#.#O...#.",
		"..OO.O......O.O.OO.....O.#..#.O..#O.#.O.....#.O..OO.OOO#..#.#OO.#.OO#.....##O....#.#O.....#..#......",
		".#..O#.#O#.....##.O.#......O#.O.#..#O.#..O#.OO.O.#.O#..O.....#.##..###O..O.OO.##.O.O..#......###...O",
		"#.O.O....O##...#..OO..#...O#O....O....O.OO.........#..OO..#...#...OOO#O.....O.#.OO.OO##.O....#......",
		"O#.O.O..#..#....O.O.......#.#...#...##O..O.O...O.OO#..#O.O..O.......#...#O.#.O.#O##.#.O..#..O....#.#",
		".O.O..#...O#..O....O..#.#.OO......O.O.O.O##........#.O#.#..#....O..O....O...O##OOO..O.#O..O.......##",
		"......#..........O#..O..O..O#.OOO#O.#........#...O.#..OO..#.#.#..O#...###.#...#.#.O#..#O.#O....O..O.",
		"..#.......O...#..O..O#..........#.......#....#...O........##O....#O..OO.###.O#.#O........OO#.#O..O#O",
		".....#.O.O........O.#.O.#.O...##.......#..#........#.O..O...###...O.O.O#....OO.#.....#.......O#..O..",
		"#..O.O.O#.....#.OOO##.....O..OOO.OO...#..#O..OO.O...##...#..###......O...##..#....#..OOO.O...#OO.#..",
		"..##...O...#O..O#.O.#...O.O#......OO..O...O.#...O.O.OO...O.O.O.#..#O.O........#.#..##.O.....#.#O...O",
		"..#O...O.O.#.O....O#.....#..O.O................O....O..##O#.#O....O.O......O...#.....O..O#...O..#.O.",
		"..#..O.#......O#.O.O...O#...O#.O#...O..O....#..O...#.#......#......#O....O.O.O..O..O..#....O..O.....",
		"..#...O....OOO.O##..O#.#.......O##.....O........OO.O...OO#.....#......O.#O#OO##O....O#.OO........#..",
		"....O##..#.....##.O##...O.#..O...#O.#...#.#...#O....O.O.....O.OO.##..#O...#O.O#.#..O#...............",
		"O..O.......O#O.......OO....O.OO....#...#O.OO#..##O....#..#.#O#OO##O...O.O...#...OO#.....O..#..OO....",
		"..#..#O....#...O.##...#O#.O..O.O....O.....OO#O.#..#.....O..O.###.OO.O.......O..#.......OO.#.#..O....",
		"#.O..O.O.#O..OOO##OO....O.OO#OO...#OO.#.#O.OO.OO...#..O.....O....O.O.O#.OO..O.#..#.O.O..OOO##..O#OO.",
		"O.#O.#.OO.#...#.##..OO...#.O##..#O.#.#.#.....O.O##..O..#..O.O.....#....O...O......O..O##.###..#O.#OO",
		"#.........#....O...O........OOOO.##.O.O.#......O####O..#..O##.O.#O.#O...#..#.....O...#.#..O#.#...#O.",
		".......#.#OOO.....O#O..#.#....#.OO.....O..OOOO#...#.............###.........O.OO..#.O#O..O...O..#.O.",
		"O#...O.......O.O..O.#.O.....OO...O...O.O#OO....##.##.#O#..........O.O.......#.....O.....O.O..OOO#..#",
		"#.O#......O.O..#O.O..O....#..#O..#.#O.O...OO...##.#.O.O..#..#O.#.#.O#.......O..O.#...#O.....#.O#...O",
		"..O..O#...##.#..##..#O..O.OO#..#.##O.O..O..#.O..##.O#.O..#.O.#...........O#.O...O.O#.#...#..#..O.O.O",
		"#.#.O...OO....O..O.#O.#.OO..O#..O..O....O.#.#O.##O.....#.O.#...#...#....#....OO.......#.O..O...#..#.",
		"..O.O#O..O.......#.......O#OO..O....O##.....O.#.O..#.#.....#...OOO.......#...OO#......O..........O.O",
		"..##.#..O.O..#....##..O...O..O#.#O....O##....##O.....OO.O..#.O.O#.#.....OO..O........#........O#O###",
		"..#.#...O#..O.......#.OO##...OO#O...#.O....#OOO......O.O.O.#O..#.....##..#O#.OOO...O......OOO#.O.O.O",
		"O...#...O..#O......O.O......#.....#...#..#.O...O.#O##..#.#...#..O#....O.#OO#..O.....#.#..#OOO#...OOO",
		"...##..#O...#...O.##O.O#.........O........OO#O..#.OO.#....O..#...OO..#O.......O..O.O#...#......OO..#",
		".....O.O....O.O#....#O#......O.............O.OO#.....#OO...........O#.......#......OO#O#.OO..O..O...",
		"O##.....O...#.#.#..O..OOOOO..#.#.O##.O...#..O...O.....O.#.#...###O.O..#.#.O..O#....#.O...O#.O.......",
		"#OO.....O.O.#...#.O#.....OO...##.O...O......O.#O.........#.O....OO.OO.#.O.#O.O......#O.###..#.#.OO#.",
		".OOO..O.OO#O..O##.......##...##......O..O...#...O...#.O.O.#.#O.#OO...O.O..#.O#...O..#...OO....#OO...",
		".##.O.OOO..O#.##O...#O#O...#...#.....#.....O...O......#..##O.OOO.O#..O.#..#.#..O........O..#.O#O....",
		"....O..O.............#......O.#O..OO...OO.....O..#..#......OO......O....OO..OOOO.....#.#O..#.#...O.O",
		".###.....##.#......###....O#....O....O...O....#.#O.O#..#..O..#O...OO..OO.O.....#O.O..#.#.....O...O#.",
		"..O..#O...O.OO.O........O...O.O.....O..#O.O.O...#........O........OO.O..OO#O.#.O..#.#O#...O.#.##.O..",
		"OO..OO.O..#OO..OO#..........#...............#.O#....O...O.#.O.O#..#.#.OO#...O.#O..##........O#.O....",
		"#.O......#....#...O.O...O....OO..O..#...O..O..#..O#........###...#..#.#..OO..O.....#.....#...#...O..",
		"..#O.O..O.....OO...............O....#.....O.....#O.OO.O..O..##O....O..#.#.#O...O#O##O#.#OOOOOO.#..O.",
		"#O#O.....O..#...O#......OO...#..OOO#O..O##.OO..#.#.##.#.OO..O.###.O.#..O...#O#....O....O##O..#O.O.O.",
		"..O.....O#.O#.....OOO.O..O......O.O..#O..O.O.........#.###..#.O.O.O.OO#.O..#...O.#..#.#...O##.#...O.",
		"#..O#..O#..#...#.#......##.....#OO.O#.OO..#.....OO#...O.O..O.......O#.#.........O.#O#O.O..#.O.......",
		"..O.#..#..O.O..##..#.#O.......OOO#.O.#.OO.##...O...........#.O..#...#.O........O..O..#.....O.#.....O",
		"O.....O.....O...O....#..#......OO#....O......O..O#.#...O...##.O####O#..O.OO.O.O.O.##...........O...O",
		"OO#.O..##O.O..O..#.#..#...#...O.O.#..#...#.#...O.#..#O......#####..##.OOOOO......O.O..#OO.O.#..O.O..",
		"O.O..#.O.OO.#.O.#....O.#......O.##..#O...O..##.#...O.#...O.O....#...O.#...O.......###....#....#.....",
		"#..#.........O....#........OO.#..OO..OOO....#.....O....##O.O.O....O.O..O....#O#.###.O...#..O.O....#O",
		"#...O##..O#..O..#.O.......O.....#....#O###O#OO.O.O..O...O..OO...#.#.#..O..O...O..#...#.#....#.#.#.OO",
		"......#..........#...#..O..#O#..O....#.O..O..#.O.#O...O#.......#........O#O....OO.....O..OO...#..#..",
		"O..O...#...#..O##....#.O....O..#...O..###..O...O....OO..O..O.........O..##OO......#O...O.O##...OOO.O",
		"..#......OO.#.......#...O#.#.O#...#...#O...O......#.....O#....OO........O#O.......#O..#OO#.O.O....#.",
		".O.##....#OO.O...O#....O..O......O..#.....#O..O.#...#.......O....OO#...O....#....#.O...OOOOO.#O.O..O",
		"....#....OO#..#O....O...#.OO#...#O......#.#.OOO..#...O#......#...O..##.OO..O....O.#O...OO.#.#.O..O..",
		"O.O....OO#...#..O...O.#OO.O..O.....#....#.O..O.#..........OO..#.#.#.O...O.O.#.#..#.#......#..O.###..",
		"....#..O.O....OO...O.....#.......#........O..#.OO##O#OOO#O.......#.#.#....#.#.O..#.O##.#..O....##.OO",
		"O.....#.#....#.........O.O....O..O#.....#O.O#O...O#O...O...##........OO.O.........OO..#...OO.O..#...",
		".O....O....#.#O..O...OOO.......O..O..O..#............O.....#.O...OO....#.O#O.O##.#.O........#...O...",
		"..##.O..OO#..O.O...O....##...O.......#..O...O.#O.O..........O.#.#OOO#..O...O..O.OOO.....O.........O.",
		"OO..#..#.....#....##...#..#..##.O.OO...O#........O..O.....#.......O.O.O..O...O.O.OO..O...#.###.#.O##",
		"O.#O...O......O.......O#.O......O#..O..#...##.O...##......O.O.#O.##.O.O#....O#O#..#..#O#...#........",
		"#OOO.O...O...O.#.#O....#....O..#.#..O.......#OO......O........O#..O..##..#.O...O...##O...........O#O",
		".O.#.OO..O.....O...##O.#O.#...O#..O.#....#..#OO#.#.OO...#.O#...OO......O.....#O..O...O...OOO#.O..O##",
		"..O...O...#..O..O#..O.OO.#.OO.O...OO.#.....#.O..#..#O..OOO.##O#O....#...#O.O....OO..#.O...OO#O#....O",
		"#O..OO.###.OOOO....O#......O.O.#OO......OO..O.O..O...O...O..#.OO.....##.O.OO...O..#O.###..O..OO.#.O.",
		"...O.O##....O.#..O#.OOO#.........#....#....#.......#.O...O...........O.O#O.O#..OO.O#....O.#.O.#O.O..",
		"#....#....O..O..#.....#..O#O.#..O.O.#O.O....OO.#.#.O.O....#O....#..O..#.O.#O...#.....O.O...#..#...O.",
		".#....O.....O..OO.O.O#..#...##.....OO....#.#.OO.#..O..O.......#......#.....OOO#O.#######.##......O#O",
		"#OO..#..OOOO.O..#.###..#.O..#.O...O#.O..#....#...#OO#..#....#OO....#...OO..#...O.O.....OO...#.O.#...",
		"#..O....OO....#.#.OO...O.O....O.##..#.O..O....OO..#.O..O.OO......#.#O..O.O........O.O.O...O.....O#.O",
		".O.#...O.#....O..O#.O..O.OO.....O...........#.##OO....O..#O###..#..O.O.O.....#.......##O#.#...O.....",
		".....O.OO.#O.O##............O.O...##.O......#....#....#...#.#...O...#OO.#.O.....O.#..O...........#..",
		"..O.OO..#O..O.O.##.O#..........OO..O...OO..OO.OO.OO#...O.#O.O...OO.......#OO.#..#..O...#....O.......",
		"O.OO....O#....#...#.#O##..O#...#..#......O.......#O..#O..O..O...#......#O..O##.O....O..O.O#.#....O..",
		".#.O..O..O.........##O..#..OOO...O##.O.#O#.O...O.#....#..O..O.......O..#..O..O#..#....O..OO.##..O...",
		"...#...OO.OOO..#...O.#.O.......#...##.O#.##..##..#..#.##.....##O..O..O..#OO...#....O..O.......#..#O.",
		".....O...O.#......##O.O..#...#.....O..O....#O#.#..#O.#....O#.O...#O....###...OOO...O#..O..O..###..O.",
		"..#..........O..O..O...O...O....O............O.....#O....O..#.....#...O.....#..O.#.#..O..O...O.#O...",
		"...OOO......O...#.O#....##..#...O..O#..##..O....#.O.OO.......#....O....##.#.#.O...O#.#O..OOO..O.O...",
		"O...#.O....O..OO..O...O....#.#.....#O#O....#O.##.O.OO.#.O............#.O#...OO..OO#..#O.#....O......",
		"#.O.....#.O#.OO.#........O..O.#.OO.........#..#...#.O#O...O...O..OO.#.......O.O..#..O.O...OO..O..O.O",
		"..#....#O#.O.O#.OO..O...#..O......O#O.......OO..OO.....#O#.O.O..#O.......OO.#.O..O....O.O.O.........",
		"#...O###O..O.O##..#.#.##..#...#.#.#...O.#....#.....#.#...O#..O.O...O..O........#.###....O...#..#...#",
		".......O........O.....##..O##...O...OO..#O..O...#...O.........#OO.....OO.O..O..#OO.#.#O..O#.OO...#O.",
		".OO..O.....#.O.....#O..#O...O....#..O.....#....#O......O...#.O###.OO......O..O.......#####.........O",
	}
)

func moveRocks(input [][]byte, delta int, vertical bool) int {
	yMax := len(input)
	xMax := len(input[0])

	moveCount := 0
	if vertical {
		for x := 0; x < xMax; x++ {
			for y := 0; y < yMax; y++ {
				curY := y
				if delta > 0 {
					curY = yMax - y - 1
				}
				if input[curY][x] != 'O' {
					continue
				}
				rockY := curY
				for rockY+delta >= 0 && rockY+delta < yMax && input[rockY+delta][x] == '.' {
					moveCount++
					input[rockY+delta][x] = 'O'
					input[rockY][x] = '.'
					rockY += delta
				}
			}
		}
	} else {
		for x := 0; x < xMax; x++ {
			for y := 0; y < yMax; y++ {
				curX := x
				if delta > 0 {
					curX = xMax - x - 1
				}
				if input[y][curX] != 'O' {
					continue
				}
				rockX := curX
				for rockX+delta >= 0 && rockX+delta < xMax && input[y][rockX+delta] == '.' {
					moveCount++
					input[y][rockX+delta] = 'O'
					input[y][rockX] = '.'
					rockX += delta
				}
			}
		}
	}

	return moveCount
}

func score(input [][]byte) int {
	yMax := len(input)
	xMax := len(input[0])
	score := 0
	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax; y++ {
			if input[y][x] == 'O' {
				score += (yMax - y)
			}
		}
	}
	return score
}

func toString(input [][]byte) string {
	var buf strings.Builder
	for _, line := range input {
		buf.Write(line)
		buf.WriteByte('\n')
	}
	return buf.String()
}
func displayInput(input [][]byte) {
	fmt.Printf("%s\n", toString(input))
}

func computeScore(lines []string) {
	inputBytes := make([][]byte, len(lines))
	for i, line := range lines {
		inputBytes[i] = []byte(line)
	}
	displayInput(inputBytes)
	moveRocks(inputBytes, -1, true)
	displayInput(inputBytes)
	s := score(inputBytes)
	fmt.Printf("Score is: %d\n", s)
}

func TestDay14Phase1(t *testing.T) {
	// 136
	computeScore(lightInput)
	// 112773
	computeScore(largeInput)
}

func computeFullTiltScore(lines []string) {
	inputBytes := make([][]byte, len(lines))
	for i, line := range lines {
		inputBytes[i] = []byte(line)
	}
	displayInput(inputBytes)
	snapshot := make([][]byte, len(lines))
	for i, line := range lines {
		snapshot[i] = []byte(line)
	}
	moves := make(map[string]string, 10_000_000)
	cur := toString(inputBytes)
	cycle := 0
	maxCycle := 1_000_000_000
	//maxCycle = 3
	for ; cycle < maxCycle; cycle++ {
		moveRocks(inputBytes, -1, true)
		moveRocks(inputBytes, -1, false)
		moveRocks(inputBytes, 1, true)
		moveRocks(inputBytes, 1, false)
		newPos := toString(inputBytes)
		moves[cur] = newPos
		cur = newPos
		fmt.Println(cur)
		if moves[cur] != "" {
			cycle++
			break
		}
	}

	if cycle < maxCycle {
		// Compute cycle length
		cycleLength := 1
		currentPos := cur
		for moves[currentPos] != cur {
			cycleLength++
			currentPos = moves[currentPos]
		}

		// Compute number of full "cycle" before going over the limit
		skeptCycle := (maxCycle - cycle) / cycleLength
		// Skip complete cycles
		cycle += skeptCycle * cycleLength

		// Iterates till final state
		for ; cycle < maxCycle; cycle++ {
			cur = moves[cur]
			k := cur
			fmt.Printf("%s\n", k)
			inputBytes = make([][]byte, len(lines))
			k = strings.TrimSpace(k)
			for i, line := range strings.Split(k, "\n") {
				inputBytes[i] = []byte(line)
			}
			s := score(inputBytes)
			fmt.Printf("Score for cycle %d is: %d\n", cycle, s)
		}
	}
	k := cur
	fmt.Printf("%s\n", k)
	inputBytes = make([][]byte, len(lines))
	k = strings.TrimSpace(k)
	for i, line := range strings.Split(k, "\n") {
		inputBytes[i] = []byte(line)
	}
	s := score(inputBytes)
	fmt.Printf("Score for cycle %d is: %d\n", cycle, s)
}

func TestDay14Phase2(t *testing.T) {
	// 64
	computeFullTiltScore(lightInput)
	// 98894
	computeFullTiltScore(largeInput)
}
