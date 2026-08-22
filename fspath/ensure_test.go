/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package fspath_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/Synertry/gosynutils/file"
	"github.com/Synertry/gosynutils/fspath"
)

func TestEnsureDir(t *testing.T) {
	t.Parallel()
	base := t.TempDir()

	existingDir := filepath.Join(base, "existing")
	if err := os.Mkdir(existingDir, 0o750); err != nil {
		t.Fatalf("failed to create test dir: %v", err)
	}

	existingFile := filepath.Join(base, "file.txt")
	if err := file.TouchFile(existingFile); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	tests := map[string]struct {
		path        string
		wantCreated bool
		wantErr     bool
		wantIsDir   bool
	}{
		"already a directory": {
			path:        existingDir,
			wantCreated: false,
			wantErr:     false,
			wantIsDir:   true,
		},
		"new directory": {
			path:        filepath.Join(base, "fresh"),
			wantCreated: true,
			wantErr:     false,
			wantIsDir:   true,
		},
		"missing parents": {
			path:        filepath.Join(base, "a", "b", "c"),
			wantCreated: true,
			wantErr:     false,
			wantIsDir:   true,
		},
		"path is a file": {
			path:        existingFile,
			wantCreated: false,
			wantErr:     true,
			wantIsDir:   false,
		},
		"empty path": {
			path:        "",
			wantCreated: false,
			wantErr:     true,
			wantIsDir:   false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			gotCreated, err := fspath.EnsureDir(tc.path, 0o750)
			if (err != nil) != tc.wantErr {
				t.Errorf("EnsureDir() error = %v, wantErr %v", err, tc.wantErr)
			}
			if gotCreated != tc.wantCreated {
				t.Errorf("EnsureDir() gotCreated = %v, want %v", gotCreated, tc.wantCreated)
			}
			if tc.wantErr {
				return
			}

			isDir, cerr := fspath.CheckDir(tc.path)
			if cerr != nil {
				t.Errorf("CheckDir() after EnsureDir(): %v", cerr)
			}
			if isDir != tc.wantIsDir {
				t.Errorf("after EnsureDir(), CheckDir() = %v, want %v", isDir, tc.wantIsDir)
			}
		})
	}
}

// Calling twice must report created only the first time.
func TestEnsureDirIsIdempotent(t *testing.T) {
	t.Parallel()
	path := filepath.Join(t.TempDir(), "repeat")

	created, err := fspath.EnsureDir(path, 0o750)
	if err != nil {
		t.Fatalf("EnsureDir() first call: %v", err)
	}
	if !created {
		t.Error("EnsureDir() first call reported created = false, want true")
	}

	created, err = fspath.EnsureDir(path, 0o750)
	if err != nil {
		t.Fatalf("EnsureDir() second call: %v", err)
	}
	if created {
		t.Error("EnsureDir() second call reported created = true, want false")
	}
}

func TestEnsureDirSentinelErrors(t *testing.T) {
	t.Parallel()
	base := t.TempDir()

	existingFile := filepath.Join(base, "blocker.txt")
	if err := file.TouchFile(existingFile); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	tests := map[string]struct {
		path    string
		wantErr error
	}{
		"empty path":   {path: "", wantErr: fspath.ErrEmptyPath},
		"path is file": {path: existingFile, wantErr: fspath.ErrNotADirectory},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, err := fspath.EnsureDir(tc.path, 0o750); !errors.Is(err, tc.wantErr) {
				t.Errorf("EnsureDir() error = %v, want %v", err, tc.wantErr)
			}
		})
	}
}

func BenchmarkEnsureDirExisting(b *testing.B) {
	path := b.TempDir()
	b.ReportAllocs()
	for b.Loop() {
		if _, err := fspath.EnsureDir(path, 0o750); err != nil {
			b.Fatalf("EnsureDir() error: %v", err)
		}
	}
}
