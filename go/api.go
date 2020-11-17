/*
 * queueservice
 *
 * Video queue service
 *
 * API version: 0.0.1
 * Contact: sbirudavolu@umass.edu
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"
)


// DefaultApiRouter defines the required methods for binding the api requests to a responses for the DefaultApi
// The DefaultApiRouter implementation should parse necessary information from the http request, 
// pass the data to a DefaultApiServicer to perform the required actions, then write the service results to the http response.
type DefaultApiRouter interface { 
	GetQueue(http.ResponseWriter, *http.Request)
}
// QueueApiRouter defines the required methods for binding the api requests to a responses for the QueueApi
// The QueueApiRouter implementation should parse necessary information from the http request, 
// pass the data to a QueueApiServicer to perform the required actions, then write the service results to the http response.
type QueueApiRouter interface { 
	AddVideo(http.ResponseWriter, *http.Request)
}


// DefaultApiServicer defines the api actions for the DefaultApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type DefaultApiServicer interface { 
	GetQueue(string) (interface{}, error)
}


// QueueApiServicer defines the api actions for the QueueApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type QueueApiServicer interface { 
	AddVideo(string, AddVideoRequest) (interface{}, error)
}