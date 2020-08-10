package pewma

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPeakEwma(t *testing.T) {
	p := New()
	// p.tau = 600 * time.Millisecond
	p.Observe(time.Now().UnixNano(), int64(1*time.Second))
	assert.Equal(t, p.Value(), int64(1*time.Second))
	p.Observe(time.Now().UnixNano(), int64(1*time.Second))
	assert.Equal(t, p.Value(), int64(1*time.Second))
	p.Observe(time.Now().UnixNano(), int64(1*time.Second))
	assert.Equal(t, p.Value(), int64(1*time.Second))
	time.Sleep(1 * time.Second)
	p.Observe(time.Now().UnixNano(), int64(1*time.Second))
	assert.Equal(t, p.Value(), int64(1*time.Second))
	p.Observe(time.Now().UnixNano(), int64(2*time.Second))
	assert.Equal(t, p.Value(), int64(2*time.Second))

	for i := 0; i <= 1000; i++ {
		time.Sleep(1 * time.Microsecond)
		p.Observe(time.Now().UnixNano(), int64(1*time.Second))
	}
	assert.True(t, p.Value() > int64(1800*time.Millisecond) && p.Value() < int64(2000*time.Millisecond), fmt.Sprintf("%d", p.Value()))
}

func BenchmarkPeakEwma(b *testing.B) {
	b.ResetTimer()
	p := New()
	for i := 0; i < b.N; i++ {
		p.Observe(time.Now().UnixNano(), int64(time.Duration(rand.Intn(10))*time.Second))
	}
	// b.Error(p.Value())
}
