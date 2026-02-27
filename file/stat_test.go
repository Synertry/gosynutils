/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package file_test

import (
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/Synertry/gosynutils/file"
)

func TestGetStat(t *testing.T) { //nolint:gocognit
	tests := map[string]struct {
		setup     func(dir string) (string, error)
		wantSize  int64
		wantCount int
		wantDirs  int
		wantFiles int
		wantErr   bool
	}{
		"single_file": {
			setup: func(dir string) (string, error) {
				path := filepath.Join(dir, "single.txt")
				return path, os.WriteFile(path, []byte("hello"), 0644)
			},
			wantSize:  5,
			wantCount: 1,
			wantDirs:  0,
			wantFiles: 1,
		},
		"empty_directory": {
			setup: func(dir string) (string, error) {
				path := filepath.Join(dir, "empty")
				return path, os.Mkdir(path, 0755)
			},
			wantSize:  0,
			wantCount: 1,
			wantDirs:  1,
			wantFiles: 0,
		},
		"nested_structure": {
			setup: func(dir string) (path string, err error) {
				path = filepath.Join(dir, "nested")

				if err = os.Mkdir(path, 0755); err != nil {
					return
				}
				if err = os.WriteFile(filepath.Join(path, "f1.txt"), []byte("abc"), 0644); err != nil {
					return
				}

				subpath := filepath.Join(path, "inner")
				if err = os.Mkdir(subpath, 0755); err != nil {
					return
				}
				if err = os.WriteFile(filepath.Join(subpath, "f2.txt"), []byte("defg"), 0644); err != nil {
					return
				}

				return
			},
			wantSize:  7,
			wantCount: 4, // nested/, f1.txt, inner/, f2.txt
			wantDirs:  2, // nested/, inner/
			wantFiles: 2, // f1.txt, f2.txt
		},
		"nonexistent_path": {
			setup: func(dir string) (string, error) {
				return filepath.Join(dir, "nonexistent"), nil
			},
			wantErr: true,
		},
	}

	dirTest := path.Join(t.TempDir(), "TestGetStats")
	if err := os.MkdirAll(dirTest, 0755); err != nil {
		t.Fatalf("failed to create test directory: %v", err)
	}

	for name, tc := range tests {
		var path string
		var err error

		if path, err = tc.setup(dirTest); err != nil {
			t.Errorf("setup failed: %v", err)
			continue
		}

		t.Run(name, func(t *testing.T) {
			// Test GetSize
			gotSize, sErr := file.GetSize(path)
			if (sErr != nil) != tc.wantErr {
				t.Errorf("GetSize() error = %v, wantErr %v", sErr, tc.wantErr)
			}
			if !tc.wantErr && gotSize != tc.wantSize {
				t.Errorf("GetSize() = %v, want %v", gotSize, tc.wantSize)
			}

			// Test GetCount
			gotCount, cErr := file.GetCount(path)
			if (cErr != nil) != tc.wantErr {
				t.Errorf("GetCount() error = %v, wantErr %v", cErr, tc.wantErr)
			}
			if !tc.wantErr && gotCount != tc.wantCount {
				t.Errorf("GetCount() = %v, want %v", gotCount, tc.wantCount)
			}

			// Test GetCountDirs
			gotDirs, cdErr := file.GetCountDirs(path)
			if (cdErr != nil) != tc.wantErr {
				t.Errorf("GetCountDirs() error = %v, wantErr %v", cdErr, tc.wantErr)
			}
			if !tc.wantErr && gotDirs != tc.wantDirs {
				t.Errorf("GetCountDirs() = %v, want %v", gotDirs, tc.wantDirs)
			}

			// Test GetCountFiles
			gotFiles, cfErr := file.GetCountFiles(path)
			if (cfErr != nil) != tc.wantErr {
				t.Errorf("GetCountFiles() error = %v, wantErr %v", cfErr, tc.wantErr)
			}
			if !tc.wantErr && gotFiles != tc.wantFiles {
				t.Errorf("GetCountFiles() = %v, want %v", gotFiles, tc.wantFiles)
			}
		})
	}
}
