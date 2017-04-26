package server

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"message-persistence-service/elastic"
	"net/http"
	"os"
	"github.com/Sirupsen/logrus"
)

const organization string = "organization"
const project string = "project"
const id string = "id"

var log = logrus.New()

func Run() {
	ctx := context.Background()

	elasticClient, err := elastic.NewElasticClient(ctx)
	persistanceAdapter:= elastic.NewElasticAdapter(ctx, elasticClient)
	writeHandler := &readHandler{persistanceAdapter, GorillaVarsGetter{}}

	if err != nil {
		spew.Dump(err)
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.Handle(fmt.Sprintf("/{%s}", organization), writeHandler).Methods("POST")
	router.Handle(fmt.Sprintf("/{%s}/{%s}", organization, project), writeHandler).Methods("POST")

	router.Handle(fmt.Sprintf("/{%s}/{%s}", organization, id), writeHandler).Methods("UPDATE", "DELETE")
	router.Handle(fmt.Sprintf("/{%s}/{%s}/{%s}", organization, project, id), writeHandler).Methods("UPDATE", "DELETE")

	http.ListenAndServe(":2021", router)
}
