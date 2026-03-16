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
	"math/rand/v2"
)

// GetRand returns a cryptographically secure random number source
func GetRand() *rand.Rand {
	var seed [32]byte
	crand.Read(seed[:]) //nolint:gosec // no error handling is necessary, as Read always succeeds.

	// https://go.dev/blog/chacha8rand#the-chacha8rand-generator
	chacha := rand.NewChaCha8(seed)
	return rand.New(chacha) //nolint:gosec // 8-round form ChaCha8 is secure too
}
