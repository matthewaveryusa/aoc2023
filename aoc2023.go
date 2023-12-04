package main

import (
	"bufio"
	"fmt"
	"io"
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

func main() {
	funcs := map[string]func(){
		"1_1": f_1_1,
		"1_2": f_1_2,
		"2_1": f_2_1,
		"2_2": f_2_2,
		"3_1": f_3_1,
		"3_2": f_3_2,
	}

	funcs[os.Args[1]]()
}
