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

	// Check tells missing (false, nil) apart from a real stat failure
	// (false, err) and from an existing path (true, nil). CheckDir alone
	// can't: it collapses "missing" and "exists as a file" into the same
	// (false, nil) result, so it can't drive this decision by itself.
	exists, err := Check(path)
	switch {
	case err != nil:
		return false, fmt.Errorf("fspath: inspect %q: %w", path, err)
	case !exists:
		if err = os.MkdirAll(path, perm); err != nil {
			return false, fmt.Errorf("fspath: create %q: %w", path, err)
		}
		return true, nil
	}

	info, err := os.Stat(path)
	if err != nil {
		return false, fmt.Errorf("fspath: inspect %q: %w", path, err)
	}
	if !info.IsDir() {
		return false, fmt.Errorf("%w: %q", ErrNotADirectory, path)
	}
	return false, nil
}
