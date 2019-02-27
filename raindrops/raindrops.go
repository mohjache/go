//Package raindrops takes integers and returns strings based on what factors it has.
package raindrops

import "strconv"

//Convert takes an input of type int and will return a string statement if it has a
// factor of 3, 5 or 7. Else it returns the stringified input
func Convert(input int) string {

	//cover edge case first to see if input is not divisible by 3,5,7
	if input%3 != 0 &&
		input%5 != 0 &&
		input%7 != 0 {
		return strconv.Itoa(input)
	}

	//instantiate value now as we are about to use it.
	var result string

	if input%3 == 0 {
		result += "Pling"
	}

	if input%5 == 0 {
		result += "Plang"
	}

	if input%7 == 0 {
		result += "Plong"
	}

	return result
}
