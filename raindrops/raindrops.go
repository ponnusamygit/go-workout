package raindrops

import (
	"fmt"
	"strconv"
	"strings"
)

// Convert will change number to string, pling, plang, plong
func Convert(number int) string {
	fmt.Println(number)
	values := []string{}
	noFactor := true
	for i := 2; i <= number; i++ {
		if number%i == 0 {
			switch i {
			case 3:
				noFactor = false
				values = append(values, "Pling")
			case 5:
				noFactor = false
				values = append(values, "Plang")
			case 7:
				noFactor = false
				values = append(values, "Plong")
			}
		}
	}
	if noFactor {
		values = append(values, strconv.Itoa(number))
	}
	return strings.Join(values, "")
}
