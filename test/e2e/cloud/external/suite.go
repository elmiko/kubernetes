/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package external

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// concreteProvider holds the implementation of the provider interface.
var concreteProvider *ExternalCloudProvider

// InitializeAndRunExternalCloudProviderTests is the entry point for running all of the
// external cloud provider tests.
func InitializeAndRunExternalCloudProviderTests(t *testing.T, p ExternalCloudProvider) {
	RegisterFailHandler(Fail)
	setConcreteProvider(&p)
	RunSpecs(t, "External Cloud Provider Suite")
}

func setConcreteProvider(p *ExternalCloudProvider) {
	concreteProvider = p
}

func getConcreteProvider() *ExternalCloudProvider {
	if concreteProvider == nil {
		Fail("concreteProvider is not set, this condition should not occur under normal operation")
	}
	return concreteProvider
}
