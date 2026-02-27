/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package slice

// Contains searches for an element in a slice and returns true if found
// Source: https://stackoverflow.com/a/70802740/5516320
//
// Deprecated: As of Go 1.21, you can use the stdlib slices package,
// which was promoted from the experimental package:
// https://stackoverflow.com/a/71181131/5516320
//
// This makes this function redundant, instead use
// slices.Contains()
//
//goland:noinspection GoDeprecation
//nolint:godoclint // We showcase a legacy function for educational purposes
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s { //nolint:modernize // We showcase a legacy function for educational purposes
		if v == e {
			return true
		}
	}
	return false
}
