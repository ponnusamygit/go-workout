package twofer

import "strings"

// Function to return two-fer
func ShareWith(name string) string {
	// var name string = name
	// Create a slice and append three strings to it.
	values := []string{}
	values = append(values, "One for ")
	if name == "" {
		values = append(values, "you")
	} else {
		values = append(values, name)
	}
	values = append(values, ", one for me.")

	// Join three strings into one.
	result := strings.Join(values, "")
	return result
}
