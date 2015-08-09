package gtnd

import (
	"testing"
)

func TestSearch(t *testing.T) {
	param := SearchParam{};
	_, err := Search(param)
	if err != nil {
		t.Errorf("failtest");
	}
}
