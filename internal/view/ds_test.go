package view_test

import (
	"testing"

	"github.com/kswapd/k11s/internal/client"
	"github.com/kswapd/k11s/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestDaemonSet(t *testing.T) {
	v := view.NewDaemonSet(client.NewGVR("apps/v1/daemonsets"))

	assert.Nil(t, v.Init(makeCtx()))
	assert.Equal(t, "DaemonSets", v.Name())
	assert.Equal(t, 15, len(v.Hints()))
}
