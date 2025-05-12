package model

type Post struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Title     string `json:"title" bson:"title" validate:"required"`
	Author    string `json:"author" bson:"author" validate:"required"`
	PageCount int    `json:"page_count" bson:"page_count" validate:"required,gt=0"`
	InStock   bool   `json:"in_stock" bson:"in_stock" validate:"required"`
}
