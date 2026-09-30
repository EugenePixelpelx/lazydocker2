package gocui

import (
	"strings"
	"testing"
)

func TestSearchUsesWrappedViewCoordinates(t *testing.T) {
	v := newView("main", 0, 0, 14, 8, OutputNormal)
	v.Wrap = true
	v.SetContent("prefix words domain value")

	if err := v.Search("domain"); err != nil {
		t.Fatal(err)
	}
	if len(v.searcher.searchPositions) != 1 {
		t.Fatalf("expected one match, got %d", len(v.searcher.searchPositions))
	}
	if got, want := v.searcher.searchPositions[0], (cellPos{x: 0, y: 1}); got != want {
		t.Fatalf("expected match at %+v, got %+v", want, got)
	}
}

func TestSearchNavigationUsesWrappedViewHeight(t *testing.T) {
	v := newView("main", 0, 0, 14, 4, OutputNormal)
	v.Wrap = true
	v.SetContent(strings.Repeat("first line words ", 4) + "domain")

	if err := v.Search("domain"); err != nil {
		t.Fatal(err)
	}
	if len(v.searcher.searchPositions) == 0 {
		t.Fatal("expected at least one match")
	}
	match := v.searcher.searchPositions[0]
	if match.y <= len(v.lines)-1 {
		t.Fatalf("expected wrapped match below logical line %d, got row %d", len(v.lines)-1, match.y)
	}
	if got := v.oy + v.cy; got != match.y {
		t.Fatalf("expected focus at row %d, got row %d", match.y, got)
	}
}
