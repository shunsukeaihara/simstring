package simstring

type SimString struct {
	db      DB
	measure Measure
}

func MakeSimString() *SimString {
	db := MakeOnMemoryDB(MakeNgramEtractor(2))
	measure := DiceIndex{}
	return &SimString{db, measure}
}

func (s *SimString) Add(w string) {
	s.db.Add(w)
}

func (s *SimString) Lookup(w string, alpha float64) map[string]struct{} {
	return search(w, alpha, s.db, s.measure)
}
