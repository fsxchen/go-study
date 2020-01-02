package entity

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title           string `json:"title"`
	Desc            string `json:"desc"`
	Content         string `json:"content"`
	ContentMarkdown string `json:"content_markdown"`
	CreatedBy       string `json:"created_by"`
	ModifiedBy      string `json:"modified_by"`
	State           int    `json:"state"`
}
