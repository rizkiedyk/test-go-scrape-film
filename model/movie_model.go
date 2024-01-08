package model

type Movie struct {
	Title    string   `json:"title"`
	Rating   string   `json:"rating"`
	Quality  string   `json:"quality"`
	Duration string   `json:"duration"`
	Genre    []string `json:"genre"`
	URLImage string   `json:"url_image"`
}
