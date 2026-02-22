/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package file_test

import (
	"errors"
	"log/slog"
	"os"
	"path"
	"testing"

	"github.com/Synertry/gosynutils/file"
)

//nolint:gocognit
func TestCopy(t *testing.T) {
	tests := map[string]struct {
		src, dst string
		success  bool
		err      error
	}{
		"simple_copy": {
			src:     "aTob.txt",
			dst:     "b.txt",
			success: true,
		},
		"empty_source": {
			src:     "", // will skip creation
			dst:     "bEmptySource.txt",
			success: false,
		},
		"empty_destination": {
			src:     "aEmptyDest.txt",
			dst:     "",
			success: false,
		},
		"invalid_source": {
			src:     "invalid",
			dst:     "bInvalidSource.txt",
			success: false,
		},
		"same_file": {
			src:     "sameFile.txt",
			dst:     "sameFile.txt",
			success: true,
		},
	}

	dirTest := path.Join(os.TempDir(), "TestCopy")

	// Ensure dirTest exists
	err := os.MkdirAll(dirTest, os.ModeDir|os.ModePerm)
	if err != nil {
		slog.Error("failed to create test directory:", "err", err)
		return
	}

	for name, tc := range tests {
		var pErr error
		//var fSrc *os.File
		strInputPath := path.Join(dirTest, tc.src)
		strOutputPath := path.Join(dirTest, tc.dst)

		// we first prepare a test file in the temp folder and then try to copy
		if tc.src != "" && tc.src != "invalid" {
			// fSrc, err = os.CreateTemp("", tc.src) // if this is skipped, nil gets passed to file.Copy, which panics
			// so this is better, since we just pass an empty string with that
			pErr = file.TouchFile(strInputPath)
		}
		if !errors.Is(pErr, nil) {
			t.Logf("failed to create source file: %v", pErr)
			continue
		}

		// Main tests
		t.Run(name, func(t *testing.T) {
			tErr := file.Copy(strInputPath, strOutputPath)
			if errors.Is(tErr, nil) != tc.success {
				t.Errorf("expected %t, got: %v\n", tc.success, tErr == nil)
				t.Logf("input file: %s", strInputPath)
				return
			}
		})

		// clean up only if a path is not empty
		if tc.src != "" && tc.src != "invalid" {
			pErr = os.Remove(strInputPath)
			if pErr != nil {
				slog.Info("failed to remove source file", "err", pErr)
				continue
			}
		}

		if tc.dst != "" && tc.success && tc.src != tc.dst {
			pErr = os.Remove(strOutputPath)
			if pErr != nil {
				slog.Info("failed to remove destination file", "err", pErr)
				continue
			}
		}
	}

	// clean up test directory
	err = os.RemoveAll(dirTest)
	if err != nil {
		slog.Info("failed to remove test directory", "err", err)
	}
}
