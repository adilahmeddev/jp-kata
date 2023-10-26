package jpservice_test

import (
	"testing"

	jpservice "github.com/adilahmeddev/jp-kata/backend/pkg/domain/jp-service"
	"github.com/adilahmeddev/jp-kata/backend/pkg/utils"
)

func TestGetUnique(t *testing.T) {
	jp := jpservice.JpProcessor{}

	gotRunes := jp.GetUnique("庭には二羽鶏がいる")
	expectedRunes := []string{"庭", "羽", "鶏", "二"}

	if utils.SlicesUnorderedEql(gotRunes, expectedRunes) {
		t.Errorf("got characters %v do not equal expected characters %v", gotRunes, expectedRunes)
	}
}
