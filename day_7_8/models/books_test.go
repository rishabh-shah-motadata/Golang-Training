package models

import (
	"testing"
)

func TestLibraryStoreInitialization(t *testing.T) {
	store := NewBookStore()

	if store == nil {
		t.Fatal("NewBookStore() returned nil")
	}

	if len(store.Books) != 3 {
		t.Fatalf("expected default 3 books, got %d", len(store.Books))
	}
}

func TestAddBookTDD(t *testing.T) {
	store := NewBookStore()

	book := Books{
		Title:  "New Book",
		Author: "Author A",
		Genre:  "Fiction",
		Rating: 4.5,
	}

	added := store.AddBook(book)

	if added.ID == 0 {
		t.Error("expected AddBook to assign a new non-zero ID")
	}

	if len(store.Books) != 4 {
		t.Errorf("expected 4 books after AddBook, got %d", len(store.Books))
	}
}

func TestGetAllBooksTDD(t *testing.T) {
	store := NewBookStore()

	books := store.GetAllBooks()
	if len(books) != 3 {
		t.Errorf("expected 3 books from GetAllBooks, got %d", len(books))
	}
}

func TestDeleteBookTDD(t *testing.T) {
	store := NewBookStore()

	ok := store.DeleteBook(1)
	if !ok {
		t.Error("expected DeleteBook to return true for an existing book")
	}

	if len(store.Books) != 2 {
		t.Errorf("expected size 2 after deletion, got %d", len(store.Books))
	}

	ok = store.DeleteBook(99)
	if ok {
		t.Error("expected DeleteBook to return false for non-existing ID")
	}
}

func TestUpdateBookTDD(t *testing.T) {
	store := NewBookStore()

	updated := Books{
		ID:     1,
		Title:  "Updated Title",
		Author: "Updated Author",
		Genre:  "Drama",
		Rating: 4.9,
	}

	result, ok := store.UpdateBook(updated)
	if !ok {
		t.Fatal("expected UpdateBook to succeed for existing ID")
	}

	if result.Title != "Updated Title" {
		t.Error("title was not updated")
	}

	// non-existing
	_, ok = store.UpdateBook(Books{ID: 987})
	if ok {
		t.Error("expected UpdateBook to fail for non-existing ID")
	}
}
