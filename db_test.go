package simstring

import (
	"testing"
)

func TestSizeWordMap(t *testing.T) {
	m := MakeSizeWordMap()
	features := makeUnicodeNgram("ゴリラ", 2)
	m.Add(len(features), "ゴリラ")
	mapping, ok := m.Lookup(len(features))
	if !ok {
		t.Fatal("Not Found")
	}
	if _, ok := mapping["ゴリラ"]; !ok {
		t.Fatal("Not Found")
	}
}

func TestSizeFeatureWordMap(t *testing.T) {
	m := MakeSizeFeatureWordMap()
	features := makeUnicodeNgram("ゴリラ", 2)
	for _, f := range features {
		m.Add(len(features), f, "ゴリラ")
	}

	for _, f := range features {
		mapping, ok := m.Lookup(len(features), f)
		if !ok {
			t.Fatal("Not Found")
		}
		if _, ok := mapping["ゴリラ"]; !ok {
			t.Fatal("Not Found")
		}
	}
}
