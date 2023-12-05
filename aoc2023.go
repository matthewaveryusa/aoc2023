package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
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

type triplets []triplet

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

func (t *triplet) calcDestRanges(seed int, rng int) ([]rng, bool) {
	//if seed <= t.src+t.rng && seed >= t.src {
	//	return seed + (t.dst - t.src), true
	//}
	//return seed, false

}

func (t *triplets) calcDest(seed int) int {
	for _, triplet := range *t {
		if dest, ok := triplet.calcDest(seed); ok {
			return dest
		}
	}
	return seed

}

func (t *triplets) calcRanges(seed int, rng int) []rng {
	//for _, triplet := range *t {
	//	if dest, ok := triplet.calcDest(seed); ok {
	//		return dest
	//	}
	//}
	//return seed

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
			if len(currentTriplets) != 0 {
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
			return true
		}
		t := triplet{}
		t.parse(line)
		currentTriplets = append(currentTriplets, t)
		return true

	})

	if len(currentTriplets) != 0 {
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
	fmt.Printf("%v", seedDest)
	println((minInt(seedDest)))
}

func f_5_2() {
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
	fmt.Printf("%v", seedDest)
	println((minInt(seedDest)))
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
	}

	funcs[os.Args[1]]()
}
