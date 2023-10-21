package utils_test

import (
	"testing"

	"github.com/adilahmeddev/jp-kata/backend/pkg/utils"
)

func TestSlicesUnorderedEql(t *testing.T) {
	a := []int8{1, 5, 10, 4}
	b := []int8{10, 4, 5, 1}

	if !utils.SlicesUnorderedEql(a, b) {
		t.Error("slices should be equal")
	}
}

func TestSlicesUnorderedEql_shouldFail(t *testing.T) {

	a := []int8{1, 5, 10, 4}
	b := []int8{10, 4, 5, 6}

	if utils.SlicesUnorderedEql(a, b) {
		t.Error("slices should not be equal")
	}
}
