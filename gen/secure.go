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

// ErrNonPositiveLength is returned when a requested byte length is zero or negative.
// Generating an empty token is never what the caller meant, so it is refused
// rather than silently returning an empty string that would pass as a credential.
var ErrNonPositiveLength = errors.New("gen: byte length must be greater than zero")

// SecureBytes returns n cryptographically secure random bytes.
//
// Unlike [GetRand], which seeds a fast userspace generator, this reads directly
// from the operating system source on every call. Use it for security material
// such as tokens, nonces, salts and keys. Use [GetRand] for simulation,
// sampling, jitter and other non-secret randomness where throughput matters.
//
// It returns [ErrNonPositiveLength] if n is not greater than zero.
func SecureBytes(n int) (b []byte, err error) {
	if n <= 0 {
		return nil, fmt.Errorf("%w, got %d", ErrNonPositiveLength, n)
	}

	b = make([]byte, n)
	if _, err = crand.Read(b); err != nil {
		// Unreachable on current Go, where crypto/rand.Read never fails, but
		// the error is surfaced rather than discarded so a future change in
		// that guarantee cannot silently hand out predictable bytes.
		return nil, fmt.Errorf("gen: read random bytes: %w", err)
	}
	return b, nil
}

// SecureToken returns n random bytes encoded as unpadded base64url.
//
// Note that n counts the bytes of entropy, not the length of the returned
// string: 32 bytes yields a 43 character token. 32 is a sound default for a
// bearer credential.
//
// The encoding is URL and filename safe and free of padding, so the result
// needs no escaping in a URL, a form field, an HTTP header or a filename.
func SecureToken(n int) (token string, err error) {
	var b []byte
	if b, err = SecureBytes(n); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// SecureHex returns n random bytes encoded as lowercase hexadecimal.
//
// As with [SecureToken], n counts bytes rather than output characters, so the
// returned string is twice as long as n. Prefer [SecureToken] when the value
// only has to round trip, and this when a human has to read, compare or dictate
// the value, since hex has no case ambiguity.
func SecureHex(n int) (s string, err error) {
	var b []byte
	if b, err = SecureBytes(n); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
