package ui_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/kswapd/k12s/internal/config"
	"github.com/kswapd/k12s/internal/render"
	"github.com/kswapd/k12s/internal/ui"
	"github.com/stretchr/testify/assert"
)

func TestBenchConfig(t *testing.T) {
	os.Setenv(config.K9sConfig, "/tmp/blee")
	assert.Equal(t, "/tmp/blee/bench-fred.yml", ui.BenchConfig("fred"))
}

func TestConfiguratorRefreshStyle(t *testing.T) {
	config.K9sStylesFile = filepath.Join("..", "config", "testdata", "black_and_wtf.yml")

	cfg := ui.Configurator{}
	cfg.RefreshStyles("")

	assert.True(t, cfg.HasSkin())
	assert.Equal(t, tcell.ColorGhostWhite.TrueColor(), render.StdColor)
	assert.Equal(t, tcell.ColorWhiteSmoke.TrueColor(), render.ErrColor)
}
