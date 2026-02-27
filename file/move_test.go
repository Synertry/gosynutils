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
	"os"
	"path"
	"testing"

	"github.com/Synertry/gosynutils/file"
)

func TestMove(t *testing.T) { //nolint:gocognit
	tests := map[string]struct {
		src, dst string
		success  bool
		err      error
	}{
		"simple_move": {
			src:     "a.txt",
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

	dirTest := path.Join(t.TempDir(), "MoveCopy")

	// Ensure dirTest exists
	err := os.MkdirAll(dirTest, os.ModeDir|os.ModePerm)
	if err != nil {
		t.Fatalf("failed to create test directory: %v", err)
	}

	for name, tc := range tests {
		var pErr error
		// var fSrc *os.File
		strInputPath := path.Join(dirTest, tc.src)
		strOutputPath := path.Join(dirTest, tc.dst)

		// we first prepare a test file in the temp folder and then try to move
		if tc.src != "" && tc.src != "invalid" {
			// fSrc, err = os.CreateTemp("", tc.src) // if this is skipped, nil gets passed to file.Move, which panics
			// so this is better, since we just pass an empty string with that
			pErr = file.TouchFile(strInputPath)
		}
		if !errors.Is(pErr, nil) {
			t.Logf("failed to create source file: %v", pErr)
			continue
		}

		// Main tests
		t.Run(name, func(t *testing.T) {
			tErr := file.Move(strInputPath, strOutputPath)
			if errors.Is(tErr, nil) != tc.success {
				t.Errorf("expected %t, got: %v\n", tc.success, tErr == nil)
				t.Logf("input file: %s", strInputPath)
				return
			}
		})

		// clean up

		// only refers to test case "empty destination", because that move operation would fail
		// which leaves behind the source file
		if tc.dst != "" && !tc.success {
			pErr = os.Remove(strInputPath)
			if pErr != nil {
				t.Logf("failed to remove source file: %v", pErr)
				continue
			}
		}

		if tc.success {
			pErr = os.Remove(strOutputPath)
			if pErr != nil {
				t.Logf("failed to remove destination file: %v", pErr)
				continue
			}
		}
	}
}
