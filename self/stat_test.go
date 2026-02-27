/*
 *           gosynutils
 *     Copyright (c) Synertry 2025 - 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package self_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Synertry/gosynutils/self"
)

func TestStat(t *testing.T) {
	pathExe, err := os.Executable()
	if err != nil {
		t.Errorf("failed get executeable to setup test: %v", err)
		return
	}
	pathExeDir := filepath.Dir(pathExe)

	t.Run("PathExe", func(t *testing.T) {
		selfPathExe := self.GetPathExe()
		if pathExe != selfPathExe {
			t.Errorf("expected: %q, got: %q\n", pathExe, selfPathExe)
			return
		}
	})
	t.Run("PathExeDir", func(t *testing.T) {
		selfPathExeDir := self.GetPathExeDir()
		if pathExeDir != selfPathExeDir {
			t.Errorf("expected: %q, got: %q\n", pathExeDir, selfPathExeDir)
			return
		}
	})
}
