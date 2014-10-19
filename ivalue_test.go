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

type concreteInterface struct {
	concrete
	value interface{}
}

func BenchmarkConcreteAssignment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &concrete{}
	}
}

func BenchmarkIValueAssignment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var d dummier = &concrete{}
		_ = d
	}
}

func BenchmarkIValueWithInterfaceFieldAssignment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var d dummier = &concreteInterface{}
		_ = d
	}
}
