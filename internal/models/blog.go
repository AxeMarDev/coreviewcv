package models

type Blog struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	AuthorId   int    `json:"author_id"`
	DatePoster int    `json:"date_posted"`
}
