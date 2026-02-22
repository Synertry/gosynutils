/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package file_test //nolint:cyclop // This unfortunately needs this level of complexity to test the stat functions

import (
	"errors"
	"os"
	"path"
	"testing"

	"github.com/Synertry/gosynutils/file"
)

// TestTouchFile tests the TouchFile function that it ensures a file exists.
// So it is like the Unix touch command.
func TestTouchFile(t *testing.T) {
	var pathTempFile = path.Join(os.TempDir(), "tempfile_for_touchfile_test.tmp")

	// Ensure the file does not exist before the test
	_ = os.Remove(pathTempFile)

	t.Run("TouchFile", func(t *testing.T) {
		err := file.TouchFile(pathTempFile)
		if err != nil {
			t.Errorf("TouchFile failed: %v", err)
			return
		}

		// Check if the file now exists
		if _, err := os.Stat(pathTempFile); errors.Is(err, os.ErrNotExist) {
			t.Errorf("File %s does not exist after TouchFile", pathTempFile)
			return
		}
	})

	// Clean up
	_ = os.Remove(pathTempFile)
}
