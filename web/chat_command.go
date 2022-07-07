package web

import (
	"encoding/json"
	"mtui/types"
	"mtui/types/command"
	"net/http"
	"time"
)

// fetch("api/bridge/execute_chatcommand", { method: "POST", body: JSON.stringify({ playername: "test", command: "status" })}).then(r => r.json()).then(res => console.log(res));
func (a *Api) ExecuteChatcommand(w http.ResponseWriter, r *http.Request, claims *types.Claims) {
	req := &command.ExecuteChatCommandRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		SendError(w, 500, err.Error())
		return
	}

	if req.Playername != claims.Username {
		// username does not match
		SendError(w, 500, "username mismatch")
		return
	}

	res, err := a.app.Bridge.ExecuteCommand(command.COMMAND_CHATCMD_REQ, req, time.Second*5)
	if err != nil {
		SendError(w, 500, err.Error())
		return
	}

	cmd, err := command.ParseCommand(res)
	if err != nil {
		SendError(w, 500, err.Error())
		return
	}

	SendJson(w, cmd)
}
