package render_test

import (
	"testing"

	"github.com/kswapd/k12s/internal/render"
	"github.com/stretchr/testify/assert"
)

func TestReferenceRender(t *testing.T) {
	o := render.ReferenceRes{
		Namespace: "ns1",
		Name:      "blee",
		GVR:       "v1/secrets",
	}

	var (
		ref = render.Reference{}
		r   render.Row
	)
	assert.Nil(t, ref.Render(o, "fred", &r))
	assert.Equal(t, "ns1/blee", r.ID)
	assert.Equal(t, render.Fields{
		"ns1",
		"blee",
		"v1/secrets",
	}, r.Fields)
}
