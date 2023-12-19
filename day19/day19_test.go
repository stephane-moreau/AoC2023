package day19

import (
	"testing"
)

func TestDay19Phase1(t *testing.T) {
	// 19114
	computePiecesIndex(lightInput)
	// 432788
	computePiecesIndex(largeInput)
}

func TestDay19Phase2(t *testing.T) {
	// 167409079868000
	computePiecesRange(lightInput)
	// 142863718918201
	computePiecesRange(largeInput)
}
