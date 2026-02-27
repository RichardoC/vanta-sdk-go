package v1

import (
	"context"
	"fmt"
)

// PageInfo is the common cursor metadata returned by Vanta list endpoints.
type PageInfo struct {
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
}

// ResultsPage is a generic envelope for list endpoints with results.data/pageInfo.
type ResultsPage[T any] struct {
	Results struct {
		Data     []T      `json:"data"`
		PageInfo PageInfo `json:"pageInfo"`
	} `json:"results"`
}

// Pager iterates over cursor-paginated endpoints.
type Pager[T any] struct {
	fetch func(context.Context, string) (*ResultsPage[T], error)
	next  string
	done  bool
}

// NewPager creates a cursor pager with an optional initial cursor.
func NewPager[T any](initialCursor string, fetch func(context.Context, string) (*ResultsPage[T], error)) *Pager[T] {
	if fetch == nil {
		fetch = func(context.Context, string) (*ResultsPage[T], error) {
			return nil, fmt.Errorf("pager fetch must not be nil")
		}
	}
	return &Pager[T]{fetch: fetch, next: initialCursor}
}

// Next fetches the next page. Returns false when exhausted.
func (p *Pager[T]) Next(ctx context.Context) (*ResultsPage[T], bool, error) {
	if p.done {
		return nil, false, nil
	}
	page, err := p.fetch(ctx, p.next)
	if err != nil {
		return nil, false, err
	}
	if page == nil {
		return nil, false, fmt.Errorf("pager fetch returned nil page")
	}
	if !page.Results.PageInfo.HasNextPage {
		p.done = true
	} else {
		if page.Results.PageInfo.EndCursor == "" {
			return nil, false, fmt.Errorf("pager fetch returned hasNextPage=true with empty endCursor")
		}
		if page.Results.PageInfo.EndCursor == p.next {
			return nil, false, fmt.Errorf("pager fetch returned repeated endCursor %q with hasNextPage=true", p.next)
		}
		p.next = page.Results.PageInfo.EndCursor
	}
	return page, true, nil
}
