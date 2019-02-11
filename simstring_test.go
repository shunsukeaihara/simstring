package simstring

import (
	"testing"
)

func TestSearch(t *testing.T) {
	db := MakeOnMemoryDB(MakeNgramEtractor(2))
	db.Add("ゴリラ")
	db.Add("チンパンジー")
	db.Add("ボノボ")
	db.Add("ニシゴリラ")
	db.Add("ウィルス性胃腸炎")
	db.Add("ロタウィルス性胃腸炎")

	if results := search("ウイルス性胃腸炎", 0.5, db, DiceIndex{}); len(results) != 2 {
		t.Fatal(results)
	}

}

func TestSimstring(t *testing.T) {
	s := MakeSimString()
	s.Add("ゴリラ")
	s.Add("チンパンジー")
	s.Add("ボノボ")
	s.Add("ニシゴリラ")
	s.Add("ウィルス性胃腸炎")
	s.Add("ロタウィルス性胃腸炎")

	if results := s.Lookup("ウイルス性胃腸炎", 0.5); len(results) != 2 {
		t.Fatal(results)
	}
}
