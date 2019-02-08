package simstring

import (
	"math"
)

type Measure interface {
	MinSize(int, float64) int
	MaxSize(int, float64) int
	MinMatch(int, int, float64) int
	Similarity([]string, []string) float64
}

type DiceIndex struct {
}

func (DiceIndex) MinSize(qsize int, alpha float64) int {
	return int(math.Ceil(alpha * 1.0 / (2 - alpha) * float64(qsize)))
}

func (DiceIndex) MaxSize(qsize int, alpha float64) int {
	return int(math.Floor((2.0 - alpha) * float64(qsize) * 1.0 / alpha))
}

func (DiceIndex) MinMatch(qsize int, ysize int, alpha float64) int {
	return int(math.Ceil(0.5 * alpha * float64(qsize+ysize)))
}

func (DiceIndex) Similarity(a []string, b []string) float64 {
	return 0.1
}
