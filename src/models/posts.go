package models

import (
	"errors"
	"strings"
	"time"
)

// Post represents a post made by the user
type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"criatedAt,omitempty"`
}

// Prepare calls methods to validate and format the posts.
func (post *Post) Prepare() error {
	if err := post.Validate(); err != nil {
		return err
	}

	post.Format()
	return nil
}

func (post *Post) Validate() error {
	if post.Title == "" {
		return errors.New("Title is mandatory and can't be left blank")
	}

	if post.Content == "" {
		return errors.New("Content is mandatory and can't be left blank")
	}

	return nil
}

func (post *Post) Format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
