package server

import (
	"encoding/json"
	"fmt"
	"message-persistence-service/common"
	"net/http"
)

type readHandler struct {
	common.Getter
	common.MuxVarsGetter
}

func (srv *readHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Info("Get")
	vars := srv.Vars(r)
	timeline := getTimeline(vars)
	id, ok := vars[id]

	if ok {
		msg, err := srv.GetById(id[1:])
		if err != nil {
			log.WithError(err).Error("Could not get Message by id")
			rw.WriteHeader(502)
			return
		}
		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			log.WithError(err).Error("Could not marshal message")
			rw.WriteHeader(502)
			return

		}
		rw.Write(jsonMsg)
		return
	} else {
		msgs, err := srv.Get(timeline, 0, 10)
		if err != nil {
			log.WithError(err).WithField("timeline", timeline).Error("Could not fetch messages for timeline")
			rw.WriteHeader(502)
		}
		jsonMsgs, err := json.Marshal(msgs)
		if err != nil {
			log.WithError(err).Error("Could not marshal messages")
			rw.WriteHeader(502)
			return
		}
		rw.Write(jsonMsgs)
		return
	}

}

func getTimeline(vars map[string]string) string {
	var timeline string
	_, hasProject := vars[project]
	if hasProject {
		timeline = fmt.Sprintf("%s/%s", vars[organization], vars[project])
	} else {
		timeline = fmt.Sprintf("%s", vars[organization])
	}
	return timeline
}
