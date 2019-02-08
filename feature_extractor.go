package simstring

const SENTINEL = " " // non breaking space

func makeUnicodeNgram(s string, n int) []string {
	s2 := []rune(SENTINEL + s + SENTINEL)
	ret := make([]string, 0, len(s2)-n+1)
	for i := 0; i < len(s2)-n+1; i++ {
		ret = append(ret, string(s2[i:i+n]))
	}
	return ret
}

type FeatureExtractor interface {
	Extract(string) []string
}

type NgramExtractor struct {
	n int
}

func MakeNgramEtractor(n int) *NgramExtractor {
	return &NgramExtractor{n: n}
}

func (n NgramExtractor) Extract(w string) []string {
	return makeUnicodeNgram(w, n.n)
}
