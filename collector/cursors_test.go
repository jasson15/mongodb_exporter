package collector

import (
	"github.com/dcu/mongodb_exporter/shared"
	"testing"
)

func Test_CursorsCollectData(t *testing.T) {
	cursors := &Cursors{}

	cursors.Export("cursors")
	if shared.Groups["cursors"] == nil {
		t.Error("Cursors group was not loaded.")
	}
}
