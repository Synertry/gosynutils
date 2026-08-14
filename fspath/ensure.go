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
	"fmt"
	"os"
)

// ErrNotADirectory is returned when a path exists but is a file, so it cannot
// serve as a directory. This is reported rather than ignored because the caller
// is about to write into that location and would otherwise fail later with a
// far less obvious error.
var ErrNotADirectory = errors.New("fspath: path exists but is not a directory")

// ErrEmptyPath is returned when an empty path is supplied. An empty path would
// otherwise resolve to the working directory, which is almost never intended.
var ErrEmptyPath = errors.New("fspath: path must not be empty")

// EnsureDir makes sure a directory exists at path, creating it and any missing
// parents if needed. It complements [Check] and [CheckDir], which only report.
//
// The created return value reports whether anything was actually made, so a
// caller can log or audit first-time creation without racing a second [Check].
//
// perm applies only to directories this call creates, and is subject to the
// process umask on Unix. On Windows the mode is largely ignored, so do not rely
// on it as an access control there. 0o750 is a reasonable default when the
// contents are not meant to be world readable.
//
// It returns [ErrEmptyPath] for an empty path and [ErrNotADirectory] if
// something exists at path that is not a directory.
func EnsureDir(path string, perm os.FileMode) (created bool, err error) {
	if path == "" {
		return false, ErrEmptyPath
	}

	var isDir bool
	if isDir, err = CheckDir(path); isDir {
		return false, nil // already a directory, nothing to do
	}

	// CheckDir reports an error for a path that does not exist yet, which is
	// the ordinary case here, so distinguish that from a real failure.
	var exists bool
	if exists, err = Check(path); err != nil && exists {
		return false, fmt.Errorf("fspath: inspect %q: %w", path, err)
	}
	if exists {
		return false, fmt.Errorf("%w: %q", ErrNotADirectory, path)
	}

	if err = os.MkdirAll(path, perm); err != nil {
		return false, fmt.Errorf("fspath: create %q: %w", path, err)
	}
	return true, nil
}
