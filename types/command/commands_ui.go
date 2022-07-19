package command

// commands from the ui to ingame

import (
	"mtui/bridge"
)

const (
	COMMAND_CHATCMD_REQ bridge.CommandRequestType = "execute_command"
)

type ExecuteChatCommandRequest struct {
	Playername string `json:"playername"`
	Command    string `json:"command"`
}

type ExecuteChatCommandResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
