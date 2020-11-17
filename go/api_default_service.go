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

import "time"

// DefaultApiService is a service that implents the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// GetQueue -
func (s *DefaultApiService) GetQueue(code string) (interface{}, error) {
	// TODO - update GetQueue with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	v1 := Video{Id: "1", Url: "https://youtu.be/dQw4w9WgXcQ", AddedAt: time.Now()}
	v2 := Video{Id: "2", Url: "https://youtu.be/o61BiBCXMCI", AddedAt: time.Now()}
	v3 := Video{Id: "3", Url: "https://youtu.be/NEwcO0QwPAk", AddedAt: time.Now()}
	videos := [3]Video{v1, v2, v3}
	result := VideoQueue{Queue: videos[:], LastEdited: time.Now()}
	return result, nil
}