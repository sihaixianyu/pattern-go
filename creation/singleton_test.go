package creation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEagerInstance(t *testing.T) {
	assert.Equal(t, GetEagerInstance(), GetEagerInstance())
}

func BenchmarkGetEagerInstanceParalell(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetEagerInstance() != GetEagerInstance() {
				b.Errorf("test failed")
			}
		}
	})
}
