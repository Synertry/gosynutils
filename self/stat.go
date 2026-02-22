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

var (
	// PathExe is the path to the executable file
	PathExe string
	// PathExeDir is the directory containing the executable file
	// it depends on PathExe
	PathExeDir string
)

func init() {
	var err error
	PathExe, err = os.Executable()
	if err != nil {
		slog.Error("failed to get executable path", "error", err)
		return
	}
	PathExeDir = filepath.Dir(PathExe)
}
