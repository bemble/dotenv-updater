package api

import (
	"dotenv-updater/core/config"
	"dotenv-updater/core/dotenv"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type EnvsSetBody struct {
	Value string `json:"value"`
}

var DOTENV_FILE = config.GetTargetDotEnvFilePath()

func EnvsSet(w http.ResponseWriter, r *http.Request) {
	name := strings.Replace(chi.URLParam(r, "name"), "%2F", "/", -1)

	envBody := EnvsSetBody{}
	decodeErr := json.NewDecoder(r.Body).Decode(&envBody)

	if decodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		content := map[string]string{
			"message": "Decode error: " + decodeErr.Error(),
		}
		body, _ := json.Marshal(content)
		w.Write(body)
		return
	}

	setErr := dotenv.SetVar(DOTENV_FILE, name, envBody.Value)
	if setErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		content := map[string]string{
			"message": "Set env error: " + setErr.Error(),
		}
		body, _ := json.Marshal(content)
		w.Write(body)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func EnvsDelete(w http.ResponseWriter, r *http.Request) {
	name := strings.Replace(chi.URLParam(r, "name"), "%2F", "/", -1)
	err := dotenv.DeleteVar(DOTENV_FILE, name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		content := map[string]string{
			"message": "Delete env error: " + err.Error(),
		}
		body, _ := json.Marshal(content)
		w.Write(body)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
