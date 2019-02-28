//Package accumulate contains function to perform a string operation on each element of a slice
//then return that slice.
package accumulate

//Accumulate performs a string operation on each element of the input slice.
func Accumulate(input []string, converter func(string) string) []string {
	for i := range input {
		input[i] = converter(input[i])
	}
	return input
}
