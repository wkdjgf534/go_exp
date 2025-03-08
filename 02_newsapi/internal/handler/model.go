package handler

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

type NewsPostRequestBody struct {
	Author    string   `json:"author"`
	Title     string   `json:"title"`
	Summary   string   `json:"summary"`
	CreatedAt string   `json:"created_at"`
	Content   string   `json:"content"`
	Source    string   `json:"source"`
	Tags      []string `json:"tags"`
}

func (n NewsPostRequestBody) Validate() (errs error){
	if n.Author == "" {
		errs = errors.Join(errs, fmt.Errorf("author is emmpty: %s", n.Author))
	}
	if n.Title == "" {
		errs = errors.Join(errs, fmt.Errorf("title is emmpty: %s", n.Title))
	}
	if n.Summary == "" {
		errs = errors.Join(errs, fmt.Errorf("summary is emmpty: %s", n.Summary))
	}
	if _, err := time.Parse(time.RFC3339, n.CreatedAt); err != nil {
		errs = errors.Join(errs, err)
	}
	if _, err := url.Parse(n.Source); err != nil {
		errs = errors.Join(errs, err)
	}
	if len(n.Tags) == 0 {
		errs = errors.Join(errs, errors.New("tags cannot be empty"))
	}
	return errs
}
