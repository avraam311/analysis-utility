package domain

type Analysis struct {
	AllocsNum  string `json:"allocs_num" binding:"required"`
	GCNum      string `json:"gc_num" binding:"required"`
	UsedMemory string `json:"used_memory" binding:"required"`
	GCLastTime string `json:"gc_last_time" binding:"required"`
}
