package utils

type Pt struct {
	C, R int
}

func (p Pt) Move(dc, dr int) Pt {
	return Pt{p.C + dc, p.R + dr}
}

func (p Pt) PMove(p2 Pt) Pt {
	return Pt{p.C + p2.C, p.R + p2.R}
}

func (p Pt) Dist(p2 Pt) (dc, dr int) {
	dc = p.C - p2.C
	dr = p.R - p2.R
	return dc, dr
}

var Dir4 = []Pt{
	{C: 0, R: 1},
	{C: -1, R: 0},
	{C: 0, R: -1},
	{C: 1, R: 0},
}

type StringGrid struct {
	NCol, NRow int
	Array      []string
}

func (g *StringGrid) IsInside(pt Pt) bool {
	nrow, ncol := len(g.Array), len(g.Array[0])
	return pt.C >= 0 && pt.C < ncol && pt.R >= 0 && pt.R < nrow
}

func (g *StringGrid) GetByte(pt Pt) byte {
	return g.Array[pt.R][pt.C]
}

func (g *StringGrid) GetInt(pt Pt) int {
	return int(g.GetByte(pt) - '0')
}

func (g *StringGrid) GetRune(pt Pt) rune {
	return rune(g.GetByte(pt))
}

type Grid[T comparable] struct {
	NRow, NCol int
	Array      [][]T
}

func (g *Grid[T]) IsInside(pt Pt) bool {
	return pt.C >= 0 && pt.C < g.NCol && pt.R >= 0 && pt.R < g.NRow
}

func (g *Grid[T]) Get(pt Pt) T {
	return g.Array[pt.R][pt.C]
}

func (g *Grid[T]) Set(pt Pt, v T) {
	g.Array[pt.R][pt.C] = v
}

func (g *Grid[T]) Swap(pt1, pt2 Pt) {
	g.Array[pt2.R][pt2.C], g.Array[pt1.R][pt1.C] = g.Array[pt1.R][pt1.C], g.Array[pt2.R][pt2.C]
}

func (g *Grid[T]) Find(start, dir Pt, target, ban T) (result Pt, find bool) {
	step := 1
	for {
		result = start.Move(dir.C*step, dir.R*step)
		if !g.IsInside(result) || g.Get(result) == ban {
			return result, false
		}
		if g.Get(result) == target {
			return result, true
		}
		step++
	}
}

type Seen map[Pt]bool
