package api

import (
	"dotenv-updater/core/config"
	"encoding/json"
	"net/http"
)

func StatusList(w http.ResponseWriter, r *http.Request) {
	var list = map[string]interface{}{
		"status":  "up",
		"version": config.AppVersion(),
	}
	body, _ := json.Marshal(list)
	w.Write(body)
}
