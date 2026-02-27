package v1

import (
	"context"
	"errors"
	"testing"
)

func TestPagerNext(t *testing.T) {
	ctx := context.Background()
	calls := 0
	pager := NewPager[string]("", func(_ context.Context, cursor string) (*ResultsPage[string], error) {
		calls++
		page := &ResultsPage[string]{}
		if calls == 1 {
			if cursor != "" {
				t.Fatalf("first cursor = %q, want empty", cursor)
			}
			page.Results.Data = []string{"a"}
			page.Results.PageInfo.HasNextPage = true
			page.Results.PageInfo.EndCursor = "cursor-1"
			return page, nil
		}
		if cursor != "cursor-1" {
			t.Fatalf("second cursor = %q, want cursor-1", cursor)
		}
		page.Results.Data = []string{"b"}
		page.Results.PageInfo.HasNextPage = false
		return page, nil
	})

	p1, ok, err := pager.Next(ctx)
	if err != nil || !ok || len(p1.Results.Data) != 1 || p1.Results.Data[0] != "a" {
		t.Fatalf("unexpected first page: ok=%v err=%v data=%v", ok, err, p1.Results.Data)
	}

	p2, ok, err := pager.Next(ctx)
	if err != nil || !ok || len(p2.Results.Data) != 1 || p2.Results.Data[0] != "b" {
		t.Fatalf("unexpected second page: ok=%v err=%v data=%v", ok, err, p2.Results.Data)
	}

	_, ok, err = pager.Next(ctx)
	if err != nil {
		t.Fatalf("unexpected error on exhausted pager: %v", err)
	}
	if ok {
		t.Fatalf("expected pager exhaustion")
	}
}

func TestPagerError(t *testing.T) {
	ctx := context.Background()
	want := errors.New("boom")
	pager := NewPager[string]("", func(context.Context, string) (*ResultsPage[string], error) {
		return nil, want
	})

	_, _, err := pager.Next(ctx)
	if !errors.Is(err, want) {
		t.Fatalf("expected %v, got %v", want, err)
	}
}

func TestPagerNilPage(t *testing.T) {
	pager := NewPager[string]("", func(context.Context, string) (*ResultsPage[string], error) {
		return nil, nil
	})
	_, _, err := pager.Next(context.Background())
	if err == nil {
		t.Fatal("expected error for nil page")
	}
}

func TestPagerMissingEndCursor(t *testing.T) {
	pager := NewPager[string]("", func(context.Context, string) (*ResultsPage[string], error) {
		page := &ResultsPage[string]{}
		page.Results.PageInfo.HasNextPage = true
		return page, nil
	})
	_, _, err := pager.Next(context.Background())
	if err == nil {
		t.Fatal("expected error for missing end cursor")
	}
}

func TestPagerNilFetch(t *testing.T) {
	pager := NewPager[string]("", nil)
	_, _, err := pager.Next(context.Background())
	if err == nil {
		t.Fatal("expected error for nil fetch")
	}
}

func TestPagerRepeatedCursor(t *testing.T) {
	pager := NewPager[string]("cursor-1", func(context.Context, string) (*ResultsPage[string], error) {
		page := &ResultsPage[string]{}
		page.Results.PageInfo.HasNextPage = true
		page.Results.PageInfo.EndCursor = "cursor-1"
		return page, nil
	})
	_, _, err := pager.Next(context.Background())
	if err == nil {
		t.Fatal("expected error for repeated cursor")
	}
}
