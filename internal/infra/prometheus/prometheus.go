package prometheus

import (
	"runtime"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	allocsMetric = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "go_memstats_allocs_total",
		Help: "Total number of allocations",
	}, func() float64 {
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		return float64(mem.Mallocs)
	})

	gcCountMetric = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "go_memstats_gc_runs_total",
		Help: "Number of completed GC cycles",
	}, func() float64 {
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		return float64(mem.NumGC)
	})

	usedMemMetric = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "go_memstats_alloc_bytes",
		Help: "Currently used heap bytes",
	}, func() float64 {
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		return float64(mem.Alloc)
	})

	lastGCTimeMetric = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "go_memstats_last_gc_time_seconds",
		Help: "Time of last garbage collection in seconds since epoch",
	}, func() float64 {
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		if mem.NumGC == 0 {
			return 0
		}
		lastPauseEnd := mem.PauseEnd[(mem.NumGC+255)%256]
		return float64(lastPauseEnd) / 1e9
	})
)

func Init() {
	prometheus.MustRegister(
		allocsMetric,
		gcCountMetric,
		usedMemMetric,
		lastGCTimeMetric,
	)
}
