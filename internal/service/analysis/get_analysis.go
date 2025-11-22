package analysis

import (
	"context"
	"runtime"
	"time"

	"github.com/avraam311/analysis-utility/internal/models/domain"
)

func (s *Service) GetAnalysis(ctx context.Context) (*domain.Analysis, error) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	lastGCPauseNs := memStats.PauseEnd[(memStats.NumGC+255)%256]
	lastGCTime := time.Unix(0, int64(lastGCPauseNs))

	analysis := &domain.Analysis{
		AllocsNum:  memStats.Mallocs,
		GCNum:      memStats.NumGC,
		UsedMemory: memStats.Alloc,
		GCLastTime: lastGCTime,
	}
	return analysis, nil
}
