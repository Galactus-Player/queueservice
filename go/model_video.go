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
	"time"
)

// Video - Video entry
type Video struct {

	Id string `json:"id,omitempty"`

	Url string `json:"url,omitempty"`

	ThumbnailUrl string `json:"thumbnailUrl,omitempty"`

	AddedAt time.Time `json:"addedAt,omitempty"`
}