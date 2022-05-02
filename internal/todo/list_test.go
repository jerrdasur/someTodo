package todo

import (
	"testing"
)

func TestList_Add(t *testing.T) {
	t.Run("simple succes case", func(t *testing.T) {
		l := NewList()
		title := "task 1 for success case"

		got := l.Add(title)
		if got.Title != title {
			t.Fatalf("got: %s\nexpected: %s\n", got.Title, title)
		}
		if len(l.todos) != 1 {
			t.Fatalf("got length: %d\nexpected: %d\n", len(l.todos), 1)
		}
		if l.todos[0].Title != title {
			t.Fatalf("got: %s\nexpected: %s\n", l.todos[0].Title, title)
		}
	})
}
