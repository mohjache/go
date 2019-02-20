// Package twofer implements function to return statements based on a name
package twofer

// ShareWith returns statement 'one for {name}, one for me'
// if no name is found then 'you' is substituted in.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
