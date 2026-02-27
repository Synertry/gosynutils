/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

// http://golang.org/src/os/stat_windows.go

package atime

import (
	"os"
	"syscall"
	"time"
)

func atime(fi os.FileInfo) time.Time {
	return time.Unix(
		0,
		fi.Sys().(*syscall.Win32FileAttributeData).LastAccessTime.Nanoseconds(), //nolint:errcheck // original code
	)
}
