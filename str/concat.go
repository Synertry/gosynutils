/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package str

import "strings"

// Concat is an efficient string builder function
func Concat(str ...string) string {
	var sb strings.Builder
	for _, s := range str {
		sb.WriteString(s)
	}
	return sb.String()
}
