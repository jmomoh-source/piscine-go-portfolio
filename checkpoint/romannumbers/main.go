package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RomanPart struct {
	val   int
	calc  string
	roman string
}

var parts = []RomanPart{
	{3000, "M+M+M", "MMM"},
	{2000, "M+M", "MM"},
	{1000, "M", "M"},
	{900, "(M-C)", "CM"},
	{800, "D+C+C+C", "DCCC"},
	{700, "D+C+C", "DCC"},
	{600, "D+C", "DC"},
	{500, "D", "D"},
	{400, "(D-C)", "CD"},
	{300, "C+C+C", "CCC"},
	{200, "C+C", "CC"},
	{100, "C", "C"},
	{90, "(C-X)", "XC"},
	{80, "L+X+X+X", "LXXX"},
	{70, "L+X+X", "LXX"},
	{60, "L+X", "LX"},
	{50, "L", "L"},
	{40, "(L-X)", "XL"},
	{30, "X+X+X", "XXX"},
	{20, "X+X", "XX"},
	{10, "X", "X"},
	{9, "(X-I)", "IX"},
	{8, "V+I+I+I", "VIII"},
	{7, "V+I+I", "VII"},
	{6, "V+I", "VI"},
	{5, "V", "V"},
	{4, "(V-I)", "IV"},
	{3, "I+I+I", "III"},
	{2, "I+I", "II"},
	{1, "I", "I"},
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 || num >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	var calcs []string
	var romans []string

	n := num
	for _, p := range parts {
		if n >= p.val {
			calcs = append(calcs, p.calc)
			romans = append(romans, p.roman)
			n -= p.val
		}
	}

	fmt.Println(strings.Join(calcs, "+"))
	fmt.Println(strings.Join(romans, ""))
}
