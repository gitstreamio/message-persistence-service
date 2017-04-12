package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"message-persistence-service/common"
	"net/http"
)

type writeHandler struct {
	common.PersistenceAdapter
}

func (srv *writeHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		srv.handlePost(rw, r)
		break
	case "DELETE":
		srv.handleDelete(rw, r)
		break
	case "UPDATE":
		srv.handleUpdate(rw, r)
		break
	default:
		rw.WriteHeader(405)
	}
}
func (srv *writeHandler) handleUpdate(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	msg := &common.Message{Organization: vars[organization], Project: vars[project]}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(400)
		return
	}
	err = json.Unmarshal(body, msg)
	if err != nil || msg.Id == "" || msg.Message == "" || msg.User == "" {
		rw.WriteHeader(400)
		return
	}
	err = srv.Update(*msg)
	if err != nil {
		rw.WriteHeader(502)
		return
	}
	rw.WriteHeader(200)
}
func (srv *writeHandler) handleDelete(rw http.ResponseWriter, r *http.Request) {
	msg := &common.Message{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(400)
		return
	}
	err = json.Unmarshal(body, msg)
	if err != nil || msg.Id == ""{
		rw.WriteHeader(400)
		return
	}
	err = srv.Delete(msg.Id)
	if err != nil {
		rw.WriteHeader(502)
		return
	}
	rw.WriteHeader(200)
}

func (srv *writeHandler) handlePost(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	msg := &common.Message{Organization: vars[organization], Project: vars[project]}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(400)
		return
	}
	err = json.Unmarshal(body, msg)
	if err != nil || msg.Message == "" || msg.User == "" {
		rw.WriteHeader(400)
		return
	}
	id, err := srv.Persist(*msg)
	if err != nil {
		rw.WriteHeader(502)
		return
	}
	rw.Write([]byte(id))
}
