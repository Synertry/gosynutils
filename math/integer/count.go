/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package integer

// CountDigits returns the number of digits in a number
func CountDigits(num int) (count int) {
	if num == 0 {
		return 1
	}
	num = Abs(num)
	for num > 0 {
		num /= 10
		count++
	}
	return
}
