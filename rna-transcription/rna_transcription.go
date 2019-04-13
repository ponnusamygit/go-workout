package strand

import (
	"strings"
)

// ToRNA is used to convert DNA
func ToRNA(dna string) string {
	RNAMap := map[string]string{"G": "C", "C": "G", "T": "A", "A": "U"}
	var result string
	for _, svar := range strings.Split(dna, "") {
		result = result + RNAMap[svar]
	}
	return result
}
