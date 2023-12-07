package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func readFile(filename string) string {
	data := make([]byte, 1024)
	output := ""
	n := 0
	var err error
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for n != 0 && err != io.EOF {
		n, err = file.Read(data)
		if n != 0 {
			output += string(data[:n])
		}
	}
	return output

}

func readLines(filename string, cb func(string) bool) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() && cb(scanner.Text()) {
	}

	return scanner.Err()
}

func isnum(c byte) bool {
	return c >= '0' && c <= '9'
}

func f_1_1() {
	sum := 0
	_ = readLines("1.input", func(line string) bool {
		var first, last int
		for i := 0; i < len(line); i++ {
			if isnum(line[i]) {
				first = int(line[i] - '0')
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if isnum(line[i]) {
				last = int(line[i] - '0')
				break
			}
		}
		println(line, first, last)
		sum += first*10 + last
		return true
	})
	println(sum)

}

func findNumber(line string) int {

	words := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	if isnum(line[0]) {
		return int(line[0] - '0')
	} else {
		for k, v := range words {
			if strings.HasPrefix(line, k) {
				return v
			}
		}
	}
	return -1
}

func f_1_2() {
	sum := 0

	_ = readLines("1.input", func(line string) bool {
		var first, last int
		for i := 0; i < len(line); i++ {
			first = findNumber(line[i:])
			if first != -1 {
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			last = findNumber(line[i:])
			if last != -1 {
				break
			}
		}
		println(line, first, last)
		sum += first*10 + last
		return true
	})
	println(sum)

}

func parse_2_1(line string) (int, []map[string]int) {
	gameNumber := 0
	gamesBalls := []map[string]int{}
	line, _ = strings.CutPrefix(line, "Game ")
	gameNumber, _ = strconv.Atoi(line[:strings.Index(line, ":")])
	line = line[strings.Index(line, ":")+1:]
	games := strings.Split(line, ";")
	for _, game := range games {
		balls := strings.Split(game, ",")
		gameBalls := map[string]int{}
		for _, ball := range balls {
			ball = strings.Trim(ball, " ")
			color := ball[strings.Index(ball, " "):]
			color = strings.Trim(color, " ")
			number, _ := strconv.Atoi(ball[:strings.Index(ball, " ")])
			gameBalls[color] = number
		}

		gamesBalls = append(gamesBalls, gameBalls)
	}
	return gameNumber, gamesBalls
}

func f_2_1() {

	sum := 0
	whatIHave := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	_ = readLines("2.input", func(line string) bool {
		gameNumber, gamesBalls := parse_2_1(line)
		ok := true
		for color, max := range whatIHave {
			if ok {
				for _, gameBalls := range gamesBalls {
					if gameBalls[color] > max {
						ok = false
						break
					}
				}
			}
		}
		println(line, ok)
		if ok {
			sum += gameNumber
		}
		return true
	})

	println(sum)

}

func f_2_2() {
	sum := 0
	_ = readLines("2.input", func(line string) bool {
		_, gamesBalls := parse_2_1(line)
		maxColors := map[string]int{}
		for _, gameBalls := range gamesBalls {
			for color, number := range gameBalls {
				if maxColors[color] < number {
					maxColors[color] = number
				}
			}
		}
		power := 1
		for _, number := range maxColors {
			power *= number
		}
		sum += power
		fmt.Printf("%s %v %d\n", line, maxColors, power)
		return true
	})

	println(sum)
}

func f_3_1() {
	sum := 0
	array := []string{}
	_ = readLines("3.input", func(line string) bool {
		array = append(array, line)
		return true
	})

	width := len(array[0])
	height := len(array)
	checkArray := []struct {
		x int
		y int
	}{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0} /*{0, 0},*/, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	isSymbol := func(c byte) bool {
		return c != '.' && !isnum(c)
	}

	checkSurroundings := func(x int, y int) (bool, bool) {
		println(x, y, "checking surroundings", string(array[y][x]))
		foundSymbol := false
		end := false
		for _, c := range checkArray {
			if x+c.x < 0 || x+c.x >= width || y+c.y < 0 || y+c.y >= height {
				println(x, y, "out of bounds", x+c.x, y+c.y)
				//out of bounds, continue rest of check
				continue
			}
			if isSymbol(array[y+c.y][x+c.x]) {
				println(x, y, "has symbol", x+c.x, y+c.y, string(array[y+c.y][x+c.x]))
				foundSymbol = true
				break
			}
		}
		//do the special check for the next character. if the next character is out of bounds or a dot, it's the end of the sequence
		//if it's a symbol, it's not ok:
		//isIsland true,  end true: .123.
		//isIsland true,  end true: .123
		//isIsland false, end false: .123x
		if x+1 >= width || !isnum(array[y][x+1]) {
			println(x, y, "end because eol or not number next")
			end = true
		}

		println(x, y, "conclusion found symbol", foundSymbol, "is end", end)
		return foundSymbol, end
	}

	accumulated := ""

	for y := 0; y < len(array); y++ {
		hasSymbol := false
		for x := 0; x < len(array[y]); x++ {
			if isnum(array[y][x]) {
				accumulated += string(array[y][x])
				foundSymbol, isEnd := checkSurroundings(x, y)
				hasSymbol = hasSymbol || foundSymbol
				if !isEnd {
					continue
				}
				if hasSymbol {
					println("accumulated", accumulated)
					num, _ := strconv.Atoi(accumulated)
					sum += num
				} else {
					println("end without symbol, resetting", accumulated)
				}
				accumulated = ""
				hasSymbol = false
			}
		}
	}
	println(sum)
}
func f_3_2() {
	array := []string{}
	_ = readLines("3.input", func(line string) bool {
		array = append(array, line)
		return true
	})

	width := len(array[0])
	height := len(array)
	checkArray := []struct {
		x int
		y int
	}{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0} /*{0, 0},*/, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	isSymbol := func(c byte) bool {
		return c != '.' && !isnum(c)
	}

	checkSurroundings := func(x int, y int) (bool, bool, []string) {
		println(x, y, "checking surroundings", string(array[y][x]))
		foundSymbol := false
		end := false
		gears := []string{}
		for _, c := range checkArray {
			if x+c.x < 0 || x+c.x >= width || y+c.y < 0 || y+c.y >= height {
				println(x, y, "out of bounds", x+c.x, y+c.y)
				//out of bounds, continue rest of check
				continue
			}
			if isSymbol(array[y+c.y][x+c.x]) {
				println(x, y, "has symbol", x+c.x, y+c.y, string(array[y+c.y][x+c.x]))
				foundSymbol = true
				if array[y+c.y][x+c.x] == '*' {
					gears = append(gears, fmt.Sprintf("(%d,%d)", x+c.x, y+c.y))
				}
			}
		}
		//do the special check for the next character. if the next character is out of bounds or a dot, it's the end of the sequence
		//if it's a symbol, it's not ok:
		//isIsland true,  end true: .123.
		//isIsland true,  end true: .123
		//isIsland false, end false: .123x
		if x+1 >= width {
			println(x, y, "end because eol or not number next")
			end = true
		} else if !isnum(array[y][x+1]) {
			println(x, y, "end because eol or not number next")
			if array[y][x+1] == '*' {
				gears = append(gears, fmt.Sprintf("(%d,%d)", x+1, y))
			}
			end = true
		}

		println(x, y, "conclusion found symbol", foundSymbol, "is end", end, "gears", fmt.Sprintf("%v", gears))
		return foundSymbol, end, gears
	}

	accumulated := ""
	allGears := map[string]bool{}
	gearMap := map[string][]int{}

	for y := 0; y < len(array); y++ {
		hasSymbol := false
		for x := 0; x < len(array[y]); x++ {
			if isnum(array[y][x]) {
				accumulated += string(array[y][x])
				foundSymbol, isEnd, gears := checkSurroundings(x, y)
				for _, gear := range gears {
					allGears[gear] = true
				}
				hasSymbol = hasSymbol || foundSymbol
				if !isEnd {
					continue
				}
				if hasSymbol {
					println("accumulated", accumulated)
					num, _ := strconv.Atoi(accumulated)
					for gear, _ := range allGears {
						gearMap[gear] = append(gearMap[gear], num)
					}
				} else {
					println("end without symbol, resetting", accumulated)
				}
				accumulated = ""
				hasSymbol = false
				allGears = map[string]bool{}
			}
		}
	}

	sum := 0
	for k, v := range gearMap {
		if len(v) != 2 {
			println("gear", k, "is not a pair", len(v))
			continue
		}
		println("gear", k, "is a pair", len(v))
		sum += v[0] * v[1]
	}
	println(sum)
}

func parse_4_1(line string) (card int, l1 map[int]bool, l2 map[int]bool) {
	l1 = map[int]bool{}
	l2 = map[int]bool{}
	splits := strings.Split(line, ":")
	splits[0], _ = strings.CutPrefix(splits[0], "Card ")
	card, _ = strconv.Atoi(strings.Trim(splits[0], " "))
	splits = strings.Split(splits[1], "|")
	for _, num := range strings.Split(splits[0], " ") {
		num = strings.Trim(num, " ")
		if len(num) != 0 {
			n, _ := strconv.Atoi(num)
			l1[n] = true
		}
	}
	for _, num := range strings.Split(splits[1], " ") {
		num = strings.Trim(num, " ")
		if len(num) != 0 {
			n, _ := strconv.Atoi(num)
			l2[n] = true
		}
	}
	return card, l1, l2
}

func f_4_1() {
	sum := 0
	readLines("4.input", func(line string) bool {
		println(line)
		game, l1, l2 := parse_4_1(line)
		gamesum := 0
		for num := range l2 {
			if l1[num] {
				if gamesum == 0 {
					gamesum = 1
				} else {
					gamesum *= 2
				}
			}
		}
		fmt.Printf("card %d l1 %v l2 %v gamesum %d\n", game, l1, l2, gamesum)
		sum += gamesum
		return true
	})

	println(sum)
}

func f_4_2() {
	sum := 0
	gameCount := map[int]int{}
	readLines("4.input", func(line string) bool {
		game, l1, l2 := parse_4_1(line)
		//count the current game
		gameCount[game]++
		matches := 0
		for num := range l2 {
			if l1[num] {
				matches++
			}
		}
		m := gameCount[game]
		for o := 0; o < matches; o++ {
			i := o + 1
			if gameCount[game+i] == 0 {
				gameCount[game+i] = m
			} else {
				gameCount[game+i] += m
			}
			fmt.Printf("card %d adding %d to card %d, card %d now has %d cards\n", game, m, game+i, game+i, gameCount[game+i])
		}
		return true
	})
	for _, v := range gameCount {
		sum += v
	}

	println(sum)
}

type triplet struct {
	src int
	dst int
	rng int
}

type rng struct {
	src int
	rng int
}

func (r rng) String() string {
	return fmt.Sprintf("(%d,%d)", r.src, r.rng)
}

type triplets struct {
	title string
	vals  []triplet
}

func (t *triplet) parse(line string) {
	arr := str2arrnum(line)
	t.dst = arr[0]
	t.src = arr[1]
	t.rng = arr[2]

}

func (t *triplet) calcDest(seed int) (int, bool) {
	if seed <= t.src+t.rng && seed >= t.src {
		return seed + (t.dst - t.src), true
	}
	return seed, false

}

func (t *triplet) calcDestRanges(srngs []rng) (overlapping []rng, notOverlapping []rng) {
	for _, srng := range srngs {
		begin := srng.src
		end := srng.src + srng.rng //excluded

		obegin := t.src
		oend := t.src + t.rng //excluded
		destOffset := t.dst - t.src

		fmt.Printf("calcDestRanges seed (%d,%d), src (%d, %d), offset %d\n", begin, end, obegin, oend, destOffset)

		if end <= obegin || begin >= oend {
			ret := []rng{srng}
			fmt.Printf("no overlap %v\n", ret)
			notOverlapping = append(notOverlapping, ret...)
			continue
		}
		if begin >= obegin && end <= oend {
			ret := []rng{{begin + destOffset, end - begin}}
			fmt.Printf("full overlap inside %v\n", ret)
			overlapping = append(overlapping, ret...)
			continue
		}
		if begin < obegin && end > oend {
			no := []rng{
				{begin, obegin - begin},
				{oend, end - oend},
			}
			o := []rng{{obegin + destOffset, oend - obegin}}
			fmt.Printf("full overlap outside %v %v\n", o, no)
			overlapping = append(overlapping, o...)
			notOverlapping = append(notOverlapping, no...)
			continue
		}

		if begin < obegin {
			no := []rng{
				{begin, obegin - begin},
			}
			o := []rng{
				{obegin + destOffset, end - obegin},
			}
			fmt.Printf("overlap start %v %v\n", o, no)
			overlapping = append(overlapping, o...)
			notOverlapping = append(notOverlapping, no...)
			continue
		}

		if begin < oend {
			o := []rng{
				{begin + destOffset, oend - begin},
			}
			no := []rng{
				{oend, end - oend},
			}
			fmt.Printf("overlap end %v %v \n", o, no)
			overlapping = append(overlapping, o...)
			notOverlapping = append(notOverlapping, no...)
			continue
		}

		panic("unreachable")
	}
	return overlapping, notOverlapping
}

func (t *triplets) calcDest(seed int) int {
	for _, triplet := range t.vals {
		if dest, ok := triplet.calcDest(seed); ok {
			return dest
		}
	}
	return seed

}

func (t *triplets) calcRanges(srngs map[string]rng) map[string]rng {
	ret := map[string]rng{}
	var notoverlapping []rng
	for _, srng := range srngs {
		notoverlapping = append(notoverlapping, srng)
	}

	for _, srng := range notoverlapping {
		in := []rng{srng}
		for _, triplet := range t.vals {
			overlapping, out := triplet.calcDestRanges(in)
			for _, v := range overlapping {
				ret[v.String()] = v
			}
			in = out
		}

		//add the no overlaps at the end
		for _, v := range in {
			ret[v.String()] = v
		}
	}
	return ret

}

func str2arrnum(line string) (seeds []int) {
	for _, seed := range strings.Split(line, " ") {
		if seed == "" {
			continue
		}
		num, _ := strconv.Atoi(seed)
		seeds = append(seeds, num)
	}
	return seeds
}

func parse_5_1() (seeds []int, maps []triplets) {
	currentTriplets := triplets{}
	readLines("5.input", func(line string) bool {
		if line == "" {
			if len(currentTriplets.vals) != 0 {
				maps = append(maps, currentTriplets)
			}
			currentTriplets = triplets{}
			return true
		}

		if strings.HasPrefix(line, "seeds: ") {
			line = strings.TrimPrefix(line, "seeds: ")
			seeds = str2arrnum(line)
			return true
		}

		if !isnum(line[0]) {
			currentTriplets.title = line
			return true
		}
		t := triplet{}
		t.parse(line)
		currentTriplets.vals = append(currentTriplets.vals, t)
		return true

	})

	if len(currentTriplets.vals) != 0 {
		maps = append(maps, currentTriplets)
	}

	return seeds, maps
}

func minInt(arr []int) int {
	min := math.MaxInt
	for _, num := range arr {
		if num < min {
			min = num
		}
	}
	return min
}

func f_5_2() {
	seeds, triplets := parse_5_1()
	seedRanges := []rng{}
	for i := 0; i < len(seeds); i = i + 2 {
		seedRanges = append(seedRanges, rng{seeds[i], seeds[i+1]})
	}
	fmt.Printf("seeds %v triplets %v (len:%d)\n", seedRanges, triplets, len(triplets))
	seedDest := []int{}
	for _, seedRange := range seedRanges {
		fmt.Printf("seed-range-start %v\n", seedRange)
		seedRangesTmp := map[string]rng{seedRange.String(): seedRange}
		for _, triplet := range triplets {
			seedRangesTmp = triplet.calcRanges(seedRangesTmp)
			fmt.Printf("seed-range-intermediate %s %v\n", triplet.title, seedRangesTmp)
		}
		fmt.Printf("seed-range-final %v\n", seedRangesTmp)

		//no more re-ordering possible, so we can just take the minimum
		//the range is irrelevant because it's always positive
		min := math.MaxInt
		for _, seed := range seedRangesTmp {
			if seed.src < min {
				min = seed.src
			}
		}
		fmt.Printf("seed-min-final %v\n", min)
		seedDest = append(seedDest, min)
	}
	fmt.Printf("%v", seedDest)
	println((minInt(seedDest)))
}

func f_5_1() {
	seeds, triplets := parse_5_1()
	fmt.Printf("seeds %v triplets %v (len:%d)\n", seeds, triplets, len(triplets))
	var seedDest []int
	for _, seed := range seeds {
		fmt.Printf("seed-start %d\n", seed)
		for _, triplet := range triplets {
			seed = triplet.calcDest(seed)
			fmt.Printf("seed gone to %d\n", seed)
		}
		fmt.Printf("final seed position %d\n", seed)
		seedDest = append(seedDest, seed)
	}
	fmt.Printf("%v\n", seedDest)
	println((minInt(seedDest)))
}

func f_6_1() {
	//Time:        46     85     75     82
	//Distance:   208   1412   1257   1410
	inputs := []struct {
		time     int
		distance int
	}{
		{46, 208},
		{85, 1412},
		{75, 1257},
		{82, 1410},
	}

	ret := 1
	for _, input := range inputs {
		fmt.Printf("time %d distance %d ", input.time, input.distance)
		wins := 0
		for i := 1; i < input.time; i++ {
			hold := i
			run := input.time - hold
			distance := run * hold
			if distance > input.distance {
				wins++
				continue
			}
			if wins != 0 {
				break
			}
		}
		fmt.Printf("%d wins\n", wins)
		ret *= wins
	}
	println(ret)
}

func f_6_2() {
	p := message.NewPrinter(language.English)
	t := int64(46857582)
	goalDistance := int64(208141212571410)

	minhold := int64(0)
	maxhold := int64(t / 2)
	var hold int64
	for {
		hold = minhold + (maxhold-minhold)/2

		lrun := t - (hold - 1)
		ldistance := lrun * (hold - 1)

		run := t - hold
		distance := run * hold

		grun := t - (hold + 1)
		gdistance := grun * (hold + 1)

		p.Printf("%d: %d @@ %d @@ %d\n", hold, goalDistance-ldistance, goalDistance-distance, goalDistance-gdistance)

		if ldistance < goalDistance && distance >= goalDistance {
			//found where it starts
			break
		}

		if distance < goalDistance && gdistance >= goalDistance {
			hold = hold + 1
			//found where it starts
			break
		}

		//on the upside, need to go left ro find min
		if distance >= goalDistance {
			p.Printf("going left\n")
			maxhold = hold
		} else {
			//on the downside, need to go right to find min
			p.Printf("going right\n")
			minhold = hold
		}
	}
	p.Printf("found max %d @@ %d @@ %d. range is %d\n", hold, t/2, t/2+(t/2-hold), (t/2-hold)*2+1)
}

func main() {
	funcs := map[string]func(){
		"1_1": f_1_1,
		"1_2": f_1_2,
		"2_1": f_2_1,
		"2_2": f_2_2,
		"3_1": f_3_1,
		"3_2": f_3_2,
		"4_1": f_4_1,
		"4_2": f_4_2,
		"5_1": f_5_1,
		"5_2": f_5_2,
		"6_1": f_6_1,
		"6_2": f_6_2,
	}

	funcs[os.Args[1]]()
}
