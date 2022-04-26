package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEagerSingleton(t *testing.T) {
	assert.Equal(t, EagerSingleton(), EagerSingleton())
}

func BenchmarkEagerSingletonParalell(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			singleton := LazySingleton()
			if EagerSingleton() != singleton {
				b.Errorf("test failed")
			}
		}
	})
}

func TestLazySingleton(t *testing.T) {
	assert.Equal(t, LazySingleton(), LazySingleton())
}

func BenchmarkLazySingletonParalell(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			singleton := LazySingleton()
			if LazySingleton() != singleton {
				b.Errorf("test failed")
			}
		}
	})
}
