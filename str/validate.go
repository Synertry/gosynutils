/*
 *           gosynutils
 *     Copyright (c) Synertry 2022 - 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package str

import "unicode"

// IsASCII checks if all characters in the string are ASCII characters
// Source: https://stackoverflow.com/a/53069799/5516320
func IsASCII(s string) bool {
	for i := range len(s) {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
