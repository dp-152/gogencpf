package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

var generate bool
var validate string
var format bool
var count int

func init() {
	flag.BoolVar(&generate, "g", true, "")
	flag.BoolVar(&generate, "generate", true, "Generate values (default)")

	// Generate flags
	flag.BoolVar(&format, "f", false, "")
	flag.BoolVar(&format, "format", false, "Format output")
	flag.IntVar(&count, "c", 1, "")
	flag.IntVar(&count, "count", 1, "Amount to generate")

	flag.StringVar(&validate, "v", "", "")
	flag.StringVar(&validate, "validate", "", "")
}

func digitAt(num, pos int) int {
	return ((num % int(math.Pow10(pos))) - (num % int(math.Pow10(pos-1)))) / int(math.Pow10(pos-1))
}

func doMod11Sum(value, digits int) int {
	sum := 0

	for pos := digits + 1; pos >= 2; pos = pos - 1 {
		num := digitAt(value, pos-1)
		sum = sum + num*pos
	}

	return digitAt((sum*10)%11, 1)
}

func gen(count int) {
	for i := 0; i < count; i += 1 {
		rnum := rand.Intn(999999999)
		rnum = rnum*10 + doMod11Sum(rnum, 9)
		rnum = rnum*10 + doMod11Sum(rnum, 10)

		var fmtstr string
		switch format {
		case true:
			var builder strings.Builder
			for idx, d := range fmt.Sprintf("%011d", rnum) {
				idx += 1
				switch {
				case idx%9 == 0:
					builder.WriteRune(d)
					builder.WriteRune('-')
				case idx%3 == 0:
					builder.WriteRune(d)
					builder.WriteRune('.')
				default:
					builder.WriteRune(d)
				}
			}
			fmtstr = builder.String()
		default:
			fmtstr = fmt.Sprintf("%011d", rnum)
		}
		fmt.Println(fmtstr)
	}
}

func check(val string) {
	re := regexp.MustCompile(`\D`)

	clean := re.ReplaceAllString(val, "")
	clNum, err := strconv.Atoi(clean)
	if err != nil {
		panic(fmt.Errorf("error parsing string as int: %w", err))
	}
	ct := 0
	for ct < 2 {
		verify := doMod11Sum(clNum/int(math.Pow10(2-ct)), 9+ct)
		if d := digitAt(clNum, (ct-2)*-1); d != verify {
			fmt.Printf("Invalid CPF %s: mod11sum for digit at position %d: expected %d, calculated %d\n", val, ct+10, d, verify)
			return
		}
		ct = ct + 1
	}
	fmt.Printf("%s: is valid\n", val)
}

func main() {
	flag.Parse()

	switch {
	case validate != "":
		check(validate)
	case generate:
		gen(count)
	}
}
