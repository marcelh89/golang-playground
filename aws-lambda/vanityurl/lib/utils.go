package lib

import "strings"

/*
 * very simple url normalize function
 * to lowercase filter param names
 */
func normalize(text string) string {
	return strings.ToLower(text)
}

/*
 * custom split function to also remove empty elements
 */
func split(c rune) bool {
	return c == '/'
}
