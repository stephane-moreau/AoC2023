package day24

import (
	"fmt"
	"math"
	"math/big"
	"sort"
	"testing"
)

func Abs(f *big.Float) *big.Float {
	v := *f
	return v.Abs(f)
}

func Div(x, y *big.Float) *big.Float {
	v := big.NewFloat(0)
	return v.Quo(x, y)
}

func Sub(x, y *big.Float) *big.Float {
	v := big.NewFloat(0)
	return v.Sub(x, y)
}

func Mul(x, y *big.Float) *big.Float {
	v := big.NewFloat(0)
	return v.Mul(x, y)
}

func Neg(x *big.Float) *big.Float {
	v := big.NewFloat(0)
	return v.Neg(x)
}

func Int64(x *big.Float) float64 {
	f64, _ := x.Float64()
	return math.Round(f64)
}

func solveBF(matrix [][]*big.Float) {
	// Triangularization
	fmt.Printf("%v\n", matrix)
	for col := 0; col < len(matrix[0])-1; col++ {
		sort.SliceStable(matrix[col:], func(i, j int) bool {
			return Abs(matrix[i+col][col]).Cmp(Abs(matrix[j+col][col])) == 1
		})
		for l := col + 1; l < len(matrix); l++ {
			if matrix[l][col].Cmp(big.NewFloat(0)) == 0 {
				continue
			}
			fmt.Printf("%v\n", matrix)
			f := Div(matrix[l][col], matrix[col][col])
			for c := col; c < len(matrix[0]); c++ {
				matrix[l][c] = Sub(matrix[l][c], Mul(f, matrix[col][c]))
			}
		}
	}
	fmt.Printf("%v\n", matrix)
	// diagonalization / normalization
	for l := len(matrix) - 1; l >= 0; l-- {
		n := matrix[l][l]
		for c := l; c < len(matrix[0]); c++ {
			matrix[l][c] = Div(matrix[l][c], n)
		}

		fmt.Printf("%v\n", matrix)
		for c := 0; c < l; c++ {
			n := matrix[c][l]
			for col := l; col < len(matrix[0]); col++ {
				matrix[c][col] = Sub(matrix[c][col], Mul(n, matrix[l][col]))
			}
		}
	}
	fmt.Printf("###########\n%v\n###########\n", matrix)
}

// func TestGaussianElimination(t *testing.T) {
// 	matrix := [][]float64{
// 		{2, 2, -1, 1, 4},
// 		{4, 3, -1, 2, 6},
// 		{8, 5, -3, 4, 12},
// 		{3, 3, -2, 2, 6},
// 	}
// 	solve(matrix)

// 	matrix = [][]float64{
// 		{2, 3, 11, 5, 2},
// 		{1, 1, 5, 2, 1},
// 		{2, 1, -3, 2, -3},
// 		{1, 1, -3, 4, -3},
// 	}
// 	solve(matrix)

// 	matrix = [][]float64{
// 		{2, 1, -1, 8},
// 		{-3, -1, 2, -11},
// 		{-2, 1, 2, -3},
// 	}
// 	solve(matrix)
// }

// pR + t*VR = p0 + t*V0
// pR + t1*VR = p1 + t1*V1
// pR + t2*VR = p2 + t2*V2
// ->
// (pR-p0) = -t (vR-V0)
// -> (scalar ratio between vectors)
// (pR-p0)x(vR-v0) = 0
// pRxvR - p0xvR - pRxv0 + p0xv0 = 0
// pRxvR - p1xvR - pRxv1 + p1xv1 = 0
// pRxvR - p2xvR - pRxv2 + p2xv2 = 0
//
// p0xvR + pRxv0 - p0xv0 = p1xvR + pRxv1 - p1xV1
// (p0-p1)xvR + pRx(v0-v1) = p0xv0-p1xv1
// (p0-p2)xvR + pRx(v0-v2) = p0xv0-p2xv2
// Expansion:
//	P0_P1.y*vR.z - P0_P1.z*vR.y + pR.y*V0_V1.z - pR.z*V0_v1.y = PV0-PV1
//	P0_P1.z*vR.x - P0_P1.x*vR.z + pR.z*V0_V1.x - pR.x*V0_v1.z = PV0-PV1
//	P0_P1.x*vR.y - P0_P1.y*vR.x + pR.x*V0_V1.y - pR.y*V0_v1.x = PV0-PV1

