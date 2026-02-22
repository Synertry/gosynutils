/*
 *           gosynutils
 *     Copyright (c) Synertry 2025 - 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package jsonx_test

import (
	"strconv"
	"testing"

	"github.com/Synertry/gosynutils/enc/jsonx"
	"github.com/Synertry/gosynutils/gen"
	"github.com/Synertry/gosynutils/math/integer"
)

// TestPrettyPrintSlice tests the PrettyPrint function to ensure it formats
// the input data structure into a human-readable JSON format.
func TestPrettyPrintSlice(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  string
	}{
		"empty slice": {
			input: []string{},
			want:  "[]",
		},
		"single element slice": {
			input: []string{"hello"},
			want: `[
	"hello"
]`,
		},
		"multiple elements slice": {
			input: []string{"hello", "world", "test"},
			want: `[
	"hello",
	"world",
	"test"
]`,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := jsonx.PrettyPrint(tc.input)
			if got != tc.want {
				t.Errorf("expected: %s, got: %s\n", tc.want, got) // %q won't display formatting
				t.Logf("input: %#v", tc.input)
				return
			}
		})
	}
}

// BenchmarkPrettyPrintSlice benchmarks the PrettyPrint function with varying slice lengths
func BenchmarkPrettyPrintSlice(b *testing.B) {
	const maxBenchArrLenExp = 3

	type benchmark struct {
		name string
		len  int
	}

	benchmarks := make([]benchmark, maxBenchArrLenExp+1) // + 1 for empty

	for i := 0; i <= maxBenchArrLenExp; i++ {
		arrLen := integer.Pow(10, i)
		benchmarks[i] = benchmark{name: "ArrLen10^" + strconv.Itoa(i), len: arrLen}
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			input := gen.SliceStrings(bm.len)
			for b.Loop() {
				jsonx.PrettyPrint(input)
			}
		})
	}
}
