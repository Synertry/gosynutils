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

func TestAbs(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"positive number": {
			input: 42,
			want:  42,
		},
		"negative number": {
			input: -42,
			want:  42,
		},
		"zero": {
			input: 0,
			want:  0,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := integer.Abs(tc.input)
			if got != tc.want {
				t.Errorf("expected: %d, got: %d\n", tc.want, got)
				t.Logf("input: %d\n", tc.input)
			}
		})
	}
}
