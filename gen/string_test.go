/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package gen_test

import (
	"strconv"
	"testing"

	"github.com/Synertry/gosynutils/math/integer"

	"github.com/Synertry/gosynutils/gen"
)

func TestString_Len(t *testing.T) {
	const maxTestArrLen = 100

	type test struct {
		input string
		want  int
	}

	tests := map[string]test{
		"Empty":  {input: "", want: 0},
		"Single": {input: "a", want: 1},
		"Simple": {input: "hello", want: 5},
	}

	// Generate tests for random strings of varying lengths
	for i := range maxTestArrLen {
		input := gen.String(i)
		tests["Random"+strconv.Itoa(i)] = test{input: input, want: len(input)}
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := len(tc.input)
			if got != tc.want {
				t.Errorf("expected: %d, got: %d", tc.want, got)
				t.Logf("input: %#v\n", tc.input)
				return
			}
		})
	}
}

func TestString_Pattern(t *testing.T) {
	// random := gen.GetRand()
	sLen := gen.GetRand().Intn(100)
	str := gen.String(sLen)
	for i, r := range str {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			t.Errorf("expected: char in range [a-Z], got: %c\n", r)
			t.Logf("input: %s[%d]\n", str, i)
			return
		}
	}
}

func TestString_Race(_ *testing.T) {
	const numGoroutines = 10
	const numIterations = 100

	done := make(chan bool)
	for range numGoroutines {
		go func() {
			for range numIterations {
				_ = gen.String(10)
			}
			done <- true
		}()
	}

	for range numGoroutines {
		<-done
	}
}

func BenchmarkString(b *testing.B) {
	const maxBenchStrLenExp = 4

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
			for b.Loop() {
				gen.String(bm.len)
			}
		})
	}
}
