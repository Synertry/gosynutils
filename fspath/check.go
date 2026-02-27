/*
 *           gosynutils
 *     Copyright (c) Synertry 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package fspath

import (
	"errors"
	"os"
)

// Check checks existence of provided path, returns error if error is not *PathError
func Check(path string) (pathExits bool, err error) {
	if _, err = os.Stat(path); err == nil { // exists
		pathExits = true
	} else if errors.Is(err, os.ErrNotExist) { //nolint:gocritic // we need both branches to display that there would be a logic difference
		pathExits = false // does not exist
	} else { // possible permission issue
		pathExits = false
		// Schrödinger: file may or may not exist. See err for details.
		// SOURCE: https://stackoverflow.com/a/12518877/5516320
	}
	return
}

// CheckDir checks if path exists and leads to a directory
func CheckDir(path string) (isDir bool, err error) {
	var exists bool
	exists, err = Check(path)
	if !exists {
		return // error no need to check further
	}

	var info os.FileInfo
	if info, err = os.Stat(path); err == nil {
		if info.IsDir() {
			isDir = true
		}
	}
	return // streamlined return values
}
