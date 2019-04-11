package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	var result = []string{}
	for i, word := range rhyme {
		if i == len(rhyme)-1 {
			// result = append(result, "For want of a "+rhyme[i]+" the "+word+" was lost.")
			result = append(result, "And all for the want of a "+rhyme[0]+".")
		} else {
			result = append(result, "For want of a "+word+" the "+rhyme[i+1]+" was lost.")
		}
	}
	return result
}
