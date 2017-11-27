package stats

import (
	"math"
	"runtime"
	"time"

	"github.com/jkomyno/palitrux/config"
)

var start = time.Now()

// MB represents 1 << 20 bytes
const MB float64 = 1048576 // 1 << 20

// RuntimeStats describes the runtime statistics
type RuntimeStats struct {
	Uptime               int64   `json:"uptime"`               // Time since the server started
	AllocatedMemory      float64 `json:"allocatedMemory"`      // MetaBytes of allocated heap objects
	TotalAllocatedMemory float64 `json:"totalAllocatedMemory"` // MetaBytes of allocated heap objects since start
	Goroutines           int     `json:"goroutines"`           // Number of goroutines dispatched
	CPUs                 int     `json:"cpus"`                 // Number of CPUs in use
	Version              string  `json:"version"`              // Current version of the app
}

// GetUptime returns the uptime expressed in UNIX format
func GetUptime() int64 {
	return time.Now().Unix() - start.Unix()
}

// GetRuntimeStats returns some useful stats regarding the server's health
func GetRuntimeStats() *RuntimeStats {
	mem := &runtime.MemStats{}
	runtime.ReadMemStats(mem)

	return &RuntimeStats{
		Uptime:               GetUptime(),
		AllocatedMemory:      toMegaBytes(mem.Alloc),
		TotalAllocatedMemory: toMegaBytes(mem.TotalAlloc),
		Goroutines:           runtime.NumGoroutine(),
		CPUs:                 runtime.NumCPU(),
		Version:              config.Version,
	}
}

func round(num float64) int {
	return int(math.Floor(num + .5))
}

func toMegaBytes(bytes uint64) float64 {
	pow10 := math.Pow10(2.0)
	num := float64(bytes) / MB

	return float64(round(num*pow10)) / pow10
}
