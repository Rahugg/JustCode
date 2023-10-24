package main

import "strings"

type Pair struct {
	first  int
	second string
}

var values = []Pair{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	ans := ""

	for _, p := range values {
		integer, roman := p.first, p.second
		repeat := num / integer
		//returns the amount of copies of roman variable

		ans += strings.Repeat(roman, repeat)
		num -= repeat * integer
	}

	return ans
}
