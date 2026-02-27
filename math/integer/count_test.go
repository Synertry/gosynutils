/*
 *           gosynutils
 *     Copyright (c) Synertry 2025 - 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package integer_test

import (
	"strconv"
	"testing"

	"github.com/Synertry/gosynutils/math/integer"
)

func TestCountDigits(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"zero": {
			input: 0,
			want:  1,
		},
		"positive single digit": {
			input: 5,
			want:  1,
		},
		"positive two digits": {
			input: 42,
			want:  2,
		},
		"negative single digit": {
			input: -3,
			want:  1,
		},
		"negative two digits": {
			input: -99,
			want:  2,
		},
		"large positive number": {
			input: 1234567890,
			want:  10,
		},
		"large negative number": {
			input: -9876543210,
			want:  10,
		},
		// Leading zeros notify another base notation, here octal instead of decimal
		// "leading zeros": {
		// 	input: 000123, // Leading zeros are ignored in integer representation
		// 	want:  3,      // Count of digits in 123
		// },
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := integer.CountDigits(tc.input)
			if got != tc.want {
				t.Errorf("expected: %d, got: %d\n", tc.want, got)
				t.Logf("input: %d\n", tc.input)
				return
			}
		})
	}
}

func BenchmarkCountDigits(b *testing.B) {
	const maxBenchDigitLength = 8

	type benchmark struct {
		name string
		num  int
	}

	benchmarks := make([]benchmark, maxBenchDigitLength+1) // + 1 for zero

	for i := 0; i <= maxBenchDigitLength; i++ {
		benchmarks[i] = benchmark{name: "Digits" + strconv.Itoa(i+1), num: integer.Pow(10, i)}
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				integer.CountDigits(bm.num)
			}
		})
	}
}
