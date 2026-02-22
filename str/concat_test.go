/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package str_test

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/Synertry/gosynutils/gen"
	"github.com/Synertry/gosynutils/math/integer"
	"github.com/Synertry/gosynutils/str"
)

func TestConcat(t *testing.T) {
	const maxTestArrLen = 100

	type test struct {
		input []string
		want  string
	}

	tests := map[string]test{
		"Empty":    {input: []string{}, want: ""},
		"Single":   {input: []string{"a"}, want: "a"},
		"Multiple": {input: []string{"a", "b", "c"}, want: "abc"},
	}

	for i := 0; i < maxTestArrLen; i++ {
		input := make([]string, i)
		var want bytes.Buffer
		for j := 0; j < i; j++ {
			input[j] = gen.String(i)
			want.WriteString(input[j])
		}
		tests["Random"+strconv.Itoa(i)] = test{input: input, want: want.String()}
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := str.Concat(tc.input...)
			if got != tc.want {
				t.Errorf("expected: %s, got: %s\n", tc.want, got)
				t.Logf("input: %#v\n", tc.input)
				return
			}
		})
	}
}

func BenchmarkConcat(b *testing.B) {
	const maxBenchArrLenExp = 4

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
			input := gen.SliceStringsFixed(bm.len, 5) // Fixed string length of 5 for consistency
			for b.Loop() {
				str.Concat(input...)
			}
		})
	}
}
