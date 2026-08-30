/*
 *           gosynutils
 *     Copyright (c) Synertry 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

// Package fspath checks path existence: Check and CheckDir.
package fspath

import (
	"errors"
	"os"
)

// Check reports whether the file or directory at path exists. A
// genuinely missing path (os.ErrNotExist) reports pathExists=false with
// err=nil. any other stat failure (permission denied, and so on) is
// returned as err so the caller can tell "does not exist" apart from
// "could not tell".
func Check(path string) (pathExists bool, err error) {
	if _, statErr := os.Stat(path); statErr == nil { // exists
		pathExists = true
	} else if errors.Is(statErr, os.ErrNotExist) {
		pathExists = false // does not exist, not an error
	} else { // possible permission issue
		pathExists = false
		err = statErr
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
