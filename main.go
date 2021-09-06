package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	from := os.Args[1]
	loop, _ := strconv.Atoi(os.Args[2])
	step, _ := strconv.Atoi(os.Args[3])
	for i := 0; i < loop; i++ {
		fmt.Println(plusIntToStr(from, step*i))
	}
}

func plusIntToStr(str string, num int) string {
	runeStr := []rune(str)
	numLen := intLen(num)
	pivot := len(runeStr) - numLen
	preRune := runeStr[:pivot]
	postRune := runeStr[pivot:]
	calResult := runeToInt(postRune) + num
	if intLen(calResult) > numLen {
		return incrementRunes(preRune) + strconv.Itoa(calResult)[1:]
	} else {
		return str[:pivot] + strconv.Itoa(calResult)
	}
}

func intLen(num int) int {
	length := 1
	for num >= 10 {
		num = num / 10
		length++
	}
	return length
}

func runeToInt(runes []rune) int {
	result := 0
	j := len(runes) - 1
	for i := 0; i < len(runes); i++ {
		result += sqrt(10, j) * int(runes[i]-'0')
		j--
	}
	return result
}

func sqrt(num int, time int) int {
	if time == 0 {
		return 1
	}
	result := 1
	for i := 0; i < time; i++ {
		result *= num
	}
	return result
}

func incrementRunes(runes []rune) string {
	incCount := 0
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == '9' {
			runes[i] = '0'
			incCount++
		} else {
			runes[i] = runes[i] + 1
			break
		}
	}
	if incCount == len(runes) {
		return "1" + string(runes)
	}
	return string(runes)
}
