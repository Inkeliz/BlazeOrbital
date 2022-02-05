package main

import (
	"github.com/inkeliz/satellites/rendering"
	"strconv"
	"sync"
	"syscall/js"
	"unsafe"
)

var (
	satellitesLengthMutex sync.Mutex
	satellitesLength      = 100
)

var (
	satellites = make([]BodyVisual, 0, 1_000_000)
	suns       = make([]BodyVisual, 2)

	satellitesInfo = make([]BodyInfo, 0, 1_000_000)
	sunsInfo       = make([]BodyInfo, 2)
)

func main() {
	OnInitialized()

	document := js.Global().Get("document")
	querySelector := document.Get("querySelector").Call("bind", document)

	input := querySelector.Invoke("input")
	input.Set("value", strconv.FormatInt(int64(satellitesLength), 10))
	input.Call("addEventListener", "input", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		i, err := strconv.ParseUint(args[0].Get("target").Get("value").String(), 10, 63)
		if err != nil {
			return nil
		}
		SpawnSatellites(int(i))
		return nil
	}))

	button := querySelector.Invoke("button")
	button.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		OnInitialized()
		return nil
	}))

	rendering.NewRender(func() (a, b []float32) {
		OnRenderFrame()
		return *(*[]float32)(unsafe.Pointer(&suns)), *(*[]float32)(unsafe.Pointer(&satellites))
	})

	select {}
}

func OnInitialized() {
	/*
		suns[0] = new Body { Pos = new Point(5000, 3700), Vel = new Point(22, 0), Mass = 2000000, Color = new Color(1, 1, 0, 1) };
		suns[1] = new Body { Pos = new Point(5000, 6300), Vel = new Point(-22, 0), Mass = 2000000, Color = new Color(1, 1, 0, 1) };
		SpawnSatellites(satellites.Length);
	*/
	suns[0] = BodyVisual{Pos: NewPoint(5000, 3700), Color: NewColor(1, 1, 0, 1)}
	sunsInfo[0] = BodyInfo{Vel: NewPoint(22, 0), Mass: 2_000_000}
	suns[1] = BodyVisual{Pos: NewPoint(5000, 6300), Color: NewColor(1, 1, 0, 1)}
	sunsInfo[1] = BodyInfo{Vel: NewPoint(-22, 0), Mass: 2_000_000}
	SpawnSatellites(satellitesLength)
}

func SpawnSatellites(amount int) {
	/*
		void SpawnSatellites(int count) = > satellites = Enumerable.Range(0, count).Select(_ = > new Body().Respawn()).ToArray()
	*/
	satellitesLengthMutex.Lock()
	defer satellitesLengthMutex.Unlock()
	l := len(satellites)

	switch {
	case l == amount:
		return
	case l > amount:
		satellites = satellites[:amount]
	case l < amount:
		satellites = append(satellites, make([]BodyVisual, amount-l)...)
		satellitesInfo = append(satellitesInfo, make([]BodyInfo, amount-l)...)
		for i := l; i < amount-l; i++ {
			satellites[i].Respawn()
			satellitesInfo[i].Respawn()
		}
	}
}

func OnRenderFrame() {
	/*
		for (var i = 0; i < suns.Length; i++)
		RunPhysicsForBody(ref suns[i]);
		for (var i = 0; i < satellites.Length; i++)
		RunPhysicsForBody(ref satellites[i]);
	*/

	for i := 0; i < len(suns); i++ {
		RunPhysicsForBody(&sunsInfo[i], &suns[i])
	}
	for i := 0; i < len(satellites); i++ {
		RunPhysicsForBody(&satellitesInfo[i], &satellites[i])
	}
}
