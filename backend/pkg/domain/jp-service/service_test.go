package jpservice_test

import (
	"testing"

	jpservice "github.com/adilahmeddev/jp-kata/backend/pkg/domain/jp-service"
	"github.com/adilahmeddev/jp-kata/backend/pkg/utils"
)

func TestGetUnique(t *testing.T) {
	jp := jpservice.JpProcessor{}

	gotRunes := jp.GetUnique("漢ウキウク")
	expectedRunes := []rune{'ウ', 'キ', 'ク'}

	if utils.SlicesUnorderedEql(gotRunes, expectedRunes) {
		t.Errorf("got characters %v do not equal expected characters %v", gotRunes, expectedRunes)
	}
}
