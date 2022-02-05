package main

type Color struct {
	R, G, B, A float32
}

func NewColor(r, g, b, a float32) Color {
	return Color{R: r, G: g, B: b, A: a}
}

// Presets is public static readonly Color[] Presets = new[] { new Color(0.678f, 0.847f, 0.902f, 1), new Color(0.941f, 0.502f, 0.502f, 1), new Color(0.878f, 1.000f, 1.000f, 1), new Color(0.980f, 0.980f, 0.824f, 1), new Color(0.827f, 0.827f, 0.827f, 1), new Color(0.565f, 0.933f, 0.565f, 1), new Color(1.000f, 0.714f, 0.757f, 1), new Color(1.000f, 0.714f, 0.757f, 1), new Color(1.000f, 0.627f, 0.478f, 1), new Color(0.125f, 0.698f, 0.667f, 1), new Color(0.529f, 0.808f, 0.980f, 1), new Color(0.467f, 0.533f, 0.600f, 1), new Color(0.690f, 0.769f, 0.871f, 1), new Color(1.000f, 1.000f, 0.878f, 1), };
var Presets = [...]Color{
	{R: 0.678, G: 0.847, B: 0.902, A: 1},
	{R: 0.941, G: 0.502, B: 0.502, A: 1},
	{R: 0.878, G: 1.000, B: 1.000, A: 1},
	{R: 0.980, G: 0.980, B: 0.824, A: 1},
	{R: 0.827, G: 0.827, B: 0.827, A: 1},
	{R: 0.565, G: 0.933, B: 0.565, A: 1},
	{R: 1.000, G: 0.714, B: 0.757, A: 1},
	{R: 1.000, G: 0.714, B: 0.757, A: 1},
	{R: 1.000, G: 0.627, B: 0.478, A: 1},
	{R: 0.125, G: 0.698, B: 0.667, A: 1},
	{R: 0.529, G: 0.808, B: 0.980, A: 1},
	{R: 0.467, G: 0.533, B: 0.600, A: 1},
	{R: 0.690, G: 0.769, B: 0.871, A: 1},
	{R: 1.000, G: 1.000, B: 0.878, A: 1},
}
