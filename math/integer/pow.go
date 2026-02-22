/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package integer

// Pow calculates base to the expth power. Since the result is an int, it is assumed that exp is a positive power
// returns 0 if m is negative
func Pow(base, exp int) int {
	if exp < 0 {
		return 0
	} else if exp == 0 {
		return 1
	}
	result := base
	for i := 2; i <= exp; i++ {
		result *= base
	}
	return result
}
