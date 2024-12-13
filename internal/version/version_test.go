// Copyright 2024 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test version package
package version

import (
	"os"
	"testing"
)

func TestGetVersion(t *testing.T) {
	filename := "/test/test.csv"

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = GetVersion(path + "/../.." + filename)
	if err != nil {
		t.Errorf("GetVersion() error %v", err)
	}
}
