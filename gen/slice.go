/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package gen

// SliceStrings generates a slice of random strings with the given length.
// gen.String generates the strings.
// The length of each string is increasing for each position.
// Index 0 has a string of length 1, index 100 has a string of length 101.
// Removed dividing length by 2 from the original code to improve performance.
func SliceStrings(lenSlice int) []string {
	slice := make([]string, lenSlice)
	for i := range lenSlice {
		slice[i] = String(i + 1)
	}
	return slice
}

// SliceStringsFixed generates a slice of strings with the given length lenSlice.
// gen.String generates the strings.
// The length of each string is defined by the passed lenString.
// This makes it possible to generate a slice of letters:
// Example: SliceStringsFixed(5, 1) -> ["a", "c", "e", "g", "i"]
// or a slice of strings with the same length:
// Example: SliceStringsFixed(5, 3) -> ["gea", "brv", "pgd", "qwo", "cie"]
func SliceStringsFixed(lenSlice, lenString int) []string {
	slice := make([]string, lenSlice)
	for i := range lenSlice {
		slice[i] = String(lenString)
	}
	return slice
}
