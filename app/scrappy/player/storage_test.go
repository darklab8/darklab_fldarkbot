package player

import (
	"darkbot/app/settings/logus"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FixturePlayerStorageMockified() *PlayerStorage {
	return NewPlayerStorage(FixturePlayerAPIMock())
}

func TestGetPlayers(t *testing.T) {
	storage := FixturePlayerStorageMockified()
	storage.Update()

	bases, err := storage.GetLatestRecord()
	logus.CheckFatal(err, "not found latest base record")

	assert.True(t, len(bases.List) > 0)
	logus.Debug("", logus.Items(bases.List, "bases.List"))
}