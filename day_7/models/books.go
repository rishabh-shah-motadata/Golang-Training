package models

import (
	"sync"
	"time"
)

type Books struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Rating    float64   `json:"rating"`
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Genre     string    `json:"genre"`
}

type LibraryStore struct {
	mu    sync.RWMutex
	Books map[int]Books
}

func NewBookStore() *LibraryStore {
	return &LibraryStore{
		Books: map[int]Books{
			1: {ID: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Genre: "Fiction", Rating: 4.2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			2: {ID: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee", Genre: "Fiction", Rating: 4.3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			3: {ID: 3, Title: "1984", Author: "George Orwell", Genre: "Dystopian", Rating: 4.4, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}
}

func (ls *LibraryStore) GetAllBooks() []Books {
	ls.mu.RLock()
	defer ls.mu.RUnlock()

	books := make([]Books, 0, len(ls.Books))
	for _, book := range ls.Books {
		// I have done controlled appending here as I have already initialized the slice with proper length
		// So, no reallocation will happen during append
		books = append(books, book)
	}

	return books
}

func (ls *LibraryStore) AddBook(book *Books) {
	if book.ID == 0 {
		book.ID = len(ls.Books) + 1
	}
	
	ls.mu.Lock()
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()
	ls.Books[book.ID] = *book
	ls.mu.Unlock()
}

func (ls *LibraryStore) DeleteBook(bookID *int) bool {
	if _, exists := ls.Books[*bookID]; !exists {
		return false
	}
	
	ls.mu.Lock()
	delete(ls.Books, *bookID)
	ls.mu.Unlock()

	return true
}

func (ls *LibraryStore) UpdateBook(book *Books) bool {
	if _, exists := ls.Books[book.ID]; !exists {
		return false
	}

	ls.mu.Lock()
	book.UpdatedAt = time.Now()
	ls.Books[book.ID] = *book
	ls.mu.Unlock()

	return true
}
