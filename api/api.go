package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router(r chi.Router) {
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(CheckApiKey)

	r.Get(`/status`, StatusList)
	r.Post(`/env/{name}`, EnvsSet)
	r.Delete(`/env/{name}`, EnvsDelete)
}
