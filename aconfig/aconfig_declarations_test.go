// Copyright 2018 Google Inc. All rights reserved.
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

package aconfig

import (
	"strings"
	"testing"

	"android/soong/android"
)

func TestAconfigDeclarations(t *testing.T) {
	bp := `
		aconfig_declarations {
			name: "module_name",
			package: "com.example.package",
			srcs: [
				"foo.aconfig",
				"bar.aconfig",
			],
		}
	`
	result := runTest(t, android.FixtureExpectsNoErrors, bp)

	module := result.ModuleForTests("module_name", "").Module().(*DeclarationsModule)

	// Check that the provider has the right contents
	depData := result.ModuleProvider(module, declarationsProviderKey).(declarationsProviderData)
	android.AssertStringEquals(t, "package", depData.Package, "com.example.package")
	if !strings.HasSuffix(depData.IntermediatePath.String(), "/intermediate.pb") {
		t.Errorf("Missing intermediates path in provider: %s", depData.IntermediatePath.String())
	}
}
