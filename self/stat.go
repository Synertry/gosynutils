/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

// Package self provides the path to the executable and its directory
package self

import (
	"log/slog"
	"os"
	"path/filepath"
)

// GetPathExe returns the path to the executable file
func GetPathExe() string {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	pathExe, err := os.Executable()
	if err != nil {
		logger.Error("failed to get executable path", "error", err)
		return ""
	}

	return pathExe
}

// GetPathExeDir returns the directory containing the executable file
// it depends on GetPathExe
func GetPathExeDir() string {
	return filepath.Dir(GetPathExe())
}
