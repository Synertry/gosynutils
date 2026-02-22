/*
 *           gosynutils
 *     Copyright (c) Synertry 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package atime_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/Synertry/gosynutils/file/internal/atime"
)

func TestStat(t *testing.T) {
	tmpDir := t.TempDir()
	existingFile := filepath.Join(tmpDir, "existing.txt")
	if err := os.WriteFile(existingFile, []byte("content"), 0644); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		path    string
		wantErr error
	}{
		{
			name:    "existing file",
			path:    existingFile,
			wantErr: nil,
		},
		{
			name:    "existing directory",
			path:    tmpDir,
			wantErr: nil,
		},
		{
			name:    "non-existent file",
			path:    filepath.Join(tmpDir, "missing.txt"),
			wantErr: os.ErrNotExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := atime.Stat(tt.path)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Stat() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("Stat() unexpected error = %v", err)
				return
			}
			if got.IsZero() {
				t.Error("Stat() returned zero time for existing path")
			}
		})
	}
}

func TestGet(t *testing.T) {
	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	if err := os.WriteFile(f, []byte("test"), 0644); err != nil {
		t.Fatal(err)
	}
	fi, err := os.Stat(f)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		fi   os.FileInfo
	}{
		{
			name: "valid FileInfo",
			fi:   fi,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := atime.Get(tt.fi)
			if got.IsZero() {
				t.Error("Get() returned zero time")
			}
		})
	}
}
