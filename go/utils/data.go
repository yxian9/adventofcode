package utils

type Pt struct {
	R, C int
}

func (p Pt) Move(dc, dr int) Pt {
	return Pt{C: p.C + dc, R: p.R + dr}
}

func (p Pt) PMove(p2 Pt) Pt {
	return Pt{C: p.C + p2.C, R: p.R + p2.R}
}

func (p Pt) Dist(p2 Pt) (dr, dc int) {
	dc = p.C - p2.C
	dr = p.R - p2.R
	return dr, dc
}

var Dir4 = []Pt{
	{R: -1, C: 0},
	{R: 0, C: 1},
	{R: 1, C: 0},
	{R: 0, C: -1},
}

type StringGrid struct {
	Array []string
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
	Array [][]T
}

func (g *Grid[T]) IsInside(pt Pt) bool {
	nrow, ncol := len(g.Array), len(g.Array[0])
	return pt.C >= 0 && pt.C < ncol && pt.R >= 0 && pt.R < nrow
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

type PtSeen map[Pt]bool
