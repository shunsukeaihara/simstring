package simstring

import (
	"sort"
)

type Searcher struct {
	Measure
}

type featureToWords struct {
	feature string
	words   map[string]struct{}
	found   bool
}

func (f featureToWords) Count() int {
	if f.found {
		return len(f.words)
	} else {
		return 0
	}
}

func (f featureToWords) HasWord(word string) bool {
	if f.found {
		_, ok := f.words[word]
		return ok
	} else {
		return false
	}
}

type featureResults []featureToWords

func (f featureResults) Len() int {
	return len(f)
}

func (f featureResults) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f featureResults) Less(i, j int) bool {
	return f[i].Count() > f[j].Count()
}

func search(q string, alpha float64, db DB, measure Measure) map[string]struct{} {
	features := db.Extract(q)
	size := len(features)
	minSize := measure.MinSize(size, alpha)
	maxSize := measure.MaxSize(size, alpha)

	results := make(map[string]struct{})
	for i := minSize; i <= maxSize; i++ {
		tau := measure.MinMatch(size, i, alpha)
		for word, _ := range overlapJoin(features, size, i, tau, db) {
			results[word] = struct{}{}
		}
	}
	return results
}

func overlapJoin(features []string, fsize int, candidateSize int, tau int, db DB) map[string]struct{} {
	lookedups := make(featureResults, 0, fsize)
	for _, feature := range features {
		words, ok := db.Lookup(candidateSize, feature)
		lookedups = append(lookedups, featureToWords{feature, words, ok})
	}
	sort.Sort(lookedups)
	wordMatchedCounts := make(map[string]int)
	for i, fw := range lookedups {
		if i > fsize-tau+1 {
			break
		}
		for word, _ := range fw.words {
			wordMatchedCounts[word] += 1
		}
	}
	//tmpWords := make([]string, 0, len(wordMatchedCounts))
	//for word, _ := range wordMatchedCounts {
	//	tmpWords = append(tmpWords, word)
	//}
	results := make(map[string]struct{})
	for word, _ := range wordMatchedCounts {
		for i := fsize - tau + 1; i < fsize; i++ {
			if lookedups[i].HasWord(word) {
				wordMatchedCounts[word] += 1
			}
			if wordMatchedCounts[word] > tau {
				results[word] = struct{}{}
				break
			}
			remains := fsize - i + 1
			if wordMatchedCounts[word]+remains < tau {
				break
			}
		}
	}

	return results
}
