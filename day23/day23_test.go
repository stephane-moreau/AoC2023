package day23

import (
	"fmt"
	"strings"
	"testing"
)

var (
	lightInput = []string{
		"#.#####################",
		"#.......#########...###",
		"#######.#########.#.###",
		"###.....#.>.>.###.#.###",
		"###v#####.#v#.###.#.###",
		"###.>...#.#.#.....#...#",
		"###v###.#.#.#########.#",
		"###...#.#.#.......#...#",
		"#####.#.#.#######.#.###",
		"#.....#.#.#.......#...#",
		"#.#####.#.#.#########v#",
		"#.#...#...#...###...>.#",
		"#.#.#v#######v###.###v#",
		"#...#.>.#...>.>.#.###.#",
		"#####v#.#.###v#.#.###.#",
		"#.....#...#...#.#.#...#",
		"#.#########.###.#.#.###",
		"#...###...#...#...#.###",
		"###.###.#.###v#####v###",
		"#...#...#.#.>.>.#.>.###",
		"#.###.###.#.###.#.#v###",
		"#.....###...###...#...#",
		"#####################.#",
	}

	largeInput = []string{
		"#.###########################################################################################################################################",
		"#.....###...###...#...###...#...###...#...###...###...#...#.........###...###...#.......#...###...#...#...###.....#...#.....#...#...#...#...#",
		"#####.###.#.###.#.#.#.###.#.#.#.###.#.#.#.###.#.###.#.#.#.#.#######.###.#.###.#.#.#####.#.#.###.#.#.#.#.#.###.###.#.#.#.###.#.#.#.#.#.#.#.#.#",
		"#.....#...#...#.#.#.#.#...#.#.#.#...#.#.#...#.#...#.#.#.#.#.#.......#...#...#.#.#.#.....#.#.#...#.#.#...#...#...#.#.#.#...#.#.#.#.#.#.#.#.#.#",
		"#.#####.#####.#.#.#.#.#.###.#.#.#.###.#.###.#.###.#.#.#.#.#.#.#######.#####.#.#.#.#.#####.#.#.###.#.#######.###.#.#.#.###.#.#.#.#.#.#.#.#.#.#",
		"#.#...#...#...#.#.#.#.#...#.#.#.#...#.#.#...#...#.#.#.#.#.#.#.#...###.....#.#.#...#...#...#.#...#.#.....#...#...#.#.#.#...#...#.#.#.#.#.#.#.#",
		"#.#.#.###.#.###.#.#.#.###.#.#.#.###.#.#.#.#####.#.#.#.#.#.#.#.#.#.#######.#.#.#######.#.###.###.#.#####.#.###.###.#.#.#.#######.#.#.#.#.#.#.#",
		"#.#.#.#...#...#.#...#.#...#...#.#...#.#.#.>.>.#.#.#.#.#.#.#.#...#.>.>.#...#.#.#.......#.#...#...#.#.>.>.#.....#...#.#.#.#.......#.#.#.#.#.#.#",
		"#.#.#.#.#####v#.#####.#.#######.#.###.#.###v#.#.#.#.#.#.#.#.#######v#.#.###.#.#.#######.#.###.###.#.#v#########.###.#.#.#.#######.#.#.#.#.#.#",
		"#...#.#.#...#.>.#.....#.....#...#.#...#.#...#...#.#.#.#.#.#.#.....#.#.#...#...#.#...#...#.....###.#.#.#.....###.....#.#.#.....###.#.#.#.#.#.#",
		"#####.#.#.#.#v###.#########.#.###.#.###.#.#######.#.#.#.#.#.#.###.#.#.###.#####.#.#.#.###########.#.#.#.###.#########.#.#####.###.#.#.#.#.#.#",
		"#...#...#.#.#...#...#...#...#...#.#.#...#.......#.#.#.#.#.#.#...#...#.#...###...#.#.#.....#...###...#...###.........#.#.#...#...#.#.#.#.#.#.#",
		"#.#.#####.#.###.###.#.#.#.#####.#.#.#.#########.#.#.#.#.#.#.###.#####.#.#####.###.#.#####.#.#.#####################.#.#.#.#.###.#.#.#.#.#.#.#",
		"#.#...###.#.....#...#.#.#...###...#.#.#.........#...#.#.#...###.....#...#...#.....#.....#...#.......#.......#.......#...#.#.....#.#.#.#.#.#.#",
		"#.###.###.#######.###.#.###.#######.#.#.#############.#.###########.#####.#.###########.###########.#.#####.#.###########.#######.#.#.#.#.#.#",
		"#...#...#...#...#.....#.#...#.......#.#.#.......#...#...###.........#.....#.............#.........#.#.....#...#...#...###.#.....#.#.#.#.#.#.#",
		"###.###.###.#.#.#######.#.###.#######.#.#.#####.#.#.#######.#########.###################.#######.#.#####.#####.#.#.#.###.#.###.#.#.#.#.#.#.#",
		"###...#.###...#.....###.#...#.........#...#...#...#...#.....#.......#.....#.....#...>.>.#.......#...#...#...#...#...#...#.#.###...#...#...#.#",
		"#####.#.###########.###.###.###############.#.#######.#.#####.#####.#####.#.###.#.###v#.#######.#####.#.###.#.#########.#.#.###############.#",
		"#...#.#...........#...#.....###...#.........#.......#.#.......#...#.#.....#.#...#.#...#...#...#...#...#.....#.#.......#.#.#.#.............#.#",
		"#.#.#.###########.###.#########.#.#.###############.#.#########.#.#.#.#####.#.###.#.#####.#.#.###.#.#########.#.#####.#.#.#.#.###########.#.#",
		"#.#...#.........#.#...#.....#...#.#...............#.#.#.........#...#.......#.#...#.#...#.#.#.#...#.......#...#.....#.#.#.#.#...#...#...#...#",
		"#.#####.#######.#.#.###.###.#.###.###############.#.#.#.#####################.#.###.#.#.#.#.#.#.#########.#.#######.#.#.#.#.###.#.#.#.#.#####",
		"#.#...#.#.......#...###...#.#...#.#...............#...#.........#.....#.....#...#...#.#.#...#.#.#.........#.#.......#...#...###...#...#.....#",
		"#.#.#.#.#.###############.#.###.#.#.###########################.#.###.#.###.#####.###.#.#####.#.#.#########.#.#############################.#",
		"#...#...#.....#...#...#...#.....#.#...........#...###...###.....#.#...#.#...#...#.....#.....#.#.#...#.......#.............###...###...#.....#",
		"#############.#.#.#.#.#.#########.###########.#.#.###.#.###.#####.#.###.#.###.#.###########.#.#.###.#.###################.###.#.###.#.#.#####",
		"#.............#.#.#.#.#...#.......#...#.......#.#...#.#...#.......#.....#.#...#.............#...###...#...................#...#.#...#.#.....#",
		"#.#############.#.#.#.###.#.#######.#.#.#######.###.#.###.###############.#.###########################.###################.###.#.###.#####.#",
		"#.........#...#.#.#.#.###.#.#...###.#.#.....###.#...#...#.#...#...........#.#...#...###...#...###.....#.....#.......#...#...#...#...#.#...#.#",
		"#########.#.#.#.#.#.#.###.#.#.#.###.#.#####v###.#.#####.#.#.#.#v###########.#.#.#.#.###.#.#.#.###.###.#####.#.#####.#.#.#.###.#####.#.#v#.#.#",
		"#.........#.#.#.#...#...#.#.#.#...#.#.....>.>.#.#...#...#.#.#.>.>.........#...#...#.....#...#...#...#.#...#...#.....#.#.#...#.#.....#.>.#...#",
		"#.#########.#.#.#######.#.#.#.###.#.#######v#.#.###.#.###.#.###v#########.#####################.###.#.#.#.#####.#####.#.###.#.#.#######v#####",
		"#.#...#...#.#.#...#.....#.#...###...#.......#.#.###.#...#...#...#.........#...#...#...#.........#...#.#.#.#...#...###.#.###.#...###...#.#...#",
		"#.#.#.#.#.#.#.###.#.#####.###########.#######.#.###.###.#####.###.#########.#.#.#.#.#.#.#########.###.#.#.#.#.###.###.#.###.#######.#.#.#.#.#",
		"#.#.#.#.#.#.#.#...#.....#.....###...#.......#.#...#.#...#.....#...###.......#...#...#...#...#...#...#.#.#.#.#.#...#...#.#...#.....#.#.#...#.#",
		"#.#.#.#.#v#.#.#.#######.#####.###.#.#######.#.###.#.#.###.#####.#####.###################.#.#.#.###.#.#.#.#.#.#.###.###.#.###.###.#.#.#####.#",
		"#.#.#...#.>.#.#.#.......#.....#...#...#...#.#...#.#.#.###.....#.#...#...#...........#...#.#.#.#.###.#.#.#.#.#.#...#...#.#.###...#.#.#...#...#",
		"#.#.#####v###.#.#.#######.#####.#####.#.#.#.###.#.#.#.#######.#.#.#.###.#.#########.#.#.#.#.#.#.###.#.#.#.#.#.###v###.#.#.#####.#.#.###.#.###",
		"#...###...#...#.#.....#...#.....#...#...#.#...#.#.#.#.#.......#...#...#...#...#...#.#.#.#.#.#.#...#.#.#.#.#.#.#.>.>...#...#.....#.#.###.#...#",
		"#######.###.###.#####.#.###.#####.#.#####.###.#.#.#.#.#.#############.#####.#.#.#.#v#.#.#.#.#.###.#.#.#.#.#.#.#.#v#########.#####.#.###.###.#",
		"#...#...###...#.#.....#...#...#...#.#...#.....#...#.#.#.#...###...###...#...#.#.#.>.>.#.#.#.#...#.#.#.#.#...#...#.......#...#...#.#...#.#...#",
		"#.#.#.#######.#.#.#######.###.#.###.#.#.###########.#.#.#.#.###.#.#####.#.###.#.###v###.#.#.###.#.#.#.#.###############.#.###.#.#.###.#.#.###",
		"#.#...#.....#...#.........#...#.#...#.#.....#.....#...#...#...#.#.......#.###...#...###.#.#.#...#...#...###...........#.#...#.#.#.....#.#...#",
		"#.#####.###.###############.###.#.###.#####.#.###.###########.#.#########.#######.#####.#.#.#.#############.#########.#.###.#.#.#######.###.#",
		"#.#...#.#...#.......#.....#.....#.....#...#...#...#...........#...#.....#.#...#...#...#...#...#.....#.....#.........#...###...#.....###.....#",
		"#.#.#.#.#.###.#####.#.###.#############.#.#####.###.#############.#.###.#.#.#.#.###.#.#########.###.#.###.#########.###############.#########",
		"#...#...#...#.....#.#...#.....#.........#.....#...#...........###...###.#.#.#.#.....#...#...###.#...#...#.........#.......#.........#...#####",
		"###########.#####.#.###.#####.#.#############.###.###########.#########.#.#.#.#########.#.#.###.#.#####.#########.#######.#.#########.#.#####",
		"#...#.......#...#.#.###...#...#.............#.#...#...........#...#...#.#.#.#.#...#.....#.#.#...#.#...#.........#.........#...........#.....#",
		"#.#.#.#######.#.#.#.#####.#.###############.#.#.###.###########.#.#.#.#.#.#.#.#.#.#.#####.#.#.###.#.#.#########.###########################.#",
		"#.#...#.....#.#.#.#.#...#.#.#...#...........#...###.............#.#.#.#...#.#...#...###...#.#...#.#.#.###...#...#...###...#.....#...........#",
		"#.#####.###.#.#.#.#.#.#.#.#.#.#.#.###############################.#.#.#####.###########.###.###.#.#.#.###.#.#v###.#.###.#.#.###.#.###########",
		"#.#...#...#.#.#.#.#.#.#.#.#.#.#.#...#####...#.....#...#...........#.#.###...#.....#...#...#...#.#.#.#.#...#.>.>.#.#.###.#.#...#.#...........#",
		"#.#.#.###.#.#.#.#.#.#.#.#.#.#.#.###v#####.#.#.###.#.#.#.###########.#.###.###.###.#.#.###.###.#.#.#.#.#.#####v#.#.#.###.#.###v#.###########.#",
		"#.#.#.#...#.#.#.#.#.#.#...#.#.#...>.>.###.#.#.#...#.#.#...........#.#...#...#...#.#.#.#...###...#.#.#.#...#...#...#...#.#.#.>.#.............#",
		"#.#.#v#.###.#.#.#.#.#.#####.#.#####v#.###.#.#.#.###.#.###########.#.###.###v###.#.#.#.#.#########.#.#.###.#.#########.#.#.#.#v###############",
		"#...#.>.#...#.#.#.#.#.....#.#...#...#...#.#.#.#...#.#.#.........#.#...#...>.>.#.#...#.#.###.......#.#...#.#.....#...#.#.#.#.#.........#...###",
		"#####v###.###.#.#.#.#####.#.###.#.#####.#.#.#.###.#.#.#.#######.#.###.#####v#.#.#####.#.###.#######.###.#.#####.#.#.#.#.#.#.#########.#.#.###",
		"#.....###...#.#.#.#.#...#.#.....#.....#...#.#.#...#.#.#.....###...#...###...#...#...#.#...#.#.....#.#...#.#.....#.#.#...#.#.#.........#.#...#",
		"#.#########.#.#.#.#.#.#.#.###########.#####.#.#.###.#.#####.#######.#####.#######.#.#.###.#.#.###.#.#.###.#.#####.#.#####.#.#.#########.###.#",
		"#.....#...#...#...#.#.#.#.#...........###...#.#...#.#...#...#...###.....#.#...###.#.#.....#...#...#.#.#...#.......#.#...#.#.#.....#...#.#...#",
		"#####.#.#.#########.#.#.#.#.#############.###.###.#.###.#.###.#.#######.#.#.#.###.#.###########.###.#.#.###########.#.#.#.#.#####.#.#.#.#.###",
		"#...#...#.......#...#.#...#.....#...#...#.#...#...#.#...#...#.#.#...#...#.#.#.....#...........#.....#...#...#.....#...#.#.#.#.....#.#...#...#",
		"#.#.###########.#.###.#########.#.#.#.#.#.#.###.###.#.#####v#.#.#.#.#.###.#.#################.###########.#.#.###.#####.#.#.#.#####.#######.#",
		"#.#.............#.....#.........#.#...#.#.#.#...#...#.....>.>.#.#.#...###.#.#.......#.....#...#...........#...###.#.....#...#.....#.#.......#",
		"#.#####################.#########.#####.#.#.#.###.#########v###.#.#######.#.#.#####.#.###.#.###.#################.#.#############.#.#.#######",
		"#.#...#...#...#####...#...#...#...###...#.#.#...#.#.....#...###.#.#.....#...#.....#.#...#...#...#...#...........#.#.....#.......#...#.......#",
		"#.#.#.#.#.#.#.#####.#.###.#.#.#.#####.###.#.###.#.#.###.#.#####.#.#.###.#########.#.###.#####.###.#.#.#########.#.#####.#.#####.###########.#",
		"#.#.#.#.#.#.#...#...#...#...#...#...#...#...#...#.#...#...#...#...#.#...#...###...#.#...#...#.....#...#.........#.......#.#...#.............#",
		"#.#.#.#.#.#.###.#.#####.#########.#.###.#####.###.###.#####.#.#####.#.###.#.###.###.#.###.#.###########.#################.#.#.###############",
		"#...#...#...#...#.#.....#...#...#.#...#.....#.#...#...#...#.#.#...#.#.#...#.#...###...###.#.#...#...###.................#...#...............#",
		"#############.###.#.#####.#.#.#.#.###.#####.#.#.###.###.#.#.#.#.#.#.#.#.###.#.###########.#.#.#.#.#.###################.###################.#",
		"#.............#...#.....#.#.#.#.#...#.#.....#...###.....#.#.#.#.#.#.#.#.#...#...#####...#.#.#.#.#.#.#...#...............###...#.............#",
		"#.#############.#######.#.#.#.#.###.#.#.#################.#.#.#.#.#.#.#.#.#####v#####.#.#.#.#.#.#.#.#.#.#.#################.#.#.#############",
		"#.....#.....###.#.......#.#.#.#...#.#.#...#...#...#...#...#.#...#.#.#.#.#.#...>.>.###.#.#.#...#.#.#.#.#.#.....#...#...#...#.#.#.............#",
		"#####.#.###.###.#.#######.#.#.###.#.#.###.#.#.#.#.#.#.#.###.#####.#.#.#.#.#.###v#.###.#.#.#####.#.#.#.#.#####.#.#.#.#.#.#.#.#.#############.#",
		"#.....#.###...#.#.#...###.#.#...#.#.#.#...#.#.#.#.#.#.#...#.....#.#.#.#.#...#...#...#.#.#.....#.#.#.#.#.#...#.#.#...#.#.#.#.#.#.............#",
		"#.#####.#####v#.#.#.#.###.#.###.#.#.#.#v###.#.#.#.#.#.###v#####.#.#.#.#.#####.#####.#.#.#####.#.#.#.#.#.#.#.#v#.#####.#.#.#.#.#.#############",
		"#.#...#.#...#.>.#...#.#...#...#.#.#.#.>.>...#.#.#.#.#.#.>.>.....#.#.#.#...#...#####.#.#.#.....#.#.#.#.#.#.#.>.>.#.....#.#.#.#.#.........#####",
		"#.#.#.#.#.#.#v#######.#.#####.#.#.#.###v#####.#.#.#.#.#.#v#######.#.#.###.#.#######.#.#.#.#####.#.#.#.#.#.###v###.#####.#.#.#.#########.#####",
		"#.#.#...#.#...#...#...#.....#...#...#...#...#...#...#.#.#.#####...#.#.#...#.....#...#.#.#...#...#.#.#.#.#.###.#...#...#.#.#.#.#...#...#.....#",
		"#.#.#####.#####.#.#.#######.#########.###.#.#########.#.#.#####.###.#.#.#######.#.###.#.###.#.###.#.#.#.#.###.#.###.#.#.#.#.#.#.#.#.#.#####.#",
		"#...###...#...#.#.#...#...#.#.......#...#.#.......###...#.....#.....#.#.#.......#.....#.#...#...#.#...#...#...#...#.#...#.#.#.#.#.#.#.#...#.#",
		"#######.###.#.#.#.###.#.#.#.#.#####.###.#.#######.###########.#######.#.#.#############.#.#####.#.#########.#####.#.#####.#.#.#.#.#.#.#v#.#.#",
		"#.....#.....#.#.#...#.#.#.#.#.#...#...#...#...#...#...........###...#...#.........#...#.#.#.....#.#.........#.....#.....#.#.#.#.#...#.>.#.#.#",
		"#.###.#######.#.###.#.#.#.#.#.#.#.###.#####.#.#.###.#############.#.#############.#.#.#.#.#.#####.#.#########.#########.#.#.#.#.#######v#.#.#",
		"#...#.....#...#.#...#.#.#...#.#.#...#...#...#.#.#...#.....#.....#.#...#...........#.#.#...#.....#.#...#.....#...#...#...#.#.#.#.#.......#...#",
		"###.#####.#.###.#.###.#.#####.#.###.###.#.###.#.#.###.###.#.###.#.###.#.###########.#.#########.#.###.#.###.###.#.#.#.###.#.#.#.#.###########",
		"###.....#...#...#...#.#.....#...###...#.#...#.#.#...#...#...###...#...#.............#.#.....###...###...###...#.#.#.#.#...#.#...#...........#",
		"#######.#####.#####.#.#####.#########.#.###.#.#.###.###.###########.#################.#.###.#################.#.#.#.#.#.###.###############.#",
		"#.......#...#.#.....#.......###...#...#.#...#...#...#...#.....#.....#...#.......#.....#...#.#.........#.......#...#...#.....#.....###.......#",
		"#.#######.#.#.#.###############.#.#.###.#.#######.###.###.###.#.#####.#.#.#####.#.#######.#.#.#######.#.#####################.###.###.#######",
		"#...#.....#...#...........#.....#...###...###...#.....###...#...###...#.#.....#.#.....#...#.#.#.......#.....#####...###...###...#.#...#######",
		"###.#.###################.#.#################.#.###########.#######.###.#####.#.#####.#.###.#.#.###########.#####.#.###.#.#####.#.#.#########",
		"#...#.###...###.........#.#...#.....###.....#.#.#...#.....#.......#.#...###...#...#...#.#...#.#.#...#...###...#...#...#.#.....#.#.#.........#",
		"#.###.###.#.###.#######.#.###.#.###.###.###.#.#.#.#.#.###.#######.#.#.#####.#####.#.###.#.###.#.#.#.#.#.#####.#.#####.#.#####.#.#.#########.#",
		"#.....#...#...#...#...#...###...###...#...#.#.#.#.#.#...#.#...#...#.#.#...#.#####...#...#.#...#.#.#.#.#.#.....#...#...#.#.....#.#.###.......#",
		"#######.#####.###v#.#.###############.###.#.#.#.#.#.###.#.#.#.#v###.#.#.#.#v#########.###.#.###.#.#.#.#.#.#######.#.###.#.#####.#.###v#######",
		"#.......#...#...#.>.#...#...###...#...###.#.#.#.#.#.#...#...#.>.>.#.#.#.#.>.>.###...#...#.#.#...#.#.#.#.#...#...#.#...#.#...#...#...>.#...###",
		"#.#######.#.###.#v#####.#.#.###.#.#.#####.#.#.#.#.#.#.#########v#.#.#.#.###v#.###.#.###.#.#.#.###.#.#.#.###v#.#.#.###.#.###.#.#######v#.#.###",
		"#.........#...#...#.....#.#.###.#...###...#...#.#.#.#.#.........#...#.#...#.#...#.#.#...#.#.#...#.#.#.#.#.>.>.#.#...#.#...#...#.....#...#...#",
		"#############.#####.#####.#.###.#######.#######.#.#.#.#.#############.###.#.###.#.#.#.###.#.###.#.#.#.#.#.#v###.###.#.###.#####.###.#######.#",
		"#...#...#...#...#...#...#.#...#.....#...#.......#.#...#...........#...#...#...#.#.#.#.#...#.#...#.#...#...#.###...#.#.#...#...#...#.#.....#.#",
		"#.#.#.#.#.#.###.#.###.#.#.###.#####.#.###.#######.###############.#.###.#####.#.#.#.#.#.###.#.###.#########.#####.#.#.#.###.#.###.#.#.###.#.#",
		"#.#...#...#.....#...#.#.#.#...#...#.#...#...#...#.#...............#...#...###.#...#...#...#.#.....#...#...#.#...#.#.#...#...#.....#.#...#.#.#",
		"#.#################.#.#.#.#.###.#.#v###.###.#.#.#.#.#################.###.###.###########.#.#######.#.#.#.#.#.#.#.#.#####.#########.###.#.#.#",
		"#.......#.....#...#.#.#...#...#.#.>.>.#...#.#.#...#.#.............#...#...#...#...###...#.#.#.....#.#.#.#.#...#.#.#...#...#.......#...#.#.#.#",
		"#######.#.###.#.#.#.#.#######.#.###v#.###.#.#.#####.#.###########.#.###.###.###.#.###.#.#.#.#.###.#.#.#.#.#####.#.###.#.###.#####.###.#.#.#.#",
		"#.......#.#...#.#.#...#...###.#.#...#...#.#...#.....#.#...........#.#...###.....#.....#.#...#...#...#...#...#...#...#.#.....#.....###...#...#",
		"#.#######.#.###.#.#####.#.###.#.#.#####.#.#####.#####.#.###########.#.#################.#######.###########.#.#####.#.#######.###############",
		"#.#...#...#.....#...#...#...#...#.....#.#.....#.......#.....#.....#...###...#.........#.....###...........#...#...#...###...#...###.........#",
		"#.#.#.#.###########.#.#####.#########.#.#####.#############.#.###.#######.#.#.#######.#####.#############.#####.#.#######.#.###.###.#######.#",
		"#.#.#.#.#...........#.#.....###.......#...#...#...#.........#.#...###.....#...#...###.......#.............#...#.#.###.....#...#.....#...#...#",
		"#.#.#.#.#.###########.#.#######.#########.#.###.#.#.#########.#.#####.#########.#.###########.#############.#.#.#.###.#######.#######.#.#.###",
		"#...#...#.............#...#...#.#.....###...#...#.#...........#.....#.#.........#.......#.....#.....#.......#...#...#...#...#.........#...###",
		"#########################.#.#.#.#.###.#######.###.#################.#.#.###############.#.#####.###.#.#############.###.#.#.#################",
		"#.........#.....#...#...#...#.#.#.#...#...#...#...#.................#...#...#...........#.......###...###.......#...###...#.........#...#...#",
		"#.#######.#.###.#.#.#.#.#####.#.#.#.###.#.#.###.###.#####################.#.#.###########################.#####.#.#################.#.#.#.#.#",
		"#.......#.#...#...#...#.#...#.#...#.....#...#...###.................###...#...#...#...#...###.............#...#...#...#...........#...#...#.#",
		"#######.#.###.#########.#.#.#.###############.#####################.###.#######.#.#.#.#.#.###.#############.#.#####.#.#.#########.#########.#",
		"#.......#...#.........#.#.#.#...#...#.........#...#...........#...#.#...#.......#.#.#.#.#...#.....#...#...#.#.#...#.#.#.........#.#...#.....#",
		"#.#########.#########.#.#.#.###.#.#.#.#########.#.#.#########.#.#.#.#.###.#######.#.#.#.###.#####.#.#.#.#.#.#.#.#.#.#.#########.#.#.#.#.#####",
		"#.........#...........#...#.....#.#...###...###.#.#.........#.#.#.#.#.....###.....#.#.#.#...#...#.#.#.#.#...#.#.#.#.#...#...#...#.#.#.#.....#",
		"#########.#######################.#######.#.###.#.#########.#.#.#.#.#########v#####.#.#.#.###.#.#.#.#.#.#####.#.#.#.###.#.#.#.###.#.#.#####.#",
		"#...#...#...............#.........#.......#...#.#.#...#.....#...#...#.......>.>.#...#...#...#.#.#...#...#.....#.#.#.#...#.#.#...#...#.......#",
		"#.#.#.#.###############.#.#########.#########.#.#.#.#.#.#############.#########.#.#########.#.#.#########.#####.#.#.#.###.#.###v#############",
		"#.#...#.................#...........#...#.....#.#...#.#.......#...#...#.....#...#...#.......#.#...#...#...#...#.#...#.#...#...>.###...#...###",
		"#.###################################.#.#.#####.#####.#######.#.#.#.###.###.#.#####.#.#######.###.#.#.#.###.#.#.#####.#.#######v###.#.#.#.###",
		"#.#...#.....#...###.....#...#...#.....#.#...###.#...#.#...#...#.#.#.....#...#...#...#...#...#.#...#.#.#...#.#.#...#...#.###...#.###.#.#.#.###",
		"#.#.#.#.###.#.#.###.###.#.#.#.#.#.#####.###.###.#.#.#.#.#.#.###.#.#######.#####.#.#####.#.#.#.#.###.#.###v#.#.###.#.###.###.#.#.###.#.#.#.###",
		"#...#...###.#.#...#.#...#.#.#.#.#.....#...#.#...#.#...#.#.#.#...#.###...#.....#.#.#####.#.#.#.#...#.#.#.>.>.#.#...#...#.#...#...#...#.#.#.###",
		"###########.#.###.#.#.###.#.#.#.#####.###.#.#.###.#####.#.#.#.###.###.#.#####.#.#.#####.#.#.#.###.#.#.#.#####.#.#####.#.#.#######.###.#.#.###",
		"#...........#...#...#.....#.#.#.#.....###...#...#.#...#.#.#.#...#.#...#.#...#.#...#.....#.#.#...#.#.#.#.....#.#.....#...#...#...#...#.#.#...#",
		"#.#############.###########.#.#.#.#############.#.#.#.#.#.#v###.#.#.###.#.#.#.#####.#####.#.###.#.#.#.#####.#.#####.#######.#.#.###.#.#.###.#",
		"#.#...#.....#...###.........#.#.#.....#...#...#.#.#.#...#.>.>.#.#.#.###.#.#.#.....#.....#.#...#.#.#.#.#...#.#.......#.......#.#...#.#.#...#.#",
		"#.#.#.#.###.#.#####.#########.#.#####v#.#.#.#.#.#.#.#########.#.#.#.###.#.#.#####.#####.#.###.#.#.#.#.#.#.#.#########.#######.###.#.#.###.#.#",
		"#.#.#.#.#...#.#.....#...#...#.#.#...>.>.#.#.#.#.#.#.......#...#.#.#.#...#.#.#...#...#...#...#.#.#.#.#.#.#.#.......#...#...#...#...#.#...#.#.#",
		"#.#.#.#.#.###.#.#####.#.#.#.#.#.#.#######.#.#.#.#.#######.#.###.#.#.#.###.#.#.#.###.#.#####.#.#.#.#.#.#.#.#######.#.###.#.#.###.###.###.#.#.#",
		"#...#...#.....#.......#...#...#...#######...#...#.........#.....#...#.....#...#.....#.......#...#...#...#.........#.....#...###.....###...#.#",
		"###########################################################################################################################################.#",
	}
)

