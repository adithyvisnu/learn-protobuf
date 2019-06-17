package model

import (
	"testing"
)

func BenchmarkNameResJSON(b *testing.B) {
	name := NameResponse{
		NameRes: "Adithya Visnu",
	}

	for i := 0; i < b.N; i++ {
		JSMarshal(&name)
	}
}
func BenchmarkNameResPB(b *testing.B) {
	name := NameResponse{
		NameRes: "Adithya Visnu",
	}

	for i := 0; i < b.N; i++ {
		PBMarshal(&name)
	}
}
