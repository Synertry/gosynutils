/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package slice

// RemoveIndex removes the element at the specified index from the slice without preserving the order of elements.
// The last element is moved to the index of the removed element.
func RemoveIndex[T comparable](s []T, idx int) []T {
	if idx < 0 || idx >= len(s) {
		return s // Index out of bounds, return original slice
	}
	s[idx] = s[len(s)-1] // Move the last element to the index
	return s[:len(s)-1]  // Return the slice without the last element
}

// RemoveElements removes all occurrences of the specified element from the slice.
// It iterates through the slice and uses RemoveIndex to remove each occurrence.
func RemoveElements[T comparable](s []T, e T) []T {
	for i := range s {
		for s[i] == e { // ensures to remove all occurrences of the element even if it's get fed from the tail consecutively
			s = RemoveIndex(s, i)
			if i >= len(s) {
				return s // If the slice is empty, return it
			}
		}
		if i+1 >= len(s) { // If we reach the end of the slice, break
			break
		}
	}
	return s
}

// Invert reverses the order of elements in the slice.
// It swaps elements from the start and end of the slice until it reaches the middle.
func Invert[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