const (
	MOVE_UP    = '^'
	MOVE_DOWN  = 'v'
	MOVE_RIGHT = '>'
	MOVE_LEFT  = '<'
)

type position struct {
	x, y int
}

func bytes(input []string) [][]byte {
	res := make([][]byte, len(input))
	for i := range input {
		res[i] = []byte(input[i])
	}
	return res
}

func display(forestMap [][]byte, path map[position]int) {
	yMax, xMax := len(forestMap), len(forestMap[0])
	str := strings.Builder{}
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			if v := path[position{x, y}]; v != 0 {
				if v == len(path) {
					fmt.Fprint(&str, "o")
				} else {
					fmt.Fprint(&str, "x")
				}
			} else {
				fmt.Fprintf(&str, "%c", forestMap[y][x])
			}
		}
		fmt.Fprint(&str, "\n")
	}
	fmt.Fprint(&str, "\n")
	fmt.Print(str.String())
}

func displayState(forestMap [][]byte, moves []position) {
	yMax, xMax := len(forestMap), len(forestMap[0])
	str := strings.Builder{}
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			found := false
			for _, move := range moves {
				if move.x == x && move.y == y {
					fmt.Fprint(&str, "o")
					found = true
					break
				}
			}
			if !found {
				fmt.Fprintf(&str, "%c", forestMap[y][x])
			}
		}
		fmt.Fprint(&str, "\n")
	}
	fmt.Fprint(&str, "\n")
	fmt.Print(str.String())
}

