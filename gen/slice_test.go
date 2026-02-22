/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package gen_test

import (
	"fmt"
	"testing"

	"github.com/Synertry/gosynutils/enc/jsonx"
	"github.com/Synertry/gosynutils/gen"
	"github.com/Synertry/gosynutils/math/integer"
)

// TestSliceStrings tests if the SliceStrings function has the expected length of the slice
// and if the strings in the slice have the expected incrementing lengths depending on the index.
func TestSliceStrings(t *testing.T) {
	const maxTestArrLenExp = 4

	type test struct {
		name string
		len  int
	}

	tests := make([]test, maxTestArrLenExp+1)

	for i := range maxTestArrLenExp {
		arrLen := integer.Pow(10, i)
		tests[i] = test{name: fmt.Sprintf("ArrLen10^%d", i), len: arrLen}
	}
	tests[maxTestArrLenExp] = test{name: "Empty", len: 0}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := gen.SliceStrings(tc.len)

			if len(slice) != tc.len {
				t.Errorf("expected: %d, got: %d\n", tc.len, len(slice))
				t.Logf("input: len(`%s`)\n", jsonx.PrettyPrint(slice))
				return
			}

			for i, str := range slice {
				expectedLen := i + 1 // Length should be incrementing by 1
				if len(str) != expectedLen {
					t.Errorf("expected: %d, got: %d at index %d\n", expectedLen, len(str), i)
					t.Logf("input: %q\n", str)
					return
				}
			}
		})
	}
}

// TestSliceStringsFixed tests if the SliceStringsFixed function has the expected length of the slice
// and if the strings in the slice have the expected fixed length.
func TestSliceStringsFixed(t *testing.T) {
	const (
		maxTestArrLenExp = 4
		maxTestStrLenExp = 4
	)

	type test struct {
		name   string
		arrLen int
		strLen int
	}

	tests := make([]test, (maxTestArrLenExp+1)*(maxTestStrLenExp+1)) // + 1 for empty

	for i := range maxTestArrLenExp + 1 {
		arrLen := integer.Pow(10, i)
		for j := range maxTestStrLenExp + 1 {
			strLen := integer.Pow(10, j)
			tests[i*(maxTestStrLenExp+1)+j] = test{
				name:   fmt.Sprintf("ArrLen10^%d/StrLen10^%d", i, j),
				arrLen: arrLen,
				strLen: strLen,
			}
		}
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := gen.SliceStringsFixed(tc.arrLen, tc.strLen)

			if len(slice) != tc.arrLen {
				t.Errorf("expected: %d, got: %d\n", tc.arrLen, len(slice))
				t.Logf("input: len(`%s`)\n", jsonx.PrettyPrint(slice))
				return
			}

			for _, str := range slice {
				if len(str) != tc.strLen {
					t.Errorf("expected: %d, got: %d\n", tc.strLen, len(str))
					t.Logf("input: %q\n", str)
					return
				}
			}
		})
	}
}

// BenchmarkSliceStrings benchmarks the performance of the SliceStrings function
func BenchmarkSliceStrings(b *testing.B) {
	const maxBenchArrLenExp = 3

	type benchmark struct {
		name string
		len  int
	}

	benchmarks := make([]benchmark, maxBenchArrLenExp+1) // + 1 for empty

	for i := range maxBenchArrLenExp + 1 {
		arrLen := integer.Pow(10, i)
		benchmarks[i] = benchmark{name: fmt.Sprintf("ArrLen10^%d", i), len: arrLen}
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				gen.SliceStrings(bm.len)
			}
		})
	}
}

// BenchmarkSliceStringsFixed benchmarks the performance of the SliceStringsFixed function
func BenchmarkSliceStringsFixed(b *testing.B) {
	const (
		maxBenchArrLenExp = 3
		maxBenchStrLenExp = 3
	)

	type benchmark struct {
		name   string
		arrLen int
		strLen int
	}

	benchmarks := make([]benchmark, (maxBenchArrLenExp+1)*(maxBenchStrLenExp+1)) // + 1 for empty

	for i := range maxBenchArrLenExp + 1 {
		arrLen := integer.Pow(10, i)
		for j := range maxBenchStrLenExp + 1 {
			strLen := integer.Pow(10, j)
			benchmarks[i*(maxBenchStrLenExp+1)+j] = benchmark{
				name:   fmt.Sprintf("ArrLen10^%d/StrLen^%d", i, j),
				arrLen: arrLen,
				strLen: strLen,
			}
		}
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				gen.SliceStringsFixed(bm.arrLen, bm.strLen)
			}
		})
	}
}
