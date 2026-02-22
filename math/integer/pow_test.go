/*
 *           gosynutils
 *     Copyright (c) Synertry 2025 - 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package integer_test

import (
	"testing"

	"github.com/Synertry/gosynutils/math/integer"
)

func TestPow(t *testing.T) {
	tests := map[string]struct {
		base     int
		exponent int
		want     int
	}{
		"zero exponent": {
			base:     2,
			exponent: 0,
			want:     1,
		},
		"positive exponent": {
			base:     2,
			exponent: 3,
			want:     8,
		},
		"negative exponent": {
			base:     2,
			exponent: -3,
			want:     0,
		},
		"negative base": {
			base:     -2,
			exponent: 3,
			want:     -8,
		},
		"zero base": {
			base:     0,
			exponent: 5,
			want:     0,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := integer.Pow(tc.base, tc.exponent)
			if got != tc.want {
				t.Errorf("expected: %d, got: %d\n", tc.want, got)
				t.Logf("input base: %d, exponent: %d\n", tc.base, tc.exponent)
				return
			}
		})
	}
}