func isValidPosition(point byte, option byte, forbidSlope bool) bool {
	if forbidSlope {
		return point == '.' || point == option
	} else {
		return point != '#'
	}
}

func nextPositions(pos position, forestMap [][]byte, curPath map[position]int, forbidSlope bool) []position {
	nexts := []position{}
	if forestMap[pos.y][pos.x] == '#' {
		return nil
	}
	if isValidPosition(forestMap[pos.y][pos.x+1], '>', forbidSlope) {
		next := position{pos.x + 1, pos.y}
		if curPath[next] == 0 {
			nexts = append(nexts, next)
		}
	}
	if isValidPosition(forestMap[pos.y][pos.x-1], '<', forbidSlope) {
		next := position{pos.x - 1, pos.y}
		if curPath[next] == 0 {
			nexts = append(nexts, next)
		}
	}
	if pos.y > 0 && isValidPosition(forestMap[pos.y-1][pos.x], '^', forbidSlope) {
		next := position{pos.x, pos.y - 1}
		if curPath[next] == 0 {
			nexts = append(nexts, next)
		}
	}
	if pos.y < len(forestMap)-1 && isValidPosition(forestMap[pos.y+1][pos.x], 'v', forbidSlope) {
		next := position{pos.x, pos.y + 1}
		if curPath[next] == 0 {
			nexts = append(nexts, next)
		}
	}
	return nexts
}

