// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package benchmarks

import "testing"

type dummier interface {
	Dummy()
}

type concrete struct {
	name string
}

func (*concrete) Dummy() {
	// noop
}

type concreteTwoFields struct {
	name  string
	value string
}

func (*concreteTwoFields) Dummy() {
	// noop
}

type concreteInterfaceField struct {
	name  string
	value interface{}
}

func (*concreteInterfaceField) Dummy() {
	// noop
}

func BenchmarkIValueConcreteAssgn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &concrete{}
	}
}

func BenchmarkIValueInterfaceAssgn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var d dummier = &concrete{}
		_ = d
	}
}

func BenchmarkIValueEmptyInterfaceAssgn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var d interface{} = &concrete{}
		_ = d
	}
}

func BenchmarkIValueInterfaceTwoFieldsAssgn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var d dummier = &concreteTwoFields{}
		_ = d
	}
}

func BenchmarkIValueInterfaceInterfaceFieldsAssgn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var d dummier = &concreteInterfaceField{}
		_ = d
	}
}
