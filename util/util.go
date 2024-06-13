package util

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

func DigitAt(num, pos int) int {
	return ((num % int(math.Pow10(pos))) - (num % int(math.Pow10(pos-1)))) / int(math.Pow10(pos-1))
}

func DoMod11Sum(value, digits int) int {
	sum := 0

	for pos := digits + 1; pos >= 2; pos = pos - 1 {
		num := DigitAt(value, pos-1)
		sum = sum + num*pos
	}

	return DigitAt((sum*10)%11, 1)
}

func Gen(format bool, count int) {
	for i := 0; i < count; i += 1 {
		rnum := rand.Intn(999999999)
		rnum = rnum*10 + DoMod11Sum(rnum, 9)
		rnum = rnum*10 + DoMod11Sum(rnum, 10)

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

func Check(val string) {
	re := regexp.MustCompile(`\D`)

	clean := re.ReplaceAllString(val, "")
	clNum, err := strconv.Atoi(clean)
	if err != nil {
		panic(fmt.Errorf("error parsing string as int: %w", err))
	}
	ct := 0
	for ct < 2 {
		verify := DoMod11Sum(clNum/int(math.Pow10(2-ct)), 9+ct)
		if d := DigitAt(clNum, (ct-2)*-1); d != verify {
			fmt.Printf("Invalid CPF %s: mod11sum for digit at position %d: expected %d, calculated %d\n", val, ct+10, d, verify)
			return
		}
		ct = ct + 1
	}
	fmt.Printf("%s: is valid\n", val)
}
