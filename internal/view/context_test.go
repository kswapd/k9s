package view_test

import (
	"testing"

	"github.com/kswapd/k11s/internal/client"
	"github.com/kswapd/k11s/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	ctx := view.NewContext(client.NewGVR("contexts"))

	assert.Nil(t, ctx.Init(makeCtx()))
	assert.Equal(t, "Contexts", ctx.Name())
	assert.Equal(t, 4, len(ctx.Hints()))
}
