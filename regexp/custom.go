package regexp

type InstPp struct {
	I    int
	Op   string
	Out  int
	Arg  int
	Rune []string
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
		}
	}
	return r
}
