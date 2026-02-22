/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package integer

// Abs returns the absolute value of an integer
func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
