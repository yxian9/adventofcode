package utils

type Pt struct {
	X, Y int
}

func (p Pt) Move(dx, dy int) Pt {
	return Pt{p.X + dx, p.Y + dy}
}

func (p Pt) Dist(p2 Pt) (dx, dy int) {
	dx = p.X - p2.X
	dy = p.Y - p2.Y
	return
}

var Dir4 = []Pt{
	{X: 0, Y: 1},
	{X: -1, Y: 0},
	{X: 0, Y: -1},
	{X: 1, Y: 0},
}

type Grid struct {
	NRow, NCol int
	Array      []string
}

func (g *Grid) IsInside(pt Pt) bool {
	return pt.X >= 0 && pt.X < g.NRow && pt.Y >= 0 && pt.Y < g.NCol
}

func (g *Grid) PByte(pt Pt) byte {
	return g.Array[pt.X][pt.Y]
}

func (g *Grid) PInt(pt Pt) int {
	return int(g.PByte(pt) - '0')
}

func (g *Grid) PRune(pt Pt) rune {
	return rune(g.PByte(pt))
}
