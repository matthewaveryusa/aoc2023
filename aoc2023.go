package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func readstdin() string {
	stdindata := make([]byte, 1024)
	output := ""
	n := 0
	var err error
	for n != 0 && err != io.EOF {
		n, err = os.Stdin.Read(stdindata)
		if n != 0 {
			output += string(stdindata[:n])
		}
	}
	return output

}

func readLines(cb func(string) bool) error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() && cb(scanner.Text()) {
	}

	return scanner.Err()
}

func isnum(c byte) bool {
	return c >= '0' && c <= '9'
}

func f_1_1() {
	sum := 0
	_ = readLines(func(line string) bool {
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

	_ = readLines(func(line string) bool {
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

	_ = readLines(func(line string) bool {
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
	_ = readLines(func(line string) bool {
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

func main() {
	funcs := map[string]func(){
		"1_1": f_1_1,
		"1_2": f_1_2,
		"2_1": f_2_1,
		"2_2": f_2_2,
	}

	funcs[os.Args[1]]()
}
