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

func coreTestCheck(t *testing.T) (dir string, pathfile string, non string) {
	var err error

	base := t.TempDir()

	pathfile = filepath.Join(base, "file.txt")
	if err = file.TouchFile(pathfile); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	dir = filepath.Join(base, "subdir")
	if err = os.Mkdir(dir, 0755); err != nil {
		t.Fatalf("failed to create test dir: %v", err)
	}

	non = filepath.Join(base, "nonexistent")
	return
}

func TestCheck(t *testing.T) {
	t.Parallel()
	dir, pathFile, nonExistent := coreTestCheck(t)
	if nonExistent == "" {
		t.Fatal("failed to create test files")
	}

	tests := map[string]struct {
		path       string
		wantExists bool
		wantErr    bool
		wantErrIs  error
	}{
		"dir": {
			path:       dir,
			wantExists: true,
			wantErr:    false,
		},
		"file": {
			path:       pathFile,
			wantExists: true,
			wantErr:    false,
		},
		"nonexistent": {
			path:       nonExistent,
			wantExists: false,
			wantErr:    true,
			wantErrIs:  os.ErrNotExist,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			gotExists, err := fspath.Check(tc.path)
			if (err != nil) != tc.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tc.wantErr)
			}
			if tc.wantErrIs != nil && !errors.Is(err, tc.wantErrIs) {
				t.Errorf("Check() error = %v, want errors.Is %v", err, tc.wantErrIs)
			}
			if gotExists != tc.wantExists {
				t.Errorf("Check() gotExists = %v, want %v", gotExists, tc.wantExists)
			}
		})
	}
}

func TestCheckDir(t *testing.T) {
	t.Parallel()
	dir, pathFile, nonExistent := coreTestCheck(t)
	if nonExistent == "" {
		t.Fatal("failed to create test files")
	}

	tests := map[string]struct {
		path      string
		wantDir   bool
		wantErr   bool
		wantErrIs error
	}{
		"dir": {
			path:    dir,
			wantDir: true,
			wantErr: false,
		},
		"file": {
			path:    pathFile,
			wantDir: false,
			wantErr: false,
		},
		"nonexistent": {
			path:      nonExistent,
			wantDir:   false,
			wantErr:   true,
			wantErrIs: os.ErrNotExist,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			gotDir, err := fspath.CheckDir(tc.path)
			if (err != nil) != tc.wantErr {
				t.Errorf("CheckDir() error = %v, wantErr %v", err, tc.wantErr)
			}
			if tc.wantErrIs != nil && !errors.Is(err, tc.wantErrIs) {
				t.Errorf("CheckDir() error = %v, want errors.Is %v", err, tc.wantErrIs)
			}
			if gotDir != tc.wantDir {
				t.Errorf("CheckDir() gotDir = %v, want %v", gotDir, tc.wantDir)
			}
		})
	}
}
