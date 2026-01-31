package services

import (
	"context"
	"fmt"
	"net/url"
)

// Client interface defines the methods that services need from the main client
type Client interface {
	Get(ctx context.Context, path string, query url.Values, result interface{}) error
	Post(ctx context.Context, path string, body interface{}, result interface{}) error
	Put(ctx context.Context, path string, body interface{}, result interface{}) error
	Delete(ctx context.Context, path string) error
}

// Pagination represents pagination parameters
type Pagination struct {
	Page   int `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
	After   int `json:"after,omitempty"`
	Before  int `json:"before,omitempty"`
}

// ToQuery converts pagination to URL query values
func (p *Pagination) ToQuery() url.Values {
	q := url.Values{}
	if p.Page > 0 {
		q.Set("page", fmt.Sprintf("%d", p.Page))
	}
	if p.PerPage > 0 {
		q.Set("per_page", fmt.Sprintf("%d", p.PerPage))
	}
	if p.After > 0 {
		q.Set("after", fmt.Sprintf("%d", p.After))
	}
	if p.Before > 0 {
		q.Set("before", fmt.Sprintf("%d", p.Before))
	}
	return q
}