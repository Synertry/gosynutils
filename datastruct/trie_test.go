/*
 *           gosynutils
 *     Copyright (c) Synertry 2025 - 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package datastruct_test

import (
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/Synertry/gosynutils/datastruct"
	"github.com/Synertry/gosynutils/gen"
	"github.com/Synertry/gosynutils/math/integer"
)

// getTestTableTrie is an internal helper functions for generating test cases for the Trie data structure
func getTestTableTrie() map[string]struct {
	words     []string
	search    string
	want      bool
	doesPanic bool
} {
	return map[string]struct {
		words     []string
		search    string
		want      bool
		doesPanic bool // used to test panic cases
	}{
		"no_words": {
			words:  []string{},
			search: "hello",
			want:   false,
		},
		"empty_word": {
			words:  []string{""},
			search: "",
			want:   true,
		},
		"single_word": {
			words:  []string{"hello"},
			search: "hello",
			want:   true,
		},
		"multiple_words": {
			words:  []string{"hello", "world", "trie"},
			search: "world",
			want:   true,
		},
		"empty_search": {
			words:  []string{"hello", "world", "trie"},
			search: "",
			want:   false,
		},
		"mixed_case": {
			words:     []string{"Hello", "World", "Trie"},
			search:    "hello",
			want:      false, // methods should not be responsible for string validation and formatting
			doesPanic: true,  // capital letters are not in lowercase alphabet ascii range, so it should panic
		},
		"with_spaces": {
			words:     []string{" hello ", " world ", " trie "},
			search:    "hello",
			want:      false, // methods should not be responsible for string validation and formatting
			doesPanic: true,  // space is not in alphabet ascii range, so it should panic
		},
		"duplicates": {
			words:  []string{"hello", "hello", "world"},
			search: "hello",
			want:   true, // Duplicates should still be found
		},
		"similar_words": {
			words:  []string{"hello", "hell", "helluva"},
			search: "hell",
			want:   true,
		},
		"long_words": {
			words:  []string{"supercalifragilisticexpialidocious", "antidisestablishmentarianism"},
			search: "supercalifragilisticexpialidocious",
			want:   true,
		},
		"short_words": {
			words:  []string{"a", "b", "c"},
			search: "b",
			want:   true,
		},
		"Random/Strings": {
			words:  gen.SliceStrings(100),
			search: "", // will be handled by the test case
			want:   true,
		},
		"Random/StringsFixed": {
			words:  gen.SliceStringsFixed(100, 10),
			search: "", // will be handled by the test case
			want:   true,
		},
	}
}

//nolint:gocognit
func coreTestTrie(t *testing.T, safe bool, tests map[string]struct {
	words     []string
	search    string
	want      bool
	doesPanic bool
}) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tc.doesPanic {
						t.Errorf("unexpected panic: %v\n", r)
						t.Logf("input words: %#v, search: %q\n", tc.words, tc.search)
						return
					}
				}
			}()

			isMultiTest := strings.HasPrefix(name, "Random")
			trie := datastruct.NewTrie()
			var got bool

			for i, w := range tc.words { // populate the trie with the words
				// separate loop, because of t.Logf(), which outputs all words
				// if we would have it in the same loop, later words in Random tests would still be in mixed case

				if safe {
					trie.SafeAdd(tc.words[i]) // use SafeAdd to normalize the words
				} else {
					if isMultiTest {
						tc.words[i] = strings.ToLower(w) // lowercase the words to ensure consistency
					}
					trie.Add(tc.words[i]) // use Add directly
				}
			}

			for _, w := range tc.words {
				if isMultiTest {
					tc.search = w // set the search term to the current word for non-single tests
				}

				if safe {
					got = trie.SafeFind(tc.search)
				} else {
					got = trie.Find(tc.search)
				}

				if got != tc.want {
					t.Errorf("expected: %t, got: %t\n", tc.want, got)
					t.Logf("input words: %#v, search: %q\n", tc.words, tc.search)
					return
				}

				// test Delete method
				if tc.want && !safe { // only test Delete if the word exists and we are not using SafeAdd
					trie.Delete(tc.search)     // delete the word from the trie
					got = trie.Find(tc.search) // check if the word is still in the trie
					if got {
						t.Errorf("expected: %t, got: %t\n", false, got)
						t.Logf("input words: %#v, search: %q\n", tc.words, tc.search)
						return
					}
				}
				if !isMultiTest {
					break // only search once for single test cases
				}
			}
		})
	}
}

func TestTrie(t *testing.T) {
	// Table-driven test cases
	tests := getTestTableTrie()

	/*
		// remove empty strings from RandomStrings
		for idx, word := range tests["RandomStrings"] {
			if len(word) == 0 {
				// slices.Delete uses slow but ordered "zeroing" delete: https://github.com/golang/go/blob/71c2bf551303930faa32886446910fa5bd0a701a/src/slices/slices.go#L230
				// tests["RandomStrings"] = slices.Delete(tests["RandomStrings"], idx, idx+1)

				// our swap and reslice is faster
				tests["RandomStrings"] = Slice.RemoveIndex(tests["RandomStrings"], idx)
			}
		}
	*/

	coreTestTrie(t, false, tests)
}

func TestTrieSafe(t *testing.T) {
	// Table-driven test cases
	tests := getTestTableTrie()

	// change want value for mixed_case and with_spaces tests
	for name, tc := range tests {
		if slices.Contains([]string{"mixed_case", "with_spaces"}, name) {
			tc.want = true
			tests[name] = tc
		} else if name == "empty_word" {
			// SafeAdd does not add empty strings, so we need to change the want value
			tc.want = false
			tests[name] = tc
		}
	}

	coreTestTrie(t, true, tests)
}

func BenchmarkTrieAdd(b *testing.B) {
	const maxBenchArrLenExp = 4
	type benchmark struct {
		name string
		len  int
	}
	benchmarks := make([]benchmark, maxBenchArrLenExp+1) // do not use maps! Order will be randomized; + 1 for 2^0

	for i := 0; i <= maxBenchArrLenExp; i++ {
		arrLen := integer.Pow(10, i)
		benchmarks[i] = benchmark{name: "ArrLen10^" + strconv.Itoa(i), len: arrLen}
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			trie, words := datastruct.NewTrie(), gen.SliceStringsFixed(bm.len, 10)
			for i, w := range words { // lowercase the words to ensure consistency
				words[i] = strings.ToLower(w)
			}
			for b.Loop() {
				for _, w := range words {
					trie.Add(w)
				}
			}
		})
	}
}

func BenchmarkTrieFind(b *testing.B) {
	const maxBenchStrLenExp = 3
	type benchmark struct {
		name string
		len  int
	}
	benchmarks := make([]benchmark, maxBenchStrLenExp+1) // + 1 for single 10^0 -> 1

	for i := 0; i <= maxBenchStrLenExp; i++ { // -1 as start, because substraction is more costly than addition
		strLen := integer.Pow(10, i)
		benchmarks[i] = benchmark{name: "StrLen10^" + strconv.Itoa(i), len: strLen}
	}

	// out of loop, to ensure consistency in wordpool between search benchmarks
	trie, words := datastruct.NewTrie(), gen.SliceStringsFixed(1000, 10)
	for _, w := range words { // lowercase the words to ensure consistency
		trie.Add(strings.ToLower(w))
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				trie.Find(strings.ToLower(gen.String(bm.len)))
			}
		})
	}
}
