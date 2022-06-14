package types_test

import (
	"mtui/bridge"
	"mtui/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

const UNKNOWN_RES bridge.CommandResponseType = "whatever"

func TestParseCommand(t *testing.T) {
	resp := &bridge.CommandResponse{
		Type: types.COMMAND_PING_RES,
		Data: []byte("{}"),
	}

	o, err := types.ParseCommand(resp)
	assert.NoError(t, err)
	assert.NotNil(t, o)

	resp = &bridge.CommandResponse{
		Type: UNKNOWN_RES,
		Data: []byte("{}"),
	}

	o, err = types.ParseCommand(resp)
	assert.NoError(t, err)
	assert.Nil(t, o)
}
