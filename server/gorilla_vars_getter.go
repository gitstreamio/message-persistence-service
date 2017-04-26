package server

import (
	"net/http"
	"github.com/gorilla/mux"
)

type GorillaVarsGetter struct{}

func (vg GorillaVarsGetter) Vars(r *http.Request) map[string]string {
	return mux.Vars(r)
}