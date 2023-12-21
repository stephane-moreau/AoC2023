package day19

import (
	"fmt"
	"strings"
	"testing"
)

var (
	lightInput = []string{
		"broadcaster -> a, b, c",
		"%a -> b",
		"%b -> c",
		"%c -> inv",
		"&inv -> a",
	}

	light2Input = []string{
		"broadcaster -> a",
		"%a -> inv, con",
		"&inv -> b",
		"%b -> con",
		"&con -> output",
	}

	largeInput = []string{
		`%dj -> fj, jn`, // rx
		`%xz -> cm`,     // &gq
		`%fn -> rj`,     // &km										&xj								&qs										&kz
		`%fv -> nt, zp`, // &jn										&mf								&ph										&zp
		`%ls -> ph, cf`, // gh fk tl gk fj dj br hf cr xt sl		vx	bf	rt	bx	sk	rj	jp		nh	cf	tk	ls	bj	dp	nz	cm	mg		sx	fv	nt	dv	mt	xg	tt	rk
		`%rk -> zp, tp`, //
		`&jn -> km, cr, vz`,
		`%nh -> ph, ls`,
		`%tx -> gb`,
		`%xg -> dv, zp`,
		`%tp -> gx`,
		`&zp -> kj, kz, gx, fv, lv, tp`,
		`&gq -> rx`,
		`%fj -> sl, jn`,
		`%cr -> vz, jn`,
		`%rt -> fn, mf`,
		`%kj -> tt`,
		`%tk -> mg, ph`,
		`%xt -> jn, gh`,
		`%qx -> bx`,
		`%lv -> sx`,
		`%nz -> dp, ph`,
		`%sx -> kj, zp`,
		`%dd -> bf`,
		`%gb -> jp`,
		`%bj -> ph, nn`,
		`%sk -> mf`,
		`%bx -> tx, mf`,
		`%mt -> xg, zp`,
		`%vz -> hf`,
		`%vx -> mf, sk`,
		`%tt -> mt, zp`,
		`%br -> jn, fk`,
		`&xj -> gq`,
		`%mg -> ph, ps`,
		`%nt -> zp, rk`,
		`&qs -> gq`,
		`%rj -> qx, mf`,
		`%bf -> vx, mf`,
		`&kz -> gq`,
		`%fk -> jn, gk`,
		`%dv -> zp`,
		`%dp -> ph`,
		`&mf -> gb, tx, xj, dd, qx, rt, fn`,
		`&ph -> nn, xz, tk, ps, qs`,
		`%ps -> xz`,
		`&km -> gq`,
		`broadcaster -> fv, cr, rt, tk`,
		`%gk -> jn, xt`,
		`%cf -> ph, nz`,
		`%tl -> jn, br`,
		`%cm -> bj, ph`,
		`%nn -> nh`,
		`%jp -> mf, dd`,
		`%gh -> jn, dj`,
		`%hf -> tl, jn`,
		`%sl -> jn`,
		`%gx -> lv`,
	}
)

type module struct {
	Name    string
	Type    string
	Targets []string
	// State of the module
	onOff    bool // Fliflop
	junction map[string]bool
}

const (
	FLIP_FLOP   = "%"
	CONJUNCTION = "&"
	BROADCAST   = "b"
)

type pulse struct {
	from, to string
	lowHigh  bool
}

func (m *module) isDefault() bool {
	switch m.Type {
	case FLIP_FLOP:
		return !m.onOff
	case CONJUNCTION:
		for _, state := range m.junction {
			if state {
				return false
			}
		}
		return true
	}
	return true
}

func (m *module) processInput(input string, state bool) []pulse {
	lowHigh := false
	switch m.Type {
	case FLIP_FLOP:
		// Generates:
		//    true -> nothing
		//    false -> true + false
		if state {
			return nil
		}
		m.onOff = !m.onOff
		lowHigh = m.onOff

	case CONJUNCTION:
		// Generates:
		// 2n-1 -> false
		// 1 -> true
		m.junction[input] = state
		allHigh := true
		for _, pulse := range m.junction {
			allHigh = allHigh && pulse
		}
		lowHigh = !allHigh
	default:
		// no op
	}

	nexts := make([]pulse, len(m.Targets))
	for i, t := range m.Targets {
		nexts[i] = pulse{m.Name, t, lowHigh}
	}
	return nexts
}

func parseInput(input []string) map[string]*module {
	mods := make(map[string]*module)
	cnj := make([]*module, 0, 10)
	for _, line := range input {
		nameAndTarget := strings.Split(line, "->")
		name := strings.TrimSpace(nameAndTarget[0][1:])
		typ := nameAndTarget[0][:1]
		if typ != FLIP_FLOP && typ != CONJUNCTION {
			name = strings.TrimSpace(nameAndTarget[0])
		}
		targets := strings.Split(nameAndTarget[1], ",")
		for i := range targets {
			targets[i] = strings.TrimSpace(targets[i])
		}
		mod := module{
			Name:     name,
			Type:     nameAndTarget[0][:1],
			Targets:  targets,
			junction: make(map[string]bool),
		}
		mods[mod.Name] = &mod
		if mod.Type == CONJUNCTION {
			cnj = append(cnj, &mod)
		}
	}
	// Initializes input of conjunctions
	for _, c := range mods {
		for _, v := range mods {
			for _, t := range v.Targets {
				if t == c.Name {
					c.junction[v.Name] = false
				}
			}
		}
	}
	return mods
}

type loopResult struct {
	low  int
	high int
}

func log(format string, params ...any) {
	if false {
		fmt.Printf(format, params...)
	}
}

