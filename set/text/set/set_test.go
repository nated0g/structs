// Copyright 2022 Robert S. Muhlestein.
// SPDX-License-Identifier: Apache-2.0

package set_test

import (
	"fmt"

	"github.com/rwxrob/structs/set/text/set"
)

func ExampleMinus() {
	s := []string{
		"one", "two", "three", "four", "five", "six", "seven",
	}
	m := []string{"two", "four", "six"}
	fmt.Println(set.Minus[string, string](s, m))
	// Output:
	// [one three five seven]
}
