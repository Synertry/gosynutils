/*
 *             gosynutils
 *     Copyright (c) gosynutils 2026.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

package jsonx

import (
	"encoding/json"
)

func PrettyPrint(v any) string {
	s, _ := json.MarshalIndent(v, "", "\t")
	//fmt.Println(string(s))
	return string(s)
}
