package models

type File struct {
	ImageId   string `json:"image_id"`
	ProjectId string `json:"project_id"`
	Filename  string `json:"file_name"`
	Mimetype  string `json:"mime_type"`
	Size      string `json:"size"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