func processCycle(mods map[string]*module, res *loopResult, shouldStop func(pulse) bool) bool {
	pulses := make([]pulse, 1, 1000)
	pulses[0] = pulse{
		"button",
		"broadcaster",
		false,
	}
	for i := 0; i < len(pulses); i++ {
		pulse := pulses[i]
		if pulse.from != "button" {
			log("%s -> %v to %s\n", pulse.from, pulse.lowHigh, pulse.to)
		}
		if pulse.lowHigh && pulse.to == "gq" {
			print("break here")
		}
		if shouldStop != nil && shouldStop(pulse) {
			return true
		}
		if pulse.lowHigh {
			res.high++
		} else {
			res.low++
		}
		module := mods[pulse.to]
		if module == nil {
			continue
		}
		nexts := module.processInput(pulse.from, pulse.lowHigh)
		pulses = append(pulses, nexts...)
	}
	return false
}

func computeSwitch(input []string, rounds int, shouldStop func(pulse) bool) int {
	mods := parseInput(input)
	var lowCount, highCount int

	loops := make(map[int]loopResult)
ButtonClick:
	for r := 0; r < rounds || rounds == -1; r++ {
		var res loopResult
		log("-------START %d---------\n", r+1)
		if processCycle(mods, &res, shouldStop) {
			return r
		}
		log("-------END %d---------\n", r+1)
		if shouldStop == nil {

			loops[r] = res
			// check is cycles done
			for _, mod := range mods {
				if !mod.isDefault() {
					continue ButtonClick
				}
			}
			break
		}
	}
	fullCyles := rounds / len(loops)
	for _, res := range loops {
		lowCount += res.low
		highCount += res.high
	}
	lowCount *= fullCyles
	highCount *= fullCyles
	for i := 0; i < rounds-fullCyles*len(loops); i++ {
		lowCount += loops[i].low
		highCount += loops[i].high
	}
	if len(loops) != rounds {
		fmt.Printf("Cycle length : %d - %v", len(loops), loops)
	}
	fmt.Printf("low: %d, high %d - Total : %d\n", lowCount, highCount, lowCount*highCount)
	return lowCount * highCount
}

func TestDay20Phase1(t *testing.T) {

	swc := computeSwitch(lightInput, 10, nil)
	if swc != 3200 {
		t.Fail()
	}
	swc = computeSwitch(light2Input, 10, nil)
	if swc != 1176 {
		t.Fail()
	}
	swc = computeSwitch(lightInput, 1000, nil)
	if swc != 32_000_000 {
		t.Fail()
	}

	swc = computeSwitch(light2Input, 1000, nil)
	if swc != 11_687_500 {
		t.Fail()
	}
	swc = computeSwitch(largeInput, 1000, nil)
	if swc != 832957356 {
		t.Fail()
	}
}

type graph map[string]*module

func cyclesToReset(g graph, loopStatus func(pulse, int) bool) int {
	loops := make(map[int]loopResult)
ButtonClick:
	for r := 0; r < 10_000; r++ {
		var res loopResult
		log("-------START %d---------\n", r+1)
		if processCycle(g, &res, func(p pulse) bool {
			return loopStatus != nil && loopStatus(p, r)
		}) {
			return r
		}
		log("-------END %d---------\n", r+1)
		loops[r] = res
		// check is cycles done
		for _, mod := range g {
			if !mod.isDefault() {
				continue ButtonClick
			}
		}
		break
	}
	return len(loops)
}

func cycleCount(input []string) int {
	mods := parseInput(input)
	// finds the end module
	var last string
	var res loopResult
	processCycle(mods, &res, func(p pulse) bool {
		last = p.to
		return false
	})
	fmt.Printf("Last activity is '%s'\n", last)
	sources := make([]module, 0)
	for _, v := range mods {
		for _, t := range v.Targets {
			if t == last {
				sources = append(sources, *v)
			}
		}
	}
	fmt.Printf("Sources are '%v'\n", sources)

	// Complete reset
	mods = parseInput(input)
	// split graph into sub parts
	cycles := make(map[string]int)
	for _, s := range sources {
		for k := range s.junction {
			cycles[k] = 0
		}
	}

	cyclesToReset(mods, func(p pulse, r int) bool {
		completed := true
		for k, v := range cycles {
			if v != 0 {
				continue
			}
			completed = false
			if p.from == k && p.lowHigh {
				cycles[k] = r + 1
			}
		}
		return completed
	})
	values := make([]int, 0, len(cycles))
	for _, v := range cycles {
		values = append(values, v)
	}

	return ppcm(values)
}

func ppcm(values []int) int {
	primes := make([]map[int]int, len(values))

	primeNumbers := []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
		73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
		179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281,
	}
	for i, period := range values {
		decompositions := make(map[int]int)
	decomposition:
		for _, p := range primeNumbers {
			for period%p == 0 {
				decompositions[p] = decompositions[p] + 1
				period = period / p
				if period < p*p {
					decompositions[period] = decompositions[period] + 1
					break decomposition
				}
			}
		}
		if len(decompositions) == 0 {
			decompositions[period] = 1
		}
		primes[i] = decompositions
	}
	count := 1
	usedPrimes := map[int]int{}
	for _, decomp := range primes {
		for p, i := range decomp {
			if usedPrimes[p] < i {
				usedPrimes[p] = i
			}
		}
	}
	for p, usedCount := range usedPrimes {
		for i := 0; i < usedCount; i++ {
			count = count * p
		}
	}
	return count
}

func TestDay20Phase2(t *testing.T) {

	lightMods := parseInput(lightInput)
	c := cyclesToReset(lightMods, nil)
	if c != 1 {
		t.Fail()
	}
	light2Mods := parseInput(light2Input)
	c = cyclesToReset(light2Mods, nil)
	if c != 4 {
		t.Fail()
	}
	c = cycleCount(largeInput)
	if c != 240162699605221 {
		t.Fail()
	}
}
