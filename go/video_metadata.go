package openapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func getVideoMetadata(videoUrl string) (string, string, error) {
	vid := getYouTubeID(videoUrl)
	thumb := fmt.Sprintf("https://img.youtube.com/vi/%s/0.jpg", vid)

	req_str := fmt.Sprintf("https://www.youtube.com/oembed?url=http://www.youtube.com/watch?v=%s&format=json", vid)
	resp, err := http.Get(req_str)
	log.Printf("Error meta: %+v\n", err)
	if err != nil {
		return thumb, "", nil
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return thumb, result["title"].(string), nil
}

func getYouTubeID(videoUrl string) string {
	u, err := url.Parse(videoUrl)
	if err != nil {
		return ""
	}
	q := u.Query()
	return q.Get("v")
}
