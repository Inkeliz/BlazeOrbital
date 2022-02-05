package rendering

import (
	"fmt"
	"runtime"
	"syscall/js"
)

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func NewRender(processFrame func() (suns, satellites []float32)) {
	var callback js.Func
	callback = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			a, b := processFrame()
			renderSuns(a)
			renderSatellites(b)
			onNextFrame(callback)
		}()
		return nil
	})
	onNextFrame(callback)
}

func renderSuns([]float32)
func renderSatellites([]float32)
func onNextFrame(f js.Func)
