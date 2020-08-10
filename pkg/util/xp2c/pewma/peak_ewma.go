package pewma

import (
	"math"
	"sync/atomic"
	"time"
)

type PeakEwma struct {
	stamp int64
	value int64
	tau   time.Duration
}

func New() *PeakEwma {
	return &PeakEwma{
		tau: 10000 * time.Millisecond,
	}
}

// Observe 计算peak指数加权移动平均值
func (p *PeakEwma) Observe(now, rtt int64) {

	stamp := atomic.SwapInt64(&p.stamp, now)
	td := now - stamp
	if td < 0 {
		td = 0
	}

	w := math.Exp(float64(-td) / float64(p.tau))
	latency := atomic.LoadInt64(&p.value)
	if rtt > latency {
		atomic.StoreInt64(&p.value, rtt)
	} else {
		atomic.StoreInt64(&p.value, int64(float64(latency)*w+float64(rtt)*(1.0-w)))
	}
}

func (p *PeakEwma) Value() int64 {
	return atomic.LoadInt64(&p.value)
}