func findStoneParamsBF(input []string) hailStone {
	stones := parseInput(input)

	stone0 := stones[0]
	stone1 := stones[1]
	stone2 := stones[2]

	stonePV0 := crossProduct(stone0.start, stone0.direction)
	stonePV1 := crossProduct(stone1.start, stone1.direction)
	stonePV2 := crossProduct(stone2.start, stone2.direction)

	stoneP0_P1 := diff(stone0.start, stone1.start)
	stoneP0_P2 := diff(stone0.start, stone2.start)

	stoneV0_V1 := diff(stone0.direction, stone1.direction)
	stoneV0_V2 := diff(stone0.direction, stone2.direction)

	stonePV0_PV1 := diff(stonePV0, stonePV1)
	stonePV0_PV2 := diff(stonePV0, stonePV2)

	stoneR := hailStone{start: position{24, 13, 10}, direction: position{-3, 1, 2}}

	vr1 := crossProduct(stoneP0_P1, stoneR.direction)
	pr1 := crossProduct(stoneR.start, stoneV0_V1)
	fmt.Printf("%v %v %v\n", vr1, pr1, stonePV0_PV1)

	vr2 := crossProduct(stoneP0_P2, stoneR.direction)
	pr2 := crossProduct(stoneR.start, stoneV0_V2)
	fmt.Printf("%v %v %v\n", vr2, pr2, stonePV0_PV2)
	matrix := [6][]*big.Float{}

	// 0 & 1  : VR.x VR.y VR.z PR.x PR.y PR.z Res
	matrix[0] = []*big.Float{
		big.NewFloat(0),
		big.NewFloat(-stoneP0_P1.z),
		big.NewFloat(stoneP0_P1.y),
		big.NewFloat(0),
		big.NewFloat(-stoneV0_V1.z),
		big.NewFloat(stoneV0_V1.y),
		big.NewFloat(stonePV0_PV1.x),
	}
	matrix[1] = []*big.Float{
		big.NewFloat(stoneP0_P1.z),
		big.NewFloat(0),
		big.NewFloat(-stoneP0_P1.x),
		big.NewFloat(stoneV0_V1.z),
		big.NewFloat(0),
		big.NewFloat(-stoneV0_V1.x),
		big.NewFloat(stonePV0_PV1.y),
	}
	matrix[2] = []*big.Float{
		big.NewFloat(-stoneP0_P1.y),
		big.NewFloat(stoneP0_P1.x),
		big.NewFloat(0),
		big.NewFloat(-stoneV0_V1.y),
		big.NewFloat(stoneV0_V1.x),
		big.NewFloat(0),
		big.NewFloat(stonePV0_PV1.z),
	}
	// 0 & 2
	matrix[3] = []*big.Float{
		big.NewFloat(0),
		big.NewFloat(-stoneP0_P2.z),
		big.NewFloat(stoneP0_P2.y),
		big.NewFloat(0),
		big.NewFloat(-stoneV0_V2.z),
		big.NewFloat(stoneV0_V2.y),
		big.NewFloat(stonePV0_PV2.x),
	}
	matrix[4] = []*big.Float{
		big.NewFloat(stoneP0_P2.z),
		big.NewFloat(0),
		big.NewFloat(-stoneP0_P2.x),
		big.NewFloat(stoneV0_V2.z),
		big.NewFloat(0),
		big.NewFloat(-stoneV0_V2.x),
		big.NewFloat(stonePV0_PV2.y),
	}
	matrix[5] = []*big.Float{
		big.NewFloat(-stoneP0_P2.y),
		big.NewFloat(stoneP0_P2.x),
		big.NewFloat(0),
		big.NewFloat(-stoneV0_V2.y),
		big.NewFloat(stoneV0_V2.x),
		big.NewFloat(0),
		big.NewFloat(stonePV0_PV2.z),
	}

	fmt.Printf("%v\n", matrix)
	solveBF(matrix[:])
	return hailStone{
		direction: position{
			x: Int64(matrix[0][6]),
			y: Int64(matrix[1][6]),
			z: Int64(matrix[2][6]),
		},
		start: position{
			x: Int64(Neg(matrix[3][6])),
			y: Int64(Neg(matrix[4][6])),
			z: Int64(Neg(matrix[5][6])),
		},
	}
}

func TestDay24Phase2BF(t *testing.T) {
	hs := findStoneParams(lightInput)
	if hs.start.x+hs.start.y+hs.start.z != 47 {
		t.Fail()
	}

	hs = findStoneParams(largeInput)
	if hs.start.x+hs.start.y+hs.start.z != 888708704663413 {
		t.Fail()
	}
	// c = findStoneParams(lightInput)
	// if c != 888708704663413 {
	// 	t.Fail()
	// }
}
