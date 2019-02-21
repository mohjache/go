// Package twofer implements function to return statements based on a name
package twofer

import "strings"

// ShareWith returns statement 'one for {name}, one for me'
// if no name is found then 'you' is substituted in.
func ShareWith(name string) string {
	var sb strings.Builder

	//check for empty string case and replace with default statement
	if name == "" {
		name = "you"
	}

	//build up the string we want to return and return it
	sb.WriteString("One for ")
	sb.WriteString(name)
	sb.WriteString(", one for me.")

	return sb.String()
}
