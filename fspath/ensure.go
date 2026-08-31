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

// ErrNotADirectory is returned when path exists but is not a directory.
var ErrNotADirectory = errors.New("fspath: path exists but is not a directory")

// ErrEmptyPath is returned when path is empty.
var ErrEmptyPath = errors.New("fspath: path must not be empty")

// EnsureDir makes sure a directory exists at path, creating it and any
// missing parents if needed. created reports whether anything was actually
// made. perm applies only to directories this call creates and is subject to
// the umask on Unix; largely ignored on Windows.
//
// Returns [ErrEmptyPath] for an empty path and [ErrNotADirectory] if path
// exists but is not a directory.
func EnsureDir(path string, perm os.FileMode) (created bool, err error) {
	if path == "" {
		return false, ErrEmptyPath
	}

	// One CheckDir call classifies every case: (true, nil) directory in
	// place, (false, nil) something else occupies the path, an error
	// wrapping os.ErrNotExist means missing and any other error is a
	// real stat failure.
	isDir, err := CheckDir(path)
	switch {
	case err == nil:
		if !isDir {
			return false, fmt.Errorf("%w: %q", ErrNotADirectory, path)
		}
		return false, nil
	case !errors.Is(err, os.ErrNotExist):
		return false, fmt.Errorf("fspath: inspect %q: %w", path, err)
	}

	if err = os.MkdirAll(path, perm); err != nil {
		return false, fmt.Errorf("fspath: create %q: %w", path, err)
	}
	return true, nil
}
