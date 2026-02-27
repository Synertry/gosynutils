/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package str_test

import (
	"testing"

	"github.com/Synertry/gosynutils/str"
)

func TestIsASCII(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		input string
		want  bool
	}{
		"valid ASCII": {
			input: "Hello, World!",
			want:  true,
		},
		"invalid ASCII": {
			input: "こんにちは", // Japanese characters
			want:  false,
		},
		"empty string": {
			input: "",
			want:  true, // Empty string is considered valid ASCII
		},
		"single ASCII character": {
			input: "A",
			want:  true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := str.IsASCII(tc.input)
			if got != tc.want {
				t.Errorf("expected: %v, got: %v\n", tc.want, got)
				t.Logf("input: %q\n", tc.input)
				return
			}
		})
	}
}
