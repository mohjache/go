package hamming

import (
	"errors"
	"strings"
)

//Distance finds the number of different characters in strings of equal length.
func Distance(a, b string) (int, error) {
	// strings a and b need to be of the same length.
	if len(a) != len(b) {
		return 0, errors.New("strings a and b are different lengths")
	}

	hammingDistance := 0

	arrayA := strings.Split(a, "")
	arrayB := strings.Split(b, "")

	for index := range arrayA {
		if arrayA[index] != arrayB[index] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
