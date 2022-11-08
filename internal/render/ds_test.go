package render_test

import (
	"testing"

	"github.com/kswapd/k12s/internal/render"
	"github.com/stretchr/testify/assert"
)

func TestDaemonSetRender(t *testing.T) {
	c := render.DaemonSet{}
	r := render.NewRow(9)
	c.Render(load(t, "ds"), "", &r)

	assert.Equal(t, "kube-system/fluentd-gcp-v3.2.0", r.ID)
	assert.Equal(t, render.Fields{"kube-system", "fluentd-gcp-v3.2.0", "2", "2", "2", "2", "2"}, r.Fields[:7])
}
