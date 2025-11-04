// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utility

type ResourceContext struct {
	ResourceDirectory      string
	E2eTestDirectory       *string
	ResourceTypeName       string
	ResourceTypeNameForE2e string
	E2eRandSuffix          string
	StackName              string
}

func InitResourceCtx(stackName, randSuffix, resourceTypeName, resourceDirectory string, e2eTestDirectory *string) ResourceContext {
	return ResourceContext{
		StackName:              stackName,
		ResourceDirectory:      resourceDirectory,
		E2eTestDirectory:       e2eTestDirectory,
		ResourceTypeName:       resourceTypeName,
		E2eRandSuffix:          randSuffix,
		ResourceTypeNameForE2e: resourceTypeName + randSuffix,
	}
}
