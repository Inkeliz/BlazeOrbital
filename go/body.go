package main

import (
	_ "unsafe"
)

type BodyVisual struct {
	Pos   Point
	Color Color
}

func (body *BodyVisual) Respawn() {
	/*
		Pos = new Point(Random.Shared.Next(10000), Random.Shared.Next(10000));
		Color = Color.Presets[Random.Shared.Next(Color.Presets.Length)];
	*/
	body.Pos = NewPoint(float32(fastrandn(10000)), float32(fastrandn(10000)))
	body.Color = Presets[fastrandn(uint32(len(Presets)))]
}

type BodyInfo struct {
	Vel  Point
	Mass float32
}

func (body *BodyInfo) Respawn() {
	/*
		Vel = new Point((float)(Random.Shared.NextDouble() * 2 - 1), ((float)Random.Shared.NextDouble() * 2 - 1)) * 25f;
	*/
	body.Vel = NewPoint(fastrandFloat32()*2-1, fastrandFloat32()*2-1).MulScale(25)
}

func RunPhysicsForBody(info *BodyInfo, visual *BodyVisual) {
	/*
		// For each sun...
		for (var sunIndex = 0; sunIndex < suns.Length; sunIndex++)
		{
			ref var sun = ref suns[sunIndex];
			if (body.Pos != sun.Pos)
			{
				// ... apply sun's gravitational force to the body
				const float G = 1f;
				var deltaPos = sun.Pos - body.Pos;
				var acceleration = G * sun.Mass / deltaPos.LengthSquared;
				body.Vel += deltaPos.Direction * acceleration;
			}
		}
	*/
	for sunIndex := 0; sunIndex < len(suns); sunIndex++ {
		if visual.Pos != suns[sunIndex].Pos {
			const G float32 = 1
			deltaPos := suns[sunIndex].Pos.Sub(visual.Pos)
			acceleration := G * sunsInfo[sunIndex].Mass / deltaPos.LengthSquared()
			info.Vel = info.Vel.Add(deltaPos.Direction().MulScale(acceleration))
		}
	}

	/*
		// Move it
		body.Pos += body.Vel;
	*/
	visual.Pos = visual.Pos.Add(info.Vel)

	/*
		// If it's lost in space, respawn it
		if (body.Pos.X < -1000 || body.Pos.X > 11000 || body.Pos.Y < -1000 || body.Pos.Y > 11000)
			body.Respawn();
		    }
	*/
	if visual.Pos.X < -1000 || visual.Pos.X > 11000 || visual.Pos.Y < -1000 || visual.Pos.Y > 11000 {
		info.Respawn()
		visual.Respawn()
	}
}

//go:linkname fastrandn runtime.fastrandn
func fastrandn(n uint32) uint32

//go:linkname fastrandFloat32 runtime.fastrand
func fastrandFloat32() float32
