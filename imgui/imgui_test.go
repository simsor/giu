package imgui_test

import (
	"testing"

	"github.com/simsor/giu/imgui"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	version := imgui.Version()
	assert.Equal(t, "1.76", version)
}
