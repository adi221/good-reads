package model

import (
	"time"

	"github.com/adi221/good-reads/pkg/util"
)

// ArticleCreateForm structure definition
type ArticleCreateForm struct {
	Title      string  `json:"title,omitempty"`
	Text       *string `json:"text,omitempty"`
	HTML       *string `json:"html,omitempty"`
	URL        *string `json:"url,omitempty"`
	Image      *string `json:"image,omitempty"`
	CategoryID *uint   `json:"categoryId,omitempty"`
}

// TruncatedTitle return truncated title
func (form ArticleCreateForm) TruncatedTitle() string {
	return util.Truncate(form.Title, 29)
}

// IsComplete tests if the form is complete
func (form ArticleCreateForm) IsComplete() bool {
	return !util.OneIsEmpty(form.Image, form.Text, form.HTML)
}

// Hash returns form hash
func (form ArticleCreateForm) Hash() string {
	key := form.Title
	if form.URL != nil {
		key += *form.URL
	}
	if form.HTML != nil {
		key += *form.HTML
	}
	return util.Hash(key)
}

// Article structure definition
type Article struct {
	ID         uint       `json:"id,omitempty"`
	UserID     *uint      `json:"userId,omitempty"`
	CategoryID *uint      `json:"categoryId,omitempty"`
	Title      string     `json:"title,omitempty"`
	Text       *string    `json:"text,omitempty"`
	HTML       *string    `json:"html,omitempty"`
	URL        *string    `json:"url,omitempty"`
	Image      *string    `json:"image,omitempty"`
	Hash       string     `json:"hash,omitempty"`
	Status     string     `json:"status,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

// Codes for article HttpErrors
const (
	CategoryNotProvided = "ERR_CATEGORY_NOT_PROVIDED"
	ArticleNonExist     = "ERR_ARTICLE_NON_EXIST"
)
