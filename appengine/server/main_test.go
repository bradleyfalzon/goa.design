// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"flag"
	"log"
	"os"
	"testing"

	"google.golang.org/appengine/aetest"
)

// global App Engine test instance, initialized and shutdown in TestMain.
var testInstance aetest.Instance

func TestMain(m *testing.M) {
	flag.Parse()

	var err error
	testInstance, err = aetest.NewInstance(nil)
	if err != nil {
		log.Fatal(err)
	}

	code := m.Run()
	testInstance.Close()
	os.Exit(code)
}
