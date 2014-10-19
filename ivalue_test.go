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
