/*
 *           gosynutils
 *     Copyright (c) Synertry 2025 - 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package slice_test

import (
	"testing"

	"github.com/Synertry/gosynutils/slice"
)

func TestContains(t *testing.T) {
	tests := map[string]struct {
		slice []int
		value int
		want  bool
	}{
		"empty slice": {
			slice: []int{},
			value: 1,
			want:  false,
		},
		"single element present": {
			slice: []int{1},
			value: 1,
			want:  true,
		},
		"single element absent": {
			slice: []int{2},
			value: 1,
			want:  false,
		},
		"multiple elements with value present": {
			slice: []int{1, 2, 3},
			value: 2,
			want:  true,
		},
		"multiple elements with value absent": {
			slice: []int{4, 5, 6},
			value: 3,
			want:  false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			//goland:noinspection GoDeprecation
			got := slice.Contains(tc.slice, tc.value) // We are testing deprecated function here. Don't remove it
			if got != tc.want {
				t.Errorf("expected: %v, got: %v\n", tc.want, got)
				t.Logf("slice: %v, value: %d\n", tc.slice, tc.value)
				return
			}
		})
	}
}
