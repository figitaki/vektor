package main

import (
	"net/http"
	"strings"

	"github.com/suborbital/vektor/vk"
)

type reqScope struct {
	ReqID string `json:"req_id"`
}

func setScopeMiddleware(r *http.Request, ctx *vk.Ctx) error {
	scope := reqScope{
		ReqID: "asdfghj",
	}

	ctx.UseScope(scope)

	return nil
}

func denyMiddleware(r *http.Request, ctx *vk.Ctx) error {
	if strings.Contains(r.URL.Path, "hack") {
		ctx.Log.ErrorString("HACKER!!")

		return vk.E(403, "begone, hacker")
	}

	return nil
}

func headerMiddleware(r *http.Request, ctx *vk.Ctx) error {
	ctx.RespHeaders.Set("X-Vektor-Test", "foobar")

	return nil
}
