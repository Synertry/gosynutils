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
	"path"
	"testing"

	"github.com/Synertry/gosynutils/file"
)

func TestMove(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		src, dst string
		success  bool
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

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Each subtest gets its own directory so parallel cases can never
			// collide on a path or race on another case's cleanup.
			dirTest := t.TempDir()
			strInputPath := path.Join(dirTest, tc.src)
			strOutputPath := path.Join(dirTest, tc.dst)

			if tc.src != "" && tc.src != "invalid" {
				if err := file.TouchFile(strInputPath); err != nil {
					t.Fatalf("failed to create source file: %v", err)
				}
			}

			tErr := file.Move(strInputPath, strOutputPath)
			if errors.Is(tErr, nil) != tc.success {
				t.Errorf("expected %t, got: %v\n", tc.success, tErr == nil)
				t.Logf("input file: %s", strInputPath)
			}
		})
	}
}
