/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package file

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/Synertry/gosynutils/file/internal/atime"
)

// Copy copies a file from src to dst. If src and dst files exist and are the same, then return success.
// Otherwise, copy the file contents from src to dst.
func Copy(src, dst string) (err error) {
	var sfi, dfi os.FileInfo

	sfi, err = os.Stat(src)
	if err != nil {
		return
	}

	if !sfi.Mode().IsRegular() { // cannot copy non-regular files (e.g., directories, symlinks, devices, etc.)
		return fmt.Errorf("non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}

	dfi, err = os.Stat(dst)
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}

		if os.SameFile(sfi, dfi) {
			return
		}
	}

	/*if err = os.Link(src, dst); err == nil { // I see problems here with network paths
		return
	}*/

	err = copyContents(src, dst)
	if err != nil {
		return
	}

	err = os.Chtimes(dst, atime.Get(sfi), sfi.ModTime())
	if err != nil {
		return fmt.Errorf("setting preserved times failed: %w", err)
	}

	err = os.Chmod(dst, sfi.Mode())
	if err != nil {
		return fmt.Errorf("setting preserved modes failed: %w", err)
	}

	return
}

// copyContents is the core function of Copy to copy file contents
func copyContents(src, dst string) (err error) {
	var in, out *os.File

	in, err = os.Open(src)
	if err != nil {
		return
	}
	defer func() {
		errCi := in.Close()
		if err == nil {
			err = errCi
		}
	}()

	out, err = os.Create(dst)
	if err != nil {
		return
	}

	defer func() {
		errCo := out.Close()
		if err == nil {
			err = errCo
		}
	}()

	if _, err = io.Copy(out, in); err != nil { // core copy
		return fmt.Errorf("core function copy failed: %w", err)
	}

	return out.Sync() // flush
}
