// Package leap has functions to determine if an integer year is a leap year or not.
package leap

// IsLeapYear returns whether the year (int) entered is a leap year.
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}

	if year%100 == 0 {
		return false
	}

	if year%4 == 0 {
		return true
	}

	return false
}
