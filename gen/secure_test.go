/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package gen_test

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strings"
	"testing"

	"github.com/Synertry/gosynutils/gen"
)

func TestSecureBytes(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		n       int
		wantLen int
		wantErr bool
	}{
		"one":      {n: 1, wantLen: 1, wantErr: false},
		"sixteen":  {n: 16, wantLen: 16, wantErr: false},
		"token":    {n: 32, wantLen: 32, wantErr: false},
		"zero":     {n: 0, wantLen: 0, wantErr: true},
		"negative": {n: -8, wantLen: 0, wantErr: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := gen.SecureBytes(tc.n)
			if (err != nil) != tc.wantErr {
				t.Errorf("SecureBytes() error = %v, wantErr %v", err, tc.wantErr)
			}
			if tc.wantErr {
				if !errors.Is(err, gen.ErrNonPositiveLength) {
					t.Errorf("SecureBytes() error = %v, want ErrNonPositiveLength", err)
				}
				if got != nil {
					t.Errorf("SecureBytes() returned %d bytes alongside an error, want nil", len(got))
				}
				return
			}
			if len(got) != tc.wantLen {
				t.Errorf("SecureBytes() length = %d, want %d", len(got), tc.wantLen)
			}
		})
	}
}

// Repeated calls must not collide.
func TestSecureBytesAreDistinct(t *testing.T) {
	t.Parallel()
	const (
		iterations = 512
		size       = 16
	)

	seen := make(map[string]struct{}, iterations)
	for i := range iterations {
		b, err := gen.SecureBytes(size)
		if err != nil {
			t.Fatalf("SecureBytes() call %d: %v", i, err)
		}
		key := string(b)
		if _, dup := seen[key]; dup {
			t.Fatalf("SecureBytes() returned a duplicate value on call %d", i)
		}
		seen[key] = struct{}{}
	}
}

func TestSecureBytesAreNotAllZero(t *testing.T) {
	t.Parallel()
	b, err := gen.SecureBytes(32)
	if err != nil {
		t.Fatalf("SecureBytes() error: %v", err)
	}

	for _, v := range b {
		if v != 0 {
			return
		}
	}
	t.Error("SecureBytes() returned all zero bytes")
}

func TestSecureToken(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		n       int
		wantLen int // unpadded base64url length for n bytes
		wantErr bool
	}{
		"sixteen bytes":  {n: 16, wantLen: 22, wantErr: false},
		"thirtytwo byte": {n: 32, wantLen: 43, wantErr: false},
		"zero":           {n: 0, wantLen: 0, wantErr: true},
		"negative":       {n: -1, wantLen: 0, wantErr: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := gen.SecureToken(tc.n)
			if (err != nil) != tc.wantErr {
				t.Errorf("SecureToken() error = %v, wantErr %v", err, tc.wantErr)
			}
			if tc.wantErr {
				if got != "" {
					t.Errorf("SecureToken() = %q alongside an error, want empty", got)
				}
				return
			}
			if len(got) != tc.wantLen {
				t.Errorf("SecureToken() length = %d, want %d", len(got), tc.wantLen)
			}

			decoded, derr := base64.RawURLEncoding.DecodeString(got)
			if derr != nil {
				t.Errorf("SecureToken() = %q, which is not unpadded base64url: %v", got, derr)
			}
			if len(decoded) != tc.n {
				t.Errorf("SecureToken() decoded to %d bytes, want %d", len(decoded), tc.n)
			}
		})
	}
}

// Token must need no escaping in a URL, form field, header or filename.
func TestSecureTokenIsTransportSafe(t *testing.T) {
	t.Parallel()
	for range 128 {
		token, err := gen.SecureToken(32)
		if err != nil {
			t.Fatalf("SecureToken() error: %v", err)
		}
		if strings.ContainsAny(token, "+/= \t\r\n?&#%\\:") {
			t.Fatalf("SecureToken() = %q contains a character needing escaping", token)
		}
	}
}

func TestSecureHex(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		n       int
		wantLen int // hex is always two characters per byte
		wantErr bool
	}{
		"eight":    {n: 8, wantLen: 16, wantErr: false},
		"sha256":   {n: 32, wantLen: 64, wantErr: false},
		"zero":     {n: 0, wantLen: 0, wantErr: true},
		"negative": {n: -4, wantLen: 0, wantErr: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := gen.SecureHex(tc.n)
			if (err != nil) != tc.wantErr {
				t.Errorf("SecureHex() error = %v, wantErr %v", err, tc.wantErr)
			}
			if tc.wantErr {
				if got != "" {
					t.Errorf("SecureHex() = %q alongside an error, want empty", got)
				}
				return
			}
			if len(got) != tc.wantLen {
				t.Errorf("SecureHex() length = %d, want %d", len(got), tc.wantLen)
			}
			if got != strings.ToLower(got) {
				t.Errorf("SecureHex() = %q, want lowercase", got)
			}
			if _, derr := hex.DecodeString(got); derr != nil {
				t.Errorf("SecureHex() = %q, which is not valid hex: %v", got, derr)
			}
		})
	}
}

func BenchmarkSecureBytes(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		if _, err := gen.SecureBytes(32); err != nil {
			b.Fatalf("SecureBytes() error: %v", err)
		}
	}
}

func BenchmarkSecureToken(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		if _, err := gen.SecureToken(32); err != nil {
			b.Fatalf("SecureToken() error: %v", err)
		}
	}
}
