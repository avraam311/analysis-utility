package domain

import "time"

type Analysis struct {
	AllocsNum  uint64      `json:"allocs_num" binding:"required"`
	GCNum      uint32      `json:"gc_num" binding:"required"`
	UsedMemory uint64      `json:"used_memory" binding:"required"`
	GCLastTime time.Time `json:"gc_last_time" binding:"required"`
}
