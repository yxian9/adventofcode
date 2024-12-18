package utils

type Pt struct {
	X, Y int
}

func (p Pt) Move(dx, dy int) Pt {
	return Pt{p.X + dx, p.Y + dy}
}

func (p Pt) PMove(p2 Pt) Pt {
	return Pt{p.X + p2.X, p.Y + p2.Y}
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

type StringGrid struct {
	NRow, NCol int
	Array      []string
}

func (g *StringGrid) IsInside(pt Pt) bool {
	return pt.X >= 0 && pt.X < g.NRow && pt.Y >= 0 && pt.Y < g.NCol
}

func (g *StringGrid) PByte(pt Pt) byte {
	return g.Array[pt.X][pt.Y]
}

func (g *StringGrid) PInt(pt Pt) int {
	return int(g.PByte(pt) - '0')
}

func (g *StringGrid) PRune(pt Pt) rune {
	return rune(g.PByte(pt))
}

type Grid[T comparable] struct {
	NRow, NCol int
	Array      [][]T
}

func (g *Grid[T]) IsInside(pt Pt) bool {
	return pt.X >= 0 && pt.X < g.NRow && pt.Y >= 0 && pt.Y < g.NCol
}

func (g *Grid[T]) Get(pt Pt) T {
	return g.Array[pt.Y][pt.X]
}

func (g *Grid[T]) Swap(pt1, pt2 Pt) {
	g.Array[pt2.Y][pt2.X], g.Array[pt1.Y][pt1.X] = g.Array[pt1.Y][pt1.X], g.Array[pt2.Y][pt2.X]
}

func (g *Grid[T]) Find(start, dir Pt, target, ban T) (result Pt, find bool) {
	step := 1
	for {
		result = start.Move(dir.X*step, dir.Y*step)
		if !g.IsInside(result) || g.Get(result) == ban {
			return result, false
		}
		if g.Get(result) == target {
			return result, true
		}
		step++
	}
}
