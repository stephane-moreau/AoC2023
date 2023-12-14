package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

const (
	SEED_HEADER = "seeds:"
	MAP_HEADER  = "map:"
	mapSize     = 100
)

type mapping struct {
	source int64
	target int64
	size   int64
}

type mappings []mapping

func (m mappings) MappedValue(v int64) int64 {
	for _, mapping := range m {
		if v >= mapping.source && v < mapping.source+mapping.size {
			return mapping.target + v - mapping.source
		}
	}
	return v
}

type seedRange struct {
	start int64
	size  int64
}

type seedRanges []seedRange

func loadData(fileName string) (seedRanges, []mappings, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, nil, err
	}
	seedNums := make(seedRanges, 0, 20)
	maps := make([]mappings, 0, 20)
	var newMap mappings
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, SEED_HEADER) {
			seedValues := strings.Split(line[len(SEED_HEADER)+1:], " ")
			for i := 0; i < len(seedValues); i += 2 {
				seedStart, err := strconv.Atoi(seedValues[i])
				if err != nil {
					return nil, nil, err
				}
				seedSize, err := strconv.Atoi(seedValues[i+1])
				if err != nil {
					return nil, nil, err
				}
				seedNums = append(seedNums, seedRange{
					start: int64(seedStart),
					size:  int64(seedSize),
				})
			}
			continue
		}
		if strings.HasSuffix(line, MAP_HEADER) {
			if newMap != nil {
				maps = append(maps, newMap)
			}
			newMap = make(mappings, 0, 100)
			continue
		}
		parts := strings.Split(line, " ")
		target, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}
		source, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}
		size, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, nil, err
		}
		newMap = append(newMap, mapping{
			source: int64(source),
			target: int64(target),
			size:   int64(size),
		})
	}
	if newMap != nil {
		maps = append(maps, newMap)
	}

	return seedNums, maps, nil
}

func compute(fileName string) {
	seeds, maps, err := loadData(fileName)
	if err != nil {
		panic(err)
	}
	result := int64(-1)
	for _, seed := range seeds {
		for i := int64(0); i < seed.size; i++ {
			v := seed.start + i
			for _, mapping := range maps {
				v = mapping.MappedValue(v)
			}
			if result == -1 || result > v {
				result = v
			}
		}
	}
	fmt.Printf("Nearest localtion is %d\n", result)
}

func TestDay5Phase1(t *testing.T) {
	compute("lightdata.txt")
	compute("largedata.txt")
}

func TestDay5Phase2(t *testing.T) {
	compute("lightdata.txt")
	compute("largedata.txt")
}
