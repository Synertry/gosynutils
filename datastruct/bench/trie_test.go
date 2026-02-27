/*
 *           gosynutils
 *     Copyright (c) Synertry 2025 - 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

// package bench_test exists to benchmark applications against other public trie libraries,
// but it is separated from the datastruct package to keep the main library dependency free
package bench_test //nolint:cyclop // It is a table-driven benchmark, what am I supposed to do about the complexity?

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Synertry/gosynutils/datastruct"
	"github.com/Synertry/gosynutils/gen"
	"github.com/Synertry/gosynutils/math/integer"

	dgtrie "github.com/dghubble/trie"
	hctrie "github.com/hashicorp/go-immutable-radix"
	"github.com/sauerbraten/radix"
)


func BenchmarkComparison(b *testing.B) {
	const (
		maxBenchArrLenExp = 4
		maxBenchStrLenExp = 2
	)
	type benchmark struct {
		name   string
		arrLen int
		strLen int
	}

	benchmarks := make([]benchmark, (maxBenchArrLenExp+1)*(maxBenchStrLenExp+1)) // + 1 for empty

	for i := 4; i <= maxBenchArrLenExp; i++ {
		arrLen := integer.Pow(10, i)
		for j := 2; j <= maxBenchStrLenExp; j++ {
			strLen := integer.Pow(10, j)
			benchmarks[i*(maxBenchStrLenExp+1)+j] = benchmark{
				name:   fmt.Sprintf("ArrLen10^%d/StrLen10^%d", i, j),
				arrLen: arrLen,
				strLen: strLen,
			}
		}
	}

	// remove empty benchmarks
	var s []benchmark
	for _, bm := range benchmarks {
		if bm.name != "" {
			s = append(s, bm)
		}
	}
	benchmarks = s // reassign to remove empty benchmarks

	for _, bm := range benchmarks {
		words := gen.SliceStringsFixed(bm.arrLen, bm.strLen)
		search := strings.ToLower(gen.String(bm.strLen))
		synTrie := datastruct.NewTrie()
		dgTrie := dgtrie.NewRuneTrie()
		sbTrie := radix.New()
		hcTrie := hctrie.New()

		for i, w := range words { // lowercase the words to ensure consistency
			words[i] = strings.ToLower(w)
			synTrie.Add(words[i])
			dgTrie.Put(words[i], true)                           // dgTrie uses a boolean value for the value
			sbTrie.Set(words[i], true)                           // sbTrie uses a boolean value for the value
			hcTrie, _, _ = hcTrie.Insert([]byte(words[i]), true) // hcTrie uses a boolean value for the value
		}

		b.Run("Insert/Synertry/"+bm.name, func(b *testing.B) {
			b.ReportAllocs()
			trie := datastruct.NewTrie()
			for b.Loop() {
				for _, w := range words {
					trie.Add(w)
				}
			}
		})
		b.Run("Insert/DgHubble/"+bm.name, func(b *testing.B) {
			b.ReportAllocs()
			trie := dgtrie.NewRuneTrie()
			for b.Loop() {
				for _, w := range words {
					trie.Put(w, true) // dgTrie uses a boolean value for the value
				}
			}
		})
		b.Run("Insert/Sauerbraten/"+bm.name, func(b *testing.B) {
			b.ReportAllocs()
			trie := radix.New()
			for b.Loop() {
				for _, w := range words {
					trie.Set(w, true) // dgTrie uses a boolean value for the value
				}
			}
		})
		b.Run("Insert/Hashicorp/"+bm.name, func(b *testing.B) {
			b.ReportAllocs()
			trie := hctrie.New()
			for b.Loop() {
				for _, w := range words {
					trie, _, _ = trie.Insert([]byte(w), true)
				}
			}
		})

		b.Run("Search/Synertry/"+bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				synTrie.Find(search)
			}
		})
		b.Run("Search/DgHubble/"+bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				dgTrie.Get(search)
			}
		})
		b.Run("Search/Sauerbraten/"+bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				sbTrie.Get(search)
			}
		})
		b.Run("Search/Hashicorp/"+bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				hcTrie.Get([]byte(search))
			}
		})
	}
}