func copyPath(path map[position]int) map[position]int {
	cpy := make(map[position]int)
	for k, v := range path {
		cpy[k] = v
	}
	return cpy
}

func computePaths(start position, path map[position]int, forest [][]byte, forbidSlope bool) int {
	end := position{len(forest) - 2, len(forest[0]) - 1}
	cur := start
	path[start] = len(path) + 1
	for cur != end {
		//display(forest, path)
		nexts := nextPositions(cur, forest, path, forbidSlope)
		if len(nexts) == 0 {
			display(forest, path)
			if cur != end {
				return -1
			}
			return len(path) - 1
		}
		if len(nexts) == 1 {
			cur = nexts[0]
			path[nexts[0]] = len(path) + 1
		} else {
			longuest := 0
			var longuestPath map[position]int
			for _, next := range nexts {
				newPath := copyPath(path)
				nextLonguest := computePaths(next, newPath, forest, forbidSlope)
				if nextLonguest > longuest {
					longuest = nextLonguest
					longuestPath = newPath
				}
			}
			display(forest, longuestPath)
			path = longuestPath
			return longuest
		}
	}
	return len(path) - 1
}

func TestDay23Phase1(t *testing.T) {
	start := position{1, 0}
	forest := bytes(lightInput)
	res := computePaths(start, make(map[position]int), forest, true)
	if res != 94 {
		t.Fail()
	}

	start = position{1, 0}
	forest = bytes(largeInput)
	res = computePaths(start, make(map[position]int), forest, true)
	if res != 1998 {
		t.Fail()
	}
}

