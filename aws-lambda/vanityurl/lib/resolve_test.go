package lib

import (
	"reflect"
	"testing"
)

func TestResolve(t *testing.T) {
	url := "/volkswagen/golf"

	got := Resolve(url)
	want := make(map[string]Match)
	want["volkswagen"] = Match{word: "volkswagen", filterKey: "t_manuf", filterOption: SearchCriteriaItem{Key: "BQ", Name: "Volkswagen"}}
	want["golf"] = Match{word: "golf", filterKey: "t_model", filterOption: SearchCriteriaItem{Key: "BQAK", Name: "Golf"}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestResolveBad(t *testing.T) {
	url := "/volkswagen1/"
	got := Resolve(url)
	want := make(map[string]Match)
	want["volkswagen1"] = Match{word: "volkswagen1", filterKey: "no-match"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
