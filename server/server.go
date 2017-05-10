package server

import (
	"context"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"message-persistence-service/elastic"
	"net/http"
	"os"
)

const organization string = "organization"
const project string = "project"
const id string = "id"
const id_pattern string = "~.*"

var log = logrus.New()

func Run() {
	ctx := context.Background()

	elasticClient, err := elastic.NewElasticClient(ctx)
	persistanceAdapter := elastic.NewElasticAdapter(ctx, elasticClient)
	getter := elastic.NewElasticGetter(ctx, elasticClient)

	writeHandler := &writeHandler{persistanceAdapter, GorillaVarsGetter{}}
	readHandler := &readHandler{getter, GorillaVarsGetter{}}

	if err != nil {
		spew.Dump(err)
		os.Exit(1)
	}

	router := mux.NewRouter()

	//Read routes
	router.Handle(fmt.Sprintf("/{%s}", organization), readHandler).Methods("GET")
	router.Handle(fmt.Sprintf("/{%s}/{%s:%s}", organization, id, id_pattern), readHandler).Methods("GET")
	router.Handle(fmt.Sprintf("/{%s}/{%s}", organization, project), readHandler).Methods("GET")
	router.Handle(fmt.Sprintf("/{%s}/{%s}/{%s:%s}", organization, project, id, id_pattern), readHandler).Methods("GET")

	//Write routes
	router.Handle(fmt.Sprintf("/{%s}", organization), writeHandler).Methods("POST")
	router.Handle(fmt.Sprintf("/{%s}/{%s}", organization, project), writeHandler).Methods("POST")

	router.Handle(fmt.Sprintf("/{%s}/{%s}", organization, id), writeHandler).Methods("UPDATE", "DELETE")
	router.Handle(fmt.Sprintf("/{%s}/{%s}/{%s}", organization, project, id), writeHandler).Methods("UPDATE", "DELETE")

	http.ListenAndServe(":2021", router)
}
