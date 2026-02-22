/*
 *           gosynutils
 *     Copyright (c) Synertry 2022 - 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package str_test

import (
	"strconv"
	"testing"

	"github.com/Synertry/gosynutils/gen"
	"github.com/Synertry/gosynutils/math/integer"
	"github.com/Synertry/gosynutils/str"
)

func TestReverse(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"simple hello": {
			input: "hello",
			want:  "olleh",
		},
		"empty string": {
			input: "",
			want:  "",
		},
		"single character": {
			input: "a",
			want:  "a",
		},
		"palindrome": {
			input: "racecar",
			want:  "racecar",
		},
		"unicode characters": {
			input: "こんにちは", // "Hello" in Japanese
			want:  "はちにんこ", // Reversed
		},
		"mixed characters": {
			input: "GoSysUtils123",
			want:  "321slitUsySoG",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := str.Reverse(tc.input)
			if got != tc.want {
				t.Errorf("expected: %q, got: %q\n", tc.want, got)
				t.Logf("input: %q\n", tc.input)
				return
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	const maxBenchStrLenExp = 6

	type benchmark struct {
		name string
		len  int
	}

	benchmarks := make([]benchmark, maxBenchStrLenExp+1) // + 1 for empty

	for i := 0; i <= maxBenchStrLenExp; i++ {
		arrLen := integer.Pow(10, i)
		benchmarks[i] = benchmark{name: "StrLen10^" + strconv.Itoa(i), len: arrLen}
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			input := gen.String(bm.len) // Assuming gen.String generates a random string of length bm.len
			for b.Loop() {
				str.Reverse(input)
			}
		})
	}
}
