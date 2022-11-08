package tchart_test

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/kswapd/k11s/internal/tchart"
	"github.com/stretchr/testify/assert"
)

func TestCoSeriesColorNames(t *testing.T) {
	c := tchart.NewComponent("fred")

	c.SetSeriesColors(tcell.ColorGreen, tcell.ColorBlue)

	assert.Equal(t, []string{"green", "blue"}, c.GetSeriesColorNames())
}

func TestComponentAsRect(t *testing.T) {
	c := tchart.NewComponent("fred")

	c.SetSeriesColors(tcell.ColorGreen, tcell.ColorBlue)

	assert.Equal(t, []string{"green", "blue"}, c.GetSeriesColorNames())
}
