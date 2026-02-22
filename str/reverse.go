/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package str

import (
	"sync"
)

const threshold = 100 // This value needs to be discussed

var runePool = sync.Pool{
	New: func() interface{} {
		return make([]rune, 0, threshold)
	},
}

// Reverse returns the inverted string of s.
// Implemented through slice inversion, directly
// inverting bytes for small strings.
// Using rune array inversion for large strings.
func Reverse(s string) string {
	if len(s) == 0 {
		return ""
	}

	if IsASCII(s) && len(s) < threshold {
		// Short string direct inversion
		bytes := []byte(s)
		for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
			bytes[i], bytes[j] = bytes[j], bytes[i]
		}
		return string(bytes)
	}

	// Inverting long strings with a rune Array
	runes := runePool.Get().([]rune)
	runes = runes[:0]
	runes = append(runes, []rune(s)...)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	s = string(runes)
	runePool.Put(runes)
	return s
}
