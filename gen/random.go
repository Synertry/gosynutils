/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package gen

import (
	crand "crypto/rand"
	"math/big"
	"math/rand" //nolint:depguard // legacy reasons, TODO: upgrade to v2
)

var (
	// Random is a global instance of a cryptographically secure random number generator
	Random *rand.Rand
)

// GetRand returns a cryptographically secure random number source
func GetRand() *rand.Rand {
	// get random seed from crypto/rand
	cnum, err := crand.Int(crand.Reader, big.NewInt(1<<63-1))
	if err != nil {
		//log.Panic(err)
		panic(err)
	}
	return rand.New(rand.NewSource(cnum.Int64()))
}
