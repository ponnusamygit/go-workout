package dna

import (
	"errors"
	"regexp"
	"strings"
)

// Histogram Custom type
type Histogram map[rune]int

// DNA custom type
type DNA string

func (d DNA) String() string {
	return string(d)
}

// Counts Returns Histogram as result.
func (d DNA) Counts() (Histogram, error) {
	re := regexp.MustCompile(`[^ACGT]`)
	invalidCharacters := len(re.FindString(d.String()))
	if invalidCharacters > 0 {
		return Histogram{}, errors.New("Invalid Charaters")
	}
	h := Histogram{
		'A': strings.Count(d.String(), "A"),
		'C': strings.Count(d.String(), "C"),
		'G': strings.Count(d.String(), "G"),
		'T': strings.Count(d.String(), "T"),
	}
	return h, nil
}
