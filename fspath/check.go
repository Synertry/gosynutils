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

// Check reports whether the file or directory at path exists.
// The error mirrors the underlying [os.Stat] call whenever pathExists is
// false: a missing path returns an error wrapping [os.ErrNotExist], any
// other failure such as permission denied returns that error unchanged.
// Callers separate the two with [errors.Is] against [os.ErrNotExist].
// CheckDir applies this same error contract against its own [os.Stat]
// call rather than calling Check, keep the two consistent.
func Check(path string) (pathExists bool, err error) {
	if _, err = os.Stat(path); err == nil { // exists
		pathExists = true
	} else if errors.Is(err, os.ErrNotExist) { //nolint:gocritic // we need both branches to display that there would be a logic difference
		pathExists = false // does not exist
	} else { // possible permission issue
		pathExists = false
		// Schrödinger: file may or may not exist. See err for details.
		// SOURCE: https://stackoverflow.com/a/12518877/5516320
	}
	return
}

// CheckDir reports whether path exists and leads to a directory.
// A directory returns (true, nil), an existing non-directory returns
// (false, nil), a missing path returns false with an error wrapping
// [os.ErrNotExist] and any other stat failure returns false with that
// error.
func CheckDir(path string) (isDir bool, err error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}
