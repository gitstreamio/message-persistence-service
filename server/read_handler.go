package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
