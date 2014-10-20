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

// TODO: moredummier requires more diverse method names.

type moredummier interface {
	Dummy1()
	Dummy2()
	Dummy3()
	Dummy4()
	Dummy5()
	Dummy6()
	Dummy7()
	Dummy8()
	Dummy9()
	Dummy10()
	Dummy11()
	Dummy12()
	Dummy13()
	Dummy14()
	Dummy15()
	Dummy16()
}

type concreteMore struct {
	name string
}

func (*concreteMore) Dummy1() {
	// noop
}

func (*concreteMore) Dummy2() {
	// noop
}

func (*concreteMore) Dummy3() {
	// noop
}

func (*concreteMore) Dummy4() {
	// noop
}

func (*concreteMore) Dummy5() {
	// noop
}

func (*concreteMore) Dummy6() {
	// noop
}

func (*concreteMore) Dummy7() {
	// noop
}

func (*concreteMore) Dummy8() {
	// noop
}

func (*concreteMore) Dummy9() {
	// noop
}

func (*concreteMore) Dummy10() {
	// noop
}

func (*concreteMore) Dummy11() {
	// noop
}

func (*concreteMore) Dummy12() {
	// noop
}

func (*concreteMore) Dummy13() {
	// noop
}

func (*concreteMore) Dummy14() {
	// noop
}

func (*concreteMore) Dummy15() {
	// noop
}

func (*concreteMore) Dummy16() {
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

func BenchmarkIValueConcreteMoreAssgn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var d moredummier = &concreteMore{}
		_ = d
	}
}
