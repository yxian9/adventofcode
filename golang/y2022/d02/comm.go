package main

type round struct {
	p1, p2 int
}

// var m = map[string]int{
// 	"A": 1,
// 	"B": 2,
// 	"C": 3,
// 	"X": 1,
// 	"Y": 2,
// 	"Z": 3,
// }

// var win = map[round]bool{
// 	{3, 1}: true,
// 	{1, 2}: true,
// 	{2, 3}: true,
// }

// var loss = map[round]bool{
// 	{1, 3}: true,
// 	{2, 1}: true,
// 	{3, 2}: true,
// }

func winFunc(r round) bool {
	return r.p2 == (r.p1+1)%3
}

func score(r round) int {
	point := 0
	switch {
	case r.p1 == r.p2:
		point += 3
	case winFunc(r):
		point += 6
	}
	return point + r.p2 + 1
}

func myChose(p1, p2 int) int {
	switch p2 {
	case 0: // loss
		for i := range 3 {
			r := round{p1, i}
			if !winFunc(r) && i != p1 {
				return i
			}
		}
		// p1s := m[p1]
		// for k := range loss {
		// 	if k.p1 == p1s {
		// 		return k.p2
		// 	}
		// }
	case 1: // round
		return p1
	case 2: // win
		for i := range 3 {
			r := round{p1, i}
			if winFunc(r) {
				return i
			}
		}
	}
	return 0
}

func parser2(s string) round {
	p1, p2 := int(s[0]-'A'), int(s[2]-'X')
	return round{p1, myChose(p1, p2)}
}

func parser(s string) round {
	return round{int(s[0] - 'A'), int(s[2] - 'X')}
	// list := strings.Split(s, " ")
	// p1, p2 := list[0], list[1]
	//
	// return round{m[p1], m[p2]}
}