type pathItem struct {
	start   *node
	visited map[position]bool
	move
	currentValue int
}

type move struct {
	target position
	cost   int
}

type edge struct {
	target position
	cost   int
}
type node struct {
	pos     position
	targets []edge
}

func pathToNextSplit(forest [][]byte, start position) edge {
	cur := start
	end := position{len(forest) - 2, len(forest[0]) - 1}
	path := map[position]int{start: 1}
	for {
		nexts := nextPositions(cur, forest, path, false)
		if len(nexts) != 1 || cur == end {
			return edge{cur, len(path)}
		}
		if len(nexts) == 1 {
			cur = nexts[0]
			path[nexts[0]] = len(path) + 1
		}
	}
}

func buildNodes(forest [][]byte) []*node {
	nodes := []*node{{pos: position{1, 0}}}
	processed := map[position]bool{}
	toProcess := map[position]bool{}
	end := position{len(forest) - 2, len(forest[0]) - 1}

	for i := 0; i < len(nodes); i++ {
		n := nodes[i]
		if processed[n.pos] || n.pos == end {
			continue
		}
		processed[n.pos] = true
		targets := nextPositions(n.pos, forest, map[position]int{n.pos: 1}, false)
		forest[n.pos.y][n.pos.x] = '#'
		for _, t := range targets {
			e := pathToNextSplit(forest, t)
			n.targets = append(n.targets, e)
		}
		forest[n.pos.y][n.pos.x] = '.'
		for _, t := range n.targets {
			if processed[t.target] || toProcess[t.target] {
				continue
			}
			toProcess[t.target] = true
			nodes = append(nodes, &node{pos: t.target})
		}
	}
	return nodes
}

