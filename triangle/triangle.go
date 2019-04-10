package triangle

import (
	"math"
)

// Kind Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

// NaT const
const NaT = "NaT"

// Equ const
const Equ = "Equ"

// Iso const
const Iso = "Iso"

// Sca const
const Sca = "Sca"

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var sides = []float64{a, b, c}

	if anyZeroOrNaN(sides) || degenerateTriangle(sides) {
		return NaT
	}
	if len(Uniq(sides)) == 1 {
		return Equ
	} else if len(Uniq(sides)) == 2 {
		return Iso
	} else if len(Uniq(sides)) == 3 {
		return Sca
	} else {
		return NaT
	}
}

// anyZero length found
func anyZeroOrNaN(list []float64) bool {
	for _, n := range list {
		if math.IsNaN(n) || math.IsInf(n, 1) || n <= 0 {
			return true
		}
	}
	return false
}

// Sum the list of side values
func Sum(list []float64) float64 {
	listSum := 0.0
	for _, n := range list {
		listSum += n
	}
	return listSum
}

func degenerateTriangle(list []float64) bool {
	listSum := Sum(list)
	for _, n := range list {
		if (listSum - n) < n {
			return true
		}
	}
	return false
}

// Uniq the numbers
func Uniq(list []float64) []float64 {
	var uniqList = []float64{}
	for _, n := range list {
		if !Contains(uniqList, n) {
			uniqList = append(uniqList, n)
		}
	}
	return uniqList
}

// Contains Check the number contains
func Contains(list []float64, number float64) bool {
	for _, n := range list {
		if n == number {
			return true
		}
	}
	return false
}
