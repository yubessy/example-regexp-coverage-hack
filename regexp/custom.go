package regexp

import "github.com/yubessy/golang-regexp-customize/regexp/syntax"

type InstPp struct {
	I    int
	Op   string
	Out  int
	Arg  int
	Rune []string
	Flag bool
}

func (re *Regexp) Inst() []InstPp {
	r := make([]InstPp, len(re.prog.Inst))
	for i, x := range re.prog.Inst {
		rn := make([]string, len(x.Rune))
		for j, y := range x.Rune {
			rn[j] = string(y)
		}
		r[i] = InstPp{
			I:    i,
			Op:   x.Op.String(),
			Out:  int(x.Out),
			Arg:  int(x.Arg),
			Rune: rn,
			Flag: x.Flag,
		}
	}
	return r
}

func (re *Regexp) Coverage() (int, int) {
	c := 0
	a := 0
	for _, x := range re.prog.Inst {
		if x.Op == syntax.InstRune || x.Op == syntax.InstRune1 || x.Op == syntax.InstRuneAny || x.Op == syntax.InstRuneAnyNotNL {
			a++
			if x.Flag {
				c++
			}
		}
	}
	return c, a
}
