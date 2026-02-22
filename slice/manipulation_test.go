/*
 *           gosynutils
 *     Copyright (c) Synertry 2025 - 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package slice_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/Synertry/gosynutils/gen"
	"github.com/Synertry/gosynutils/math/integer"
	"github.com/Synertry/gosynutils/slice"
)

// TestRemoveIndex tests the RemoveIndex function to ensure it correctly removes an element at a specified index
func TestRemoveIndex(t *testing.T) {
	tests := map[string]struct {
		slice []int
		index int
		want  []int
	}{
		"empty slice": {
			slice: []int{},
			index: 0,
			want:  []int{},
		},
		"single element": {
			slice: []int{1},
			index: 0,
			want:  []int{},
		},
		"remove first element": {
			slice: []int{1, 2, 3},
			index: 0,
			want:  []int{3, 2},
		},
		"remove last element": {
			slice: []int{1, 2, 3},
			index: 2,
			want:  []int{1, 2},
		},
		"remove middle element": {
			slice: []int{1, 2, 3, 4},
			index: 1,
			want:  []int{1, 4, 3},
		},
		"index out of bounds": {
			slice: []int{1, 2, 3},
			index: 5,
			want:  []int{1, 2, 3}, // should return original
		},
		"negative index": {
			slice: []int{1, 2, 3},
			index: -1,
			want:  []int{1, 2, 3}, // should return original
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := slice.RemoveIndex(tc.slice, tc.index)
			if len(got) != len(tc.want) {
				t.Errorf("expected length: %d, got: %d\n", len(tc.want), len(got))
				t.Logf("input slice: %#v, index: %d\n", tc.slice, tc.index)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected slice: %#v, got: %#v\n", tc.want, got)
				t.Logf("input slice: %#v, index: %d\n", tc.slice, tc.index)
				return
			}
		})
	}
}

// TestRemoveElements tests the RemoveElements function to ensure
// it correctly removes all occurrences of a specified element from the slice.
func TestRemoveElements(t *testing.T) {
	tests := map[string]struct {
		slice []int
		elem  int
		want  []int
	}{
		"empty slice": {
			slice: []int{},
			elem:  0,
			want:  []int{},
		},
		"no occurrences": {
			slice: []int{1, 2, 3},
			elem:  4,
			want:  []int{1, 2, 3},
		},
		"simple case": {
			slice: []int{1, 2, 2, 3},
			elem:  2,
			want:  []int{1, 3},
		},
		"element at start": {
			slice: []int{2, 1, 2, 3},
			elem:  2,
			want:  []int{3, 1}, // 3 at the start, because we move the last element to the index of the removed element
		},
		"element at end": {
			slice: []int{1, 2, 3, 2},
			elem:  2,
			want:  []int{1, 3},
		},
		"all same elements": {
			slice: []int{2, 2, 2},
			elem:  2,
			want:  []int{},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := slice.RemoveElements(tc.slice, tc.elem)
			if len(got) != len(tc.want) {
				t.Errorf("expected length: %d, got: %d\n", len(tc.want), len(got))
				t.Logf("input slice: %#v, elem: %d\n", tc.slice, tc.elem)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected slice: %#v, got: %#v\n", tc.want, got)
				t.Logf("input slice: %#v, elem: %d\n", tc.slice, tc.elem)
				return
			}
		})
	}
}

func TestInvert(t *testing.T) {
	tests := map[string]struct {
		slice []int
		want  []int
	}{
		"empty slice": {
			slice: []int{},
			want:  []int{},
		},
		"single element": {
			slice: []int{1},
			want:  []int{1},
		},
		"multiple elements": {
			slice: []int{1, 2, 3, 4},
			want:  []int{4, 3, 2, 1},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := slice.Invert(tc.slice)
			if len(got) != len(tc.want) {
				t.Errorf("expected length: %d, got: %d\n", len(tc.want), len(got))
				t.Logf("input slice: %#v\n", tc.slice)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected slice: %#v, got: %#v\n", tc.want, got)
				t.Logf("input slice: %#v\n", tc.slice)
				return
			}
		})
	}
}

func BenchmarkRemoveIndex(b *testing.B) {
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
			input := []byte(gen.String(bm.len))
			for b.Loop() {
				slice.RemoveIndex(input, 0) // Remove the middle element
			}
		})
	}
}

func BenchmarkRemoveElements(b *testing.B) {
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
			input := []byte(gen.String(bm.len))
			for b.Loop() {
				slice.RemoveElements(input, input[0]) // Remove all occurrences of the first element
			}
		})
	}
}

func BenchmarkInvert(b *testing.B) {
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
			input := []byte(gen.String(bm.len))
			for b.Loop() {
				slice.Invert(input)
			}
		})
	}
}
