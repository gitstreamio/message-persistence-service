package server

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"io/ioutil"
	"message-persistence-service/common"
	"net/http"
)

var log = logrus.New()

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
	log.Info("Update")
	vars := mux.Vars(r)
	_, hasProject := vars[project]

	var msg *common.Message
	if hasProject {
		msg = &common.Message{Timeline: fmt.Sprintf("%s/%s", vars[organization], vars[project])}
	} else {
		msg = &common.Message{Timeline: fmt.Sprintf("%s", vars[organization])}
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(400)
		return
	}
	err = json.Unmarshal(body, msg)

	id, hasId := vars[id]

	if err != nil || !hasId || id == "" || msg.Body == "" || msg.Author == "" {
		log.WithError(err).
			WithField("body", msg.Body).
			WithField("id", id).
			WithField("author", msg.Author).
			Error("Create failed")
		rw.WriteHeader(400)
		return
	}
	err = srv.Update(id, *msg)
	if err != nil {
		log.WithError(err).Error("Update failed")
		rw.WriteHeader(502)
		return
	}
	rw.WriteHeader(200)
}
func (srv *writeHandler) handleDelete(rw http.ResponseWriter, r *http.Request) {
	log.Info("Delete")
	vars := mux.Vars(r)
	id, ok := vars[id]

	if !ok || id == "" {
		rw.WriteHeader(400)
		return
	}
	err := srv.Delete(id)
	if err != nil {
		log.WithError(err).Error("Delete failed")
		rw.WriteHeader(502)
		return
	}
	rw.WriteHeader(200)
}

func (srv *writeHandler) handlePost(rw http.ResponseWriter, r *http.Request) {
	log.Info("Create")
	vars := mux.Vars(r)

	_, hasProject := vars[project]
	var msg *common.Message
	if hasProject {
		msg = &common.Message{Timeline: fmt.Sprintf("%s/%s", vars[organization], vars[project])}
	} else {
		msg = &common.Message{Timeline: fmt.Sprintf("%s", vars[organization])}
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).Error("Create failed")
		rw.WriteHeader(400)
		return
	}
	err = json.Unmarshal(body, msg)
	if err != nil || msg.Body == "" || msg.Author == "" {
		log.WithError(err).
			WithField("body", msg.Body).
			WithField("author", msg.Author).
			Error("Create failed")
		rw.WriteHeader(400)
		return
	}
	id, err := srv.Create(*msg)
	if err != nil {
		log.WithError(err).Error("Create failed")
		rw.WriteHeader(502)
		return
	}
	rw.Write([]byte(id))
}
