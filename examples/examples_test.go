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

package examples

import (
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/testing/integration"
)

func TestAccRealmTs(t *testing.T) {
	test := getJSBaseOptions().
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "realm", "ts"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccRealmPython(t *testing.T) {
	test := getPythonBaseOptions().
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "realm", "python"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccRealmCsharp(t *testing.T) {
	test := getCSBaseOptions().
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "realm", "csharp"),
		})

	integration.ProgramTest(t, &test)
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		ExpectRefreshChanges: true,
	}
}

func getJSBaseOptions() integration.ProgramTestOptions {
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/keycloak",
		},
	})

	return baseJS
}

func getCSBaseOptions() integration.ProgramTestOptions {
	base := getBaseOptions()
	csharpBase := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"Pulumi.Keycloak",
		},
	})

	return csharpBase
}

func getPythonBaseOptions() integration.ProgramTestOptions {
	base := getBaseOptions()
	basePython := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			filepath.Join("..", "sdk", "python", "bin"),
		},
	})

	return basePython
}
