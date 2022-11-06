package animal

type Animal struct {
}

func (a *Animal) Add(l, r int64) int64 {
	return l + r
}