func pathLength(graph map[position]*node, start, end position, visits map[position]bool) int {
	startNode := graph[start]
	currentPath := 0
	for _, t := range startNode.targets {
		if visits[t.target] {
			continue
		}
		visits[t.target] = true
		newPath := pathLength(graph, t.target, end, visits)
		if newPath+t.cost > currentPath {
			currentPath = newPath + t.cost
		}
		visits[t.target] = false
	}
	return currentPath
}

func longestPath(nds []*node, start, end position) int {
	graph := make(map[position]*node)
	for _, n := range nds {
		graph[n.pos] = n
	}

	visitedNodes := map[position]bool{}
	visitedNodes[start] = true
	return pathLength(graph, start, end, visitedNodes)
}

func TestDay23Phase2(t *testing.T) {
	start := position{1, 0}
	forest := bytes(lightInput)
	res := computePaths(start, make(map[position]int), forest, false)
	if res != 154 {
		t.Fail()
	}

	nodes := buildNodes(forest)
	for _, n := range nodes {
		fmt.Printf("%v\n", *n)
	}
	end := position{len(forest) - 2, len(forest[0]) - 1}
	res = longestPath(nodes, start, end)
	if res != 154 {
		t.Fail()
	}
	forest = bytes(largeInput)
	nodes = buildNodes(forest)
	end = position{len(forest) - 2, len(forest[0]) - 1}
	res = longestPath(nodes, start, end)
	if res != 6434 {
		t.Fail()
	}
}