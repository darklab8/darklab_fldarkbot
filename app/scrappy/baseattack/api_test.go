package baseattack

import (
	"darkbot/app/settings/logus"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {
	api := FixtureBaseAttackAPIMock()
	result, _ := api.GetBaseAttackData()
	data := string(result)
	logus.Debug(data)
}

func TestDetectLPAttack(t *testing.T) {
	api := NewMock("data_lp.json")
	result, _ := api.GetBaseAttackData()
	data := string(result)
	assert.True(t, strings.Contains(data, "LP-7743"))
}