/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package gen_test

import (
	"math/rand/v2"
	"testing"

	"github.com/Synertry/gosynutils/gen"
)

// TestGetRand tests the GetRand function to check whether we successfully created a *[rand.Rand] object
// and that it is not nil. This is a basic test to ensure the function works as
func TestGetRand(t *testing.T) {
	t.Run("GetRand", func(t *testing.T) {
		random := gen.GetRand()
		if !isTypeRand(random) {
			t.Errorf("expected: type *rand.Rand, got %T\n", random)
		}
	})
}

// isTypeRand is a private helper function to check if the provided interface is of type *[rand.Rand]
func isTypeRand(t any) bool {
	switch t.(type) {
	case *rand.Rand:
		return true
	default:
		return false
	}
}

// func BenchmarkGetRand benchmarks the GetRand function. No params, so no loop count is needed.
func BenchmarkGetRand(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		gen.GetRand()
	}
}
