package romannumerals

import (
	"errors"
	"sort"
	"strings"
)

var coreNumbers = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func getCoreNum(number1 int) int {
	var base int
	for _, num1 := range coreNumberKeys() {
		if num1 <= number1 {
			base = num1
		}
	}
	return base
}

// ToRomanNumeral will convert int to roman number
func ToRomanNumeral(number int) (string, error) {
	if number > 3000 || number == 0 || number == -1 {
		return "NaN", errors.New("Not converted")
	}
	var romans []string
	for number > 0 {
		var roman string
		var base = getCoreNum(number)
		var quaterent = number / base
		number = number % base
		for index := 0; index < quaterent; index++ {
			roman = coreNumbers[base]
			romans = append(romans, roman)
		}
	}
	strRoman := strings.Join(romans, "")
	return strRoman, nil
}

func coreNumberKeys() []int {
	var keys []int
	for key := range coreNumbers {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}
