/*
 * queueservice
 *
 * Video queue service
 *
 * API version: 0.0.1
 * Contact: sbirudavolu@umass.edu
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
	"github.com/tkanos/gonfig"
)

type Config struct {
	DbName     string
	DbPassword string
}

func main() {
	log.Printf("Server started")

	config := Config{}
	err := gonfig.GetConf("./config.json", &config)
	if err != nil {
		log.Fatal(err)
	}

	db, err := openapi.DbConnect(config.DbName, config.DbPassword)
	if err != nil {
		log.Fatal(err)
	}

	DefaultApiService := openapi.NewDefaultApiService(db)
	DefaultApiController := openapi.NewDefaultApiController(DefaultApiService)

	QueueApiService := openapi.NewQueueApiService(db)
	QueueApiController := openapi.NewQueueApiController(QueueApiService)

	router := openapi.NewRouter(DefaultApiController, QueueApiController)

	log.Fatal(http.ListenAndServe(":9090", router))
}
