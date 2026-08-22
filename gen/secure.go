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
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

// ErrNonPositiveLength is returned when n is zero or negative.
var ErrNonPositiveLength = errors.New("gen: byte length must be greater than zero")

// SecureBytes returns n cryptographically secure random bytes, read from the
// OS source on every call. Use [GetRand] instead for non-secret randomness.
// Returns [ErrNonPositiveLength] if n <= 0.
func SecureBytes(n int) (b []byte, err error) {
	if n <= 0 {
		return nil, fmt.Errorf("%w, got %d", ErrNonPositiveLength, n)
	}

	b = make([]byte, n)
	if _, err = crand.Read(b); err != nil {
		return nil, fmt.Errorf("gen: read random bytes: %w", err)
	}
	return b, nil
}

// SecureToken returns n random bytes as unpadded base64url. n counts bytes of
// entropy, not output length: 32 bytes yields a 43 character token.
func SecureToken(n int) (token string, err error) {
	var b []byte
	if b, err = SecureBytes(n); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// SecureHex returns n random bytes as lowercase hex. Output is twice as long
// as n. Prefer [SecureToken] unless the value needs to be human-readable.
func SecureHex(n int) (s string, err error) {
	var b []byte
	if b, err = SecureBytes(n); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
