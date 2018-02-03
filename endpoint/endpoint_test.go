/*
Copyright 2017 The Kubernetes Authors.

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

package endpoint

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEndpoint(t *testing.T) {
	e := NewEndpoint("example.org", "foo.com", "CNAME")
	if e.DNSName != "example.org" || e.Target != "foo.com" || e.RecordType != "CNAME" {
		t.Error("endpoint is not initialized correctly")
	}
	if e.Labels == nil {
		t.Error("Labels is not initialized")
	}

	w := NewEndpoint("example.org.", "load-balancer.com.", "")
	if w.DNSName != "example.org" || w.Target != "load-balancer.com" || w.RecordType != "" {
		t.Error("endpoint is not initialized correctly")
	}
}

func TestNewEndpointFrom(t *testing.T) {
	e1 := NewEndpoint("example.org", "foo.com", "CNAME")
	e1.ProviderAnnotations = map[string]string{
		"test": "foo",
	}

	e2 := NewEndpointFrom(e1, "example.org", "foo.com", "TXT")
	assert.Equal(t, e1.ProviderAnnotations, e2.ProviderAnnotations)
}
