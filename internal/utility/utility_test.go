// Copyright 2025 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test utility package
package utility

import "testing"

func TestGetHome(t *testing.T) {
	_, err := GetHome()
	if err != nil {
		t.Errorf("GetHome() error %v", err)
	}
}
